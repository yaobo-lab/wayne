package router

import (
	"wayne/internal/api/cronjob"

	"github.com/beego/beego/v2/server/web"
)

// 注册cronjob路由
func regCronJobRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*cronjob.CronjobController).List),
		web.NSCtrlPost("/", (*cronjob.CronjobController).Create),
		web.NSCtrlDelete("/:id", (*cronjob.CronjobController).Delete),
		web.NSCtrlPut("/:id", (*cronjob.CronjobController).Update),
		web.NSCtrlGet("/:id([0-9]+)", (*cronjob.CronjobController).Get),
		web.NSCtrlGet("/names", (*cronjob.CronjobController).GetNames),
		web.NSCtrlPut("/updateorders", (*cronjob.CronjobController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*cronjob.CronjobTplController).List),
		web.NSCtrlPost("/", (*cronjob.CronjobTplController).Create),
		web.NSCtrlPut("/:id", (*cronjob.CronjobTplController).Update),
		web.NSCtrlDelete("/:id", (*cronjob.CronjobTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*cronjob.CronjobTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/cronjobs", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/cronjobs/tpls", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes cronjob路由
func regKubeCronjobRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/:cronjob/namespaces/:namespace/clusters/:cluster", (*cronjob.KubeCronjobController).Get),
		web.NSCtrlDelete("/:cronjob/namespaces/:namespace/clusters/:cluster", (*cronjob.KubeCronjobController).Delete),
		web.NSCtrlPost("/:cronjobId/tpls/:tplId/clusters/:cluster", (*cronjob.KubeCronjobController).Create),
		web.NSCtrlPost("/:cronjobId/tpls/:tplId/clusters/:cluster/suspend", (*cronjob.KubeCronjobController).Suspend),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/cronjobs", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
