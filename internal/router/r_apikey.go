package router

import (
	"wayne/internal/api/apikey"

	"github.com/beego/beego/v2/server/web"
)

// 注册apikey路由
func regApikeyRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/", (*apikey.ApiKeyController).List),
		web.NSCtrlPost("/", (*apikey.ApiKeyController).Create),
		web.NSCtrlPut("/:id", (*apikey.ApiKeyController).Update),
		web.NSCtrlDelete("/:id", (*apikey.ApiKeyController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*apikey.ApiKeyController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		// 路由中携带namespaceid
		web.NSNamespace("/namespaces/:namespaceid([0-9]+)/apikeys", r...),
		web.NSNamespace("/apps/:appid([0-9]+)/apikeys", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
