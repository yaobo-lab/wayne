package router

import (
	"wayne/internal/api/event"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes event路由
func regKubeEventRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/namespaces/:namespace/clusters/:cluster", (*event.KubeEventController).List),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/events", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
