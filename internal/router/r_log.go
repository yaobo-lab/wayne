package router

import (
	"wayne/internal/api/log"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes log路由
func regKubeLogRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/:pod/containers/:container/namespaces/:namespace/clusters/:cluster", (*log.KubeLogController).List),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/podlogs", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
