package router

import (
	"wayne/internal/api/config"

	"github.com/beego/beego/v2/server/web"
)

// 注册config路由
func regConfigRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*config.BaseConfigController).ListBase),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*config.ConfigController).List),
		web.NSCtrlPost("/", (*config.ConfigController).Create),
		web.NSCtrlPut("/:id([0-9]+)", (*config.ConfigController).Update),
		web.NSCtrlGet("/:id([0-9]+)", (*config.ConfigController).Get),
		web.NSCtrlDelete("/:id([0-9]+)", (*config.ConfigController).Delete),
		web.NSCtrlGet("/system", (*config.ConfigController).ListSystem),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/configs/base", r1...),
		web.NSNamespace("/configs", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
