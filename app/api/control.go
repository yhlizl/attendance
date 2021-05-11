package api

import (
	_ "fmt"
	_ "github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Control = controlApi{}

type controlApi struct {
}

func (a *controlApi) Close(r *ghttp.Request) {

	v := g.View()
	v.SetPath("template/Afternoon-tea-main")
	r_id := r.GetString("id")
	r_date := r.GetString("date")
	//r_store := r.GetString("store")
	sql := `update meal_info set status='close' where id='` + r_id + `';`

	db := g.DB("MITD")
	_, _ = db.Exec(sql)
	sql2 := `delete from order_info where store_id='` + r_id + `' and date = '` + r_date + `';`
	_, _ = db.Exec(sql2)

	r.Response.WriteJson(g.Map{
		"success": true,
		"sql":     sql,
		"sql2":    sql2,
	})

}
