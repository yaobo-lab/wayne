package router

import (
	"wayne/internal/api/appstarred"

	"github.com/beego/beego/v2/server/web"
)

// 注册appstarred路由
func regAppStarredRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/", (*appstarred.AppStarredController).Create),
		web.NSCtrlDelete("/:appId", (*appstarred.AppStarredController).Delete),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/stars", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
