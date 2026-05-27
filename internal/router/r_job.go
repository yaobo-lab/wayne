package router

import (
	"wayne/internal/api/job"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes job路由
func regKubeJobRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/namespaces/:namespace/clusters/:cluster", (*job.KubeJobController).ListJobByCronJob),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/jobs", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
