package router

import (
	"wayne/internal/api/ingress"

	"github.com/beego/beego/v2/server/web"
)

// 注册ingress路由
func regIngressRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*ingress.IngressController).List),
		web.NSCtrlPost("/", (*ingress.IngressController).Create),
		web.NSCtrlDelete("/:id([0-9]+)", (*ingress.IngressController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*ingress.IngressController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*ingress.IngressController).Update),
		web.NSCtrlGet("/names", (*ingress.IngressController).GetNames),
		web.NSCtrlPut("/updateorders", (*ingress.IngressController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*ingress.IngressTplController).List),
		web.NSCtrlPost("/", (*ingress.IngressTplController).Create),
		web.NSCtrlGet("/:id([0-9]+)", (*ingress.IngressTplController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*ingress.IngressTplController).Update),
		web.NSCtrlDelete("/:id([0-9]+)", (*ingress.IngressTplController).Delete),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/ingresses", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/ingresses/tpls", r2...),
	)

	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes ingress路由
func regKubeIngressRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/:ingressId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster", (*ingress.KubeIngressController).Create),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/ingresses", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
