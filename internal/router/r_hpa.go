package router

import (
	"wayne/internal/api/hpa"

	"github.com/beego/beego/v2/server/web"
)

// 注册hpa路由
func regHPARouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*hpa.HPAController).List),
		web.NSCtrlPost("/", (*hpa.HPAController).Create),
		web.NSCtrlDelete("/:id([0-9]+)", (*hpa.HPAController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*hpa.HPAController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*hpa.HPAController).Update),
		web.NSCtrlGet("/names", (*hpa.HPAController).GetNames),
		web.NSCtrlPut("/updateorders", (*hpa.HPAController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*hpa.HPATplController).List),
		web.NSCtrlPost("/", (*hpa.HPATplController).Create),
		web.NSCtrlGet("/:id([0-9]+)", (*hpa.HPATplController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*hpa.HPATplController).Update),
		web.NSCtrlDelete("/:id([0-9]+)", (*hpa.HPATplController).Delete),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/hpas", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/hpas/tpls", r2...),
	)

	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes hpa路由
func regKubeHPARouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/:hpaId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster", (*hpa.KubeHPAController).Create),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/hpas", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
