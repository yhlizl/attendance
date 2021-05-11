package api

import (
	"fmt"
	_ "github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Detail = detailApi{}

type detailApi struct {
}

func (a *detailApi) Index(r *ghttp.Request) {

	v := g.View()
	v.SetPath("template/Afternoon-tea-main")
	store := r.GetString("store")

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

	r.Response.WriteTpl("detail/detail.html", g.Map{
		"id":    s_id,
		"store": store,
		"type":  s_stype,
		"date":  s_date,
		"memo":  s_memo,
		"image": s_image,
	})

}
func (a *detailApi) Getinit(r *ghttp.Request) {

	r_id := r.GetString("id")
	r_store := r.GetString("store")
	r_type := r.GetString("type")
	r_date := r.GetString("date")

	var (
		id          string
		store_id    string
		store       string
		stype       string
		name        string
		food_name   string
		drink_name  string
		drink_size  string
		drink_ice   string
		drink_sugar string
		price       string
		date        string
		memo        string
	)

	data := g.List{}
	sql := `select id,store_id,store,type,name,food_name,drink_name,drink_size,drink_ice,drink_sugar,price,memo,date 
	from order_info where store_id = '` + r_id + `' and store='` + r_store + `' and type='` + r_type + `' and date='` + r_date + `' order by food_name,drink_name`
	db := g.DB("MITD")
	row, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	for row.Next() {
		row.Scan(&id, &store_id, &store, &stype, &name, &food_name, &drink_name, &drink_size, &drink_ice, &drink_sugar, &price, &memo, &date)
		if stype == "飲料" {
			data = append(data, g.Map{
				"sql":       sql,
				"order_id":  id,
				"store_id":  store_id,
				"store":     store,
				"type":      stype,
				"name":      name,
				"food":      drink_name,
				"sugar_ice": drink_size + "," + drink_ice + "," + drink_sugar + ",",
				"size":      drink_size,
				"ice":       drink_ice,
				"suger":     drink_sugar,
				"price":     price,
				"memo":      memo,
				"date":      date,
			})
		} else {
			data = append(data, g.Map{
				"sql":      sql,
				"order_id": id,
				"store_id": store_id,
				"store":    store,
				"type":     stype,
				"name":     name,
				"food":     food_name,

				"size":  drink_size,
				"ice":   drink_ice,
				"suger": drink_sugar,
				"price": price,
				"memo":  memo,
				"date":  date,
			})

		}
	}
	r.Response.WriteJson(data)

}
func (a *detailApi) Register(r *ghttp.Request) {
	r_order_id := r.GetString("order_id")
	r_id := r.GetString("id")
	r_store := r.GetString("store")
	r_type := r.GetString("type")
	r_member := r.GetString("member")
	r_food_name := r.GetString("food_name")
	r_drink_name := r.GetString("drink_name")
	r_ice := r.GetString("ice")
	r_size := r.GetString("size")
	r_sugar := r.GetString("sugar")
	r_memo := r.GetString("memo")
	r_date := r.GetString("date")
	r_price := r.GetString("price")
	// r_orderType := r.GetString("orderType")
	// r_order_id := r.GetString("order_id")

	/*
			db:=g.DB("MITD")
			_, err := db.Model("member").Data(g.Map{"name": name,"number":number,"dep":dep,}).Insert()

		 	if err != nil{
			panic(err)
		 	}
	*/
	db := g.DB("MITD")
	sql := ""
	if r_order_id != "" {
		sql = `insert into order_info 
	(id,store_id,store,type,name,food_name,drink_name,drink_size,drink_ice,drink_sugar,price,memo,date)
	values('` + r_order_id + `','` + r_id + `','` + r_store + `','` + r_type + `','` + r_member + `','` + r_food_name + `','` + r_drink_name + `','` + r_size + `','` + r_ice + `',
	'` + r_sugar + `','` + r_price + `','` + r_memo + `','` + r_date + `')
	on duplicate key update
	id='` + r_order_id + `',
	store_id='` + r_id + `',
	store='` + r_store + `',
	type='` + r_type + `',
	name='` + r_member + `',
	food_name='` + r_food_name + `',
	drink_name='` + r_drink_name + `',
	drink_size='` + r_size + `',
	drink_ice='` + r_ice + `',
	drink_sugar='` + r_sugar + `',
	price='` + r_price + `',
	memo='` + r_memo + `',
	date='` + r_date + `'
	`
	} else {
		sql = `insert into order_info 
		(store_id,store,type,name,food_name,drink_name,drink_size,drink_ice,drink_sugar,price,memo,date)
		values('` + r_id + `','` + r_store + `','` + r_type + `','` + r_member + `','` + r_food_name + `','` + r_drink_name + `','` + r_size + `','` + r_ice + `',
		'` + r_sugar + `','` + r_price + `','` + r_memo + `','` + r_date + `')
		on duplicate key update
		
		store_id='` + r_id + `',
		store='` + r_store + `',
		type='` + r_type + `',
		name='` + r_member + `',
		food_name='` + r_food_name + `',
		drink_name='` + r_drink_name + `',
		drink_size='` + r_size + `',
		drink_ice='` + r_ice + `',
		drink_sugar='` + r_sugar + `',
		price='` + r_price + `',
		memo='` + r_memo + `',
		date='` + r_date + `'
		`
	}
	_, _ = db.Exec(sql)
	//r.Response.RedirectBack()
	r.Response.WriteJson(g.Map{
		"success": true,
	})
}

func (a *detailApi) Deleter(r *ghttp.Request) {

	r_id := r.GetString("id")

	db := g.DB("MITD")
	sql := `delete from order_info 
	where id = '` + r_id + `'
	`
	_, _ = db.Exec(sql)
	//r.Response.RedirectBack()
	r.Response.WriteJson(g.Map{
		"success": true,
		"sql":     sql,
	})
}
