package router

import (
	"attendence/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	g.View().SetPath("template")

	s.Group("/attendence", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Attendence)
	})
}
