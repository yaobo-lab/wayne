package router

import (
	"wayne/internal/api/daemonset"

	"github.com/beego/beego/v2/server/web"
)

// 注册daemonset路由
func regDaemonSetRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*daemonset.DaemonSetController).List),
		web.NSCtrlPost("/", (*daemonset.DaemonSetController).Create),
		web.NSCtrlDelete("/:id([0-9]+)", (*daemonset.DaemonSetController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*daemonset.DaemonSetController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*daemonset.DaemonSetController).Update),
		web.NSCtrlGet("/names", (*daemonset.DaemonSetController).GetNames),
		web.NSCtrlPut("/updateorders", (*daemonset.DaemonSetController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*daemonset.DaemonSetTplController).List),
		web.NSCtrlPost("/", (*daemonset.DaemonSetTplController).Create),
		web.NSCtrlPut("/:id", (*daemonset.DaemonSetTplController).Update),
		web.NSCtrlDelete("/:id", (*daemonset.DaemonSetTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*daemonset.DaemonSetTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/daemonsets", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/daemonsets/tpls", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes daemonset路由
func regKubeDaemonSetRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/:daemonSet/namespaces/:namespace/clusters/:cluster", (*daemonset.KubeDaemonSetController).Get),
		web.NSCtrlDelete("/:daemonSet/namespaces/:namespace/clusters/:cluster", (*daemonset.KubeDaemonSetController).Delete),
		web.NSCtrlPost("/:daemonSetId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster", (*daemonset.KubeDaemonSetController).Create),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/daemonsets", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
