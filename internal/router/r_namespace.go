package router

import (
	"wayne/internal/api/namespace"

	"github.com/beego/beego/v2/server/web"
)

// 注册namespace路由
func regNamespaceRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/", (*namespace.NamespaceController).List),
		web.NSCtrlPost("/", (*namespace.NamespaceController).Create),
		web.NSCtrlDelete("/:id", (*namespace.NamespaceController).Delete),
		web.NSCtrlPut("/:id", (*namespace.NamespaceController).Update),
		web.NSCtrlGet("/:id([0-9]+)", (*namespace.NamespaceController).Get),
		web.NSCtrlGet("/:id([0-9]+)/statistics", (*namespace.NamespaceController).Statistics),
		web.NSCtrlGet("/init", (*namespace.NamespaceController).InitDefault),
		web.NSCtrlGet("/names", (*namespace.NamespaceController).GetNames),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/namespaces", r...),
	)

	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes namespace路由
func regKubeNamespaceRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/:name/clusters/:cluster", (*namespace.KubeNamespaceController).Create),
		web.NSCtrlGet("/:namespaceid([0-9]+)/resources", (*namespace.KubeNamespaceController).Resources),
		web.NSCtrlGet("/:namespaceid([0-9]+)/statistics", (*namespace.KubeNamespaceController).Statistics),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/namespaces", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
