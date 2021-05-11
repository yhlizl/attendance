package main

import (
	_ "attendence/boot"
	_ "attendence/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
