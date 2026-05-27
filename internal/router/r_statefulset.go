package router

import (
	"wayne/internal/api/statefulset"

	"github.com/beego/beego/v2/server/web"
)

// 注册statefulset路由
func regStatefulSetRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*statefulset.StatefulsetController).List),
		web.NSCtrlPost("/", (*statefulset.StatefulsetController).Create),
		web.NSCtrlDelete("/:id([0-9]+)", (*statefulset.StatefulsetController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*statefulset.StatefulsetController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*statefulset.StatefulsetController).Update),
		web.NSCtrlGet("/names", (*statefulset.StatefulsetController).GetNames),
		web.NSCtrlPut("/updateorders", (*statefulset.StatefulsetController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*statefulset.StatefulsetTplController).List),
		web.NSCtrlPost("/", (*statefulset.StatefulsetTplController).Create),
		web.NSCtrlPut("/:id", (*statefulset.StatefulsetTplController).Update),
		web.NSCtrlDelete("/:id", (*statefulset.StatefulsetTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*statefulset.StatefulsetTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/statefulsets", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/statefulsets/tpls", r2...),
	)

	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes statefulset路由
func regKubeStatefulSetRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/:statefulset/namespaces/:namespace/clusters/:cluster", (*statefulset.KubeStatefulsetController).Get),
		web.NSCtrlPost("/:statefulsetId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster", (*statefulset.KubeStatefulsetController).Create),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/statefulsets", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
