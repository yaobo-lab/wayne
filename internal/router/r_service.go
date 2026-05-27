package router

import (
	"wayne/internal/api/service"

	"github.com/beego/beego/v2/server/web"
)

// 注册service路由
func regServiceRouter() {

	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*service.ServiceController).List),
		web.NSCtrlPost("/", (*service.ServiceController).Create),
		web.NSCtrlDelete("/:id([0-9]+)", (*service.ServiceController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*service.ServiceController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*service.ServiceController).Update),
		web.NSCtrlGet("/names", (*service.ServiceController).GetNames),
		web.NSCtrlPut("/updateorders", (*service.ServiceController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*service.ServiceTplController).List),
		web.NSCtrlPost("/", (*service.ServiceTplController).Create),
		web.NSCtrlPut("/:id", (*service.ServiceTplController).Update),
		web.NSCtrlDelete("/:id", (*service.ServiceTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*service.ServiceTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/services", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/services/tpls", r2...),
	)

	web.AddNamespace(ns)
}

// 注册kubernetes service路由
func regKubeServiceRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/:service/detail/namespaces/:namespace/clusters/:cluster", (*service.KubeServiceController).Get),
		web.NSCtrlPost("/:serviceId/tpls/:tplId/clusters/:cluster", (*service.KubeServiceController).Create),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/services", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
