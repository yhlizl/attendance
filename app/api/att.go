package api

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"sort"
	"time"
)

var Attendence = attApi{}

type nullstring = sql.NullString
type attApi struct {
}

var currenttime = time.Now()

// Index is a demonstration route handler for output "Hello World!".
func (a *attApi) Index(r *ghttp.Request) {
	g.View().SetPath("template")
	r.Response.WriteTpl("layout.html", g.Map{
		"header": "F18-MITD",
		"action": "Index",
	})
}

func (a *attApi) Gettoday(r *ghttp.Request) {
	db := g.DB("MITD")
	sql := `SELECT NAME,EMPID,DATE_S, (case when arr_time is not null then 'V' end) arr,REMARK FROM attendance 
	`
	row, _ := db.Query(sql)
	var (
		NAME   nullstring
		EMPID  nullstring
		DATE_S nullstring
		ARR    nullstring
		REMARK nullstring
	)
	res := g.List{}
	for row.Next() {
		row.Scan(&NAME, &EMPID, &DATE_S, &ARR, &REMARK)
		res = append(res, g.Map{
			"NAME":   NAME.String,
			"EMPID":  EMPID.String,
			"DATE_S": DATE_S.String,
			"ARR":    ARR.String,
			"REMARK": REMARK.String,
		})
	}
	row.Close()
	r.Response.WriteJson(g.Map{
		"status": "success",
		"data":   res,
	})
}

func (a *attApi) Edit(r *ghttp.Request) {

	date_s := r.GetString("date_s")
	empid := r.GetString("empid")
	remark := r.GetString("remark")

	db := g.DB("MITD")
	sql := `update attendance set remark = '` + remark + `' where date_s = '` + date_s + `' and empid = '` + empid + `'
	`
	_, _ = db.Exec(sql)

	r.Response.WriteJson(g.Map{
		"status": "success",
		"sql":    sql,
	})
}

func (a *attApi) History(r *ghttp.Request) {
	// s := g.Server()
	// s.SetConfigWithMap(g.Map{
	// 	"ServerRoot": "/Users/royale/go/src/test/attendence/public",
	// })
	g.View().SetPath("template")
	action := r.GetString("action")
	if action == "" {
		action = currenttime.Format("2006-01")
	}
	n0 := currenttime.Format("2006-01")
	n1 := currenttime.AddDate(0, -1, 0).Format("2006-01")
	n2 := currenttime.AddDate(0, -2, 0).Format("2006-01")
	n3 := currenttime.AddDate(0, -3, 0).Format("2006-01")
	n4 := currenttime.AddDate(0, -4, 0).Format("2006-01")
	date_s := []string{}
	date_s = append(date_s, n0, n1, n2, n3, n4)
	r.Response.WriteTpl("history.html", g.Map{
		"header": "F18-MITD",
		"action": "History",
		"dayget": action,
		"dates":  date_s,
	})
}

