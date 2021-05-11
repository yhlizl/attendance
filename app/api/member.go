package api

import (
	_ "github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Member = memberApi{}

type memberApi struct {
}

// Index is a demonstration route handler for output "Hello World!".
func (a *memberApi) Register(r *ghttp.Request) {
	name := r.GetString("name")
	number := r.GetString("number")
	dep := r.GetString("dep")
	r.Response.WriteJson(g.Map{
		"name":   name,
		"number": number,
		"dep":    dep,
	})
	/*
			db:=g.DB("MITD")
			_, err := db.Model("member").Data(g.Map{"name": name,"number":number,"dep":dep,}).Insert()

		 	if err != nil{
			panic(err)
		 	}
	*/
	db := g.DB("MITD")
	_, _ = db.Exec(`insert into member (name,number,dep) values ('` + name + `','` + number + `','` + dep + `')`)

}

// Index is a demonstration route handler for output "Hello World!".
func (a *memberApi) Deleter(r *ghttp.Request) {
	name := r.GetString("name")
	number := r.GetString("number")
	dep := r.GetString("dep")
	r.Response.WriteJson(g.Map{
		"name":   name,
		"number": number,
		"dep":    dep,
	})
	/*
		db:=g.DB("MITD")
		_, err := db.Model("member").Data(g.Map{"name": name,"number":number,"dep":dep,}).Insert()

		 if err != nil{
		panic(err)
		 }
	*/
	db := g.DB("MITD")
	_, _ = db.Exec(`delete from member where number = '` + number + `'`)

}

func (a *memberApi) Show(r *ghttp.Request) {

	r.Response.WriteTpl("member/member.html")

}

func (a *memberApi) Getinit(r *ghttp.Request) {
	var (
		id     string
		name   string
		number string
		dep    string
	)

	data := g.List{}

	db := g.DB("MITD")
	row, _ := db.Query(`select * from member`)
	for row.Next() {
		row.Scan(&id, &name, &dep, &number)
		data = append(data, g.Map{
			"name":   name,
			"dep":    dep,
			"number": number,
		})
	}
	r.Response.WriteJson(data)

}

func Index(r *ghttp.Request) {

	data := r.Session.Map()["name"]
	if data != nil && data != "" {

		r.Response.WriteTpl("main/main.html",
			g.Map{
				"user": data,
			})
	} else {
		r.Response.WriteTpl("index.html", g.Map{"user": data})
	}
}

func Checkmember(r *ghttp.Request) {

	number := r.GetString("number")
	name := ""
	db := g.DB("MITD")
	row, _ := db.Query(`select name from member where number = '` + number + `'`)
	for row.Next() {
		row.Scan(&name)

	}
	if name != "" {
		r.Session.Set("name", name)
		r.Response.WriteJson(g.Map{"name": name})
	}
}

func (a *memberApi) MiddlewareAuth(r *ghttp.Request) {
	if r.Session.Contains("name") {
		r.Middleware.Next()
	} else {
		// 获取用错误码
		r.Response.WriteJson(g.Map{
			"code": 403,
			"msg":  "已登出",
		})
	}
}
func Logout(r *ghttp.Request) {

	r.Session.Clear()
	v := g.View()
	v.SetPath("template/Afternoon-tea-main")
	r.Response.WriteTpl("index.html")

}
