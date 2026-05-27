package router

import (
	"wayne/internal/api/app"

	"github.com/beego/beego/v2/server/web"
)

// 注册app路由
func regAppRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/", (*app.AppController).List),
		web.NSCtrlPost("/", (*app.AppController).Create),
		web.NSCtrlPut("/:id", (*app.AppController).Update),
		web.NSCtrlDelete("/:id", (*app.AppController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*app.AppController).Get),
		web.NSCtrlGet("/names", (*app.AppController).GetNames),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/namespaces/:namespaceid([0-9]+)/apps", r...),
		web.NSCtrlGet("/apps/statistics", (*app.AppController).AppStatistics),
	)
	//注册路由组
	web.AddNamespace(ns)
}
