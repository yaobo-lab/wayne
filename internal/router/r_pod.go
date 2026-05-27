package router

import (
	"wayne/internal/api/pod"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes pod路由
func regKubePodRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/:pod/terminal/namespaces/:namespace/clusters/:cluster", (*pod.KubePodController).Terminal),
		web.NSCtrlGet("/namespaces/:namespace/clusters/:cluster", (*pod.KubePodController).List),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/pods", r...),
		web.NSCtrlGet("/kubernetes/pods/statistics", (*pod.KubePodController).PodStatistics),
	)
	//注册路由组
	web.AddNamespace(ns)
}
