package router

import (
	"wayne/internal/api/cluster"

	"github.com/beego/beego/v2/server/web"
)

// 注册cluster路由
func regClusterRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/", (*cluster.ClusterController).Create),
		web.NSCtrlGet("/", (*cluster.ClusterController).List),
		web.NSCtrlPut("/:name", (*cluster.ClusterController).Update),
		web.NSCtrlGet("/:name", (*cluster.ClusterController).Get),
		web.NSCtrlDelete("/:name", (*cluster.ClusterController).Delete),
		web.NSCtrlGet("/names", (*cluster.ClusterController).GetNames),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/clusters", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
