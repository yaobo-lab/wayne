package router

import (
	"wayne/internal/api/node"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes node路由
func regKubeNodeRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/:name/clusters/:cluster", (*node.KubeNodeController).Get),
		web.NSCtrlPut("/:name/clusters/:cluster", (*node.KubeNodeController).Update),
		web.NSCtrlDelete("/:name/clusters/:cluster", (*node.KubeNodeController).Delete),
		web.NSCtrlPost("/:name/clusters/:cluster/label", (*node.KubeNodeController).AddLabel),
		web.NSCtrlDelete("/:name/clusters/:cluster/label", (*node.KubeNodeController).DeleteLabel),
		web.NSCtrlGet("/:name/clusters/:cluster/labels", (*node.KubeNodeController).GetLabels),
		web.NSCtrlPost("/:name/clusters/:cluster/labels", (*node.KubeNodeController).AddLabels),
		web.NSCtrlDelete("/:name/clusters/:cluster/labels", (*node.KubeNodeController).DeleteLabels),
		web.NSCtrlPost("/:name/clusters/:cluster/taint", (*node.KubeNodeController).SetTaint),
		web.NSCtrlDelete("/:name/clusters/:cluster/taint", (*node.KubeNodeController).DeleteTaint),
		web.NSCtrlGet("/clusters/:cluster", (*node.KubeNodeController).List),
		web.NSCtrlGet("/statistics", (*node.KubeNodeController).NodeStatistics),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/nodes", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
