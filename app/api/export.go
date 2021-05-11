package api

import (
	_ "fmt"
	_ "github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Export = exportApi{}

type exportApi struct {
}

func (a *exportApi) Index(r *ghttp.Request) {
	id := r.GetString("id")
	store := r.GetString("store")
	date := r.GetString("date")
	stype := r.GetString("type")
	memo := r.GetString("memo")
	//status := r.GetString("status")

	v := g.View()
	v.SetPath("template/Afternoon-tea-main")

	r.Response.WriteTpl("export/export.html",
		g.Map{
			"id":    id,
			"store": store,
			"type":  stype,
			"date":  date,
			"memo":  memo,
		})

}