func (a *attApi) Gethistory(r *ghttp.Request) {
	date_s := r.GetString("action")

	if date_s == "" {
		date_s = currenttime.Format("2006-01")
	}

	db := g.DB("MITD")
	sql := `
		select a.name,a.empid,a.date_s,a.arr,a.weekday,b.workday,b.holiday,b.lossday from 
		
		(SELECT NAME,EMPID,DATE_S, (case when REMARK not in ('') and REMARK is not null then REMARK when (arr_time is not null and arr_time not in ('')) then 'V' end) arr,a.weekday FROM attendance a
		where date_format(date_s,'%Y-%m') = ('` + date_s + `')
		order by DATE_S asc)a,

		
		(select name,empid,sum(case when s.weekday not in ('Saturday','Sunday') and arr in ('V') then 1 else 0 end) workday,
		sum(case when s.weekday not in ('Saturday','Sunday') and arr not in ('V') then 1 else 0 end) holiday,
		sum(case when s.weekday not in ('Saturday','Sunday') and ( arr is null)  then 1 else 0 end) lossday
		from (
		SELECT NAME,EMPID,DATE_S, (case when REMARK not in ('') and REMARK is not null then REMARK when (arr_time is not null and arr_time not in ('')) then 'V' end) arr,a.weekday FROM attendance a
		where date_format(date_s,'%Y-%m') = ('` + date_s + `')
		order by DATE_S asc)s
		group by name,empid)b
		where a.name=b.name and a.empid=b.empid
		
		

	`

	row, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	var (
		NAME    nullstring
		EMPID   nullstring
		DATE_S  nullstring
		ARR     nullstring
		WEEKDAY nullstring
		WORKDAY nullstring
		HOLIDAY nullstring
		LOSSDAY nullstring
	)
	res := g.List{}

	for row.Next() {
		row.Scan(&NAME, &EMPID, &DATE_S, &ARR, &WEEKDAY, &WORKDAY, &HOLIDAY, &LOSSDAY)
		res = append(res, g.Map{
			"NAME":    NAME.String,
			"EMPID":   EMPID.String,
			"DATE_S":  DATE_S.String,
			"ARR":     ARR.String,
			"WEEKDAY": WEEKDAY.String,
			"WORKDAY": WORKDAY.String,
			"HOLIDAY": HOLIDAY.String,
			"LOSSDAY": LOSSDAY.String,
		})
	}
	row.Close()

	col_temp := g.Map{}
	empid_temp := g.Map{}
	data_name := map[string](map[string]string){}
	data_hol := map[string]string{}
	data_work := map[string]string{}
	data_loss := map[string]string{}
	for _, v := range res {
		//	fmt.Println(v)
		col_temp[v["DATE_S"].(string)] = v["WEEKDAY"].(string)
		//	fmt.Println(col_temp[v["DATE_S"].(string)])
		empid_temp[v["EMPID"].(string)] = v["NAME"].(string)
		//	fmt.Println(empid_temp[v["EMPID"].(string)])
		if ok := data_name[v["DATE_S"].(string)]; ok == nil {
			data_name[v["DATE_S"].(string)] = map[string]string{}
		}
		data_name[v["DATE_S"].(string)][v["EMPID"].(string)] = v["ARR"].(string)
		//	fmt.Println(data_name[v["DATE_S"].(string)][v["EMPID"].(string)])
		data_hol[v["EMPID"].(string)] = v["HOLIDAY"].(string)
		data_work[v["EMPID"].(string)] = v["WORKDAY"].(string)
		data_loss[v["EMPID"].(string)] = v["LOSSDAY"].(string)
	}
	col_res := g.List{{"title": "NAME"}, {"title": "EMPID"}, {"title": "出勤日"}, {"title": "請假日"}, {"title": "缺"}}
	col_temp2 := []string{}
	for i, _ := range col_temp {
		col_temp2 = append(col_temp2, i)
	}
	sort.Strings(col_temp2)
	data_res := [][]string{}
	//data_week := []string{"", "WEEK", "", "", ""}
	for _, d := range col_temp2 {
		//data_week = append(data_week, col_temp[d].(string))
		tp := fmt.Sprintf("%s\r%s", d, col_temp[d].(string))
		col_res = append(col_res, g.Map{"title": tp})
	}
	//data_res = append(data_res, data_week)
	empid_res := []string{}
	for i, _ := range empid_temp {
		empid_res = append(empid_res, i)
	}

	for _, e := range empid_res {
		data_temp := []string{empid_temp[e].(string), e, data_work[e], data_hol[e], data_loss[e]}
		for _, d := range col_temp2 {
			data_temp = append(data_temp, data_name[d][e])
		}
		data_res = append(data_res, data_temp)
	}
	r.Response.WriteJson(g.Map{
		"status": "success",
		"column": col_res,
		"datas":  data_res,
		"month":  date_s,
		"SQL":    sql,
	})
}
