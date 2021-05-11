package api

import (
	"fmt"
	_ "github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Meal = mealApi{}

type mealApi struct {
}

func (a *mealApi) Getmeal(r *ghttp.Request) {
	today := r.GetString("today")

	/*
		db:=g.DB("MITD")
		_, err := db.Model("member").Data(g.Map{"name": name,"number":number,"dep":dep,}).Insert()

		 if err != nil{
		panic(err)
		 }
	*/
		var (
			id     string
			store  string
			stype  string
			date   string
			memo   string
			status string
		)

	res1 := g.List{}
	res2 := g.List{}
	db := g.DB("MITD")
	row1, _ := db.Query(`select id, store, type, date, memo from meal_info`)
	for row1.Next() {

		row1.Scan(&id, &store, &stype, &date, &memo)
		res1 = append(res1, g.Map{
			"id":    id,
			"store": store,
			"type":  stype,
			"date":  date,
			"memo":  memo,
		})
	}
	row1.Close()
	row2, _ := db.Query(`select id, store, type, date, memo, status from meal_info where date='` + today + `' and status='open' `)
	for row2.Next() {

		row2.Scan(&id, &store, &stype, &date, &memo, &status)
		res2 = append(res2, g.Map{
			"id":    id,
			"store": store,
			"type":  stype,
			"date":  date,
			"memo":  memo,
		})
	}
	row2.Close()

	r.Response.WriteJson(g.Map{
		"today_meal": res2,
		"data":       res1,
	})

}

func (a *mealApi) Create(r *ghttp.Request) {
	store := r.GetString("store")
	date := r.GetString("date")
	stype := r.GetString("type")
	memo := r.GetString("memo")
	confirm := r.Get("con")
	imagestring := r.GetString("imagestring")
	status := r.GetString("status")

	var (
		s_id    string
		s_store string
		s_stype string
		s_date  string
		s_memo  string
	)
	s_id = ""

	db := g.DB("MITD")
	if confirm == nil {
		res1 := g.List{}

		row1, _ := db.Query(`select * from meal_info where store = '` + store + `'`)
		for row1.Next() {

			row1.Scan(&s_id, &s_store, &s_stype, &s_date, &s_memo)
			res1 = append(res1, g.Map{
				"id":    s_id,
				"store": store,
				"type":  stype,
				"date":  date,
				"memo":  memo,
			})
		}
		row1.Close()
		if len(res1) > 0 {

			r.Response.WriteJson(g.Map{
				"success": false,
				"status":  "此店家已存在，是否修改店家資訊？",
				"data":    res1,
			})
			return
		}
	}
	sql := ""
	if status == "close" {
		sql = `insert into meal_info(image,date,type,store,memo) values('` + imagestring + `','` + date + `','` + stype + `','` + store + `','` + memo + `')
			on duplicate key update 
			image = '` + imagestring + `',
			date = '` + date + `',
			type = '` + stype + `',
			store='` + store + `',
			memo = '` + memo + `'
		`
	} else {
		sql = `insert into meal_info(image,date,type,store,memo,status,order_count) values('` + imagestring + `','` + date + `','` + stype + `','` + store + `','` + memo + `','open',1)
		on duplicate key update 
			image = '` + imagestring + `',
			date = '` + date + `',
			type = '` + stype + `',
			store='` + store + `',
			memo = '` + memo + `',
			status='open',
			order_count='1'
		
		`
	}
	row2p, _ := db.Prepare(sql)
	_, err := row2p.Exec()
	if err != nil {
		fmt.Println(err)
	}
	r.Response.WriteJson(g.Map{
		"success": true,
		"data": g.Map{
			"sql":   sql,
			"store": store,
			"type":  stype,
			"date":  date,
			"memo":  memo,
			//"image": imagestring,
		},
	})

}

func (a *mealApi) Index(r *ghttp.Request) {

	v := g.View()
	v.SetPath("template/Afternoon-tea-main")
	store := r.GetString("store")
	//id := r.GetString("id")
	//stype := r.GetString("type")
	//smemo := r.GetString("memo")
	var (
		s_image string
		s_id    string
		//s_store       string
		s_stype       string
		s_date        string
		s_memo        string
		s_status      string
		s_order_count string
	)
	sql := `select id, image, date, type, memo, status, order_count from meal_info  where store like '` + store + `'`
	if store != "" {
		db := g.DB("MITD")
		row1, err := db.Query(sql)
		if err != nil {
			fmt.Println(err)
		}
		for row1.Next() {

			row1.Scan(&s_id, &s_image, &s_date, &s_stype, &s_memo, &s_status, &s_order_count)
			//fmt.Println("data:", s_id, s_image, s_date, s_memo)
		}
		row1.Close()
	}

	r.Response.WriteTpl("meal/create.html", g.Map{
		"id":    s_id,
		"store": store,
		"type":  s_stype,
		"date":  s_date,
		"memo":  s_memo,
		"image": s_image,
	})

	//test

	// r.Response.WriteJson(g.Map{
	// 	"sql":    sql,
	// 	"id":     s_id,
	// 	"store":  s_store,
	// 	"store2": store,
	// 	"type":   s_stype,
	// 	"date":   s_date,
	// 	"memo":   s_memo,
	// 	"image":  s_image,
	// })

}
