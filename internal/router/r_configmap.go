package router

import (
	"wayne/internal/api/configmap"

	"github.com/beego/beego/v2/server/web"
)

// 注册configmap路由
func regConfigMapRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*configmap.ConfigMapController).List),
		web.NSCtrlPost("/", (*configmap.ConfigMapController).Create),
		web.NSCtrlDelete("/:id", (*configmap.ConfigMapController).Delete),
		web.NSCtrlPut("/:id", (*configmap.ConfigMapController).Update),
		web.NSCtrlGet("/:id([0-9]+)", (*configmap.ConfigMapController).Get),
		web.NSCtrlGet("/names", (*configmap.ConfigMapController).GetNames),
		web.NSCtrlPut("/updateorders", (*configmap.ConfigMapController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*configmap.ConfigMapTplController).List),
		web.NSCtrlPost("/", (*configmap.ConfigMapTplController).Create),
		web.NSCtrlPut("/:id", (*configmap.ConfigMapTplController).Update),
		web.NSCtrlDelete("/:id", (*configmap.ConfigMapTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*configmap.ConfigMapTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/configmaps", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/configmaps/tpls", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes configmap路由
func regKubeConfigMapRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/:configMapId/tpls/:tplId/clusters/:cluster", (*configmap.KubeConfigMapController).Create),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/configmaps", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
