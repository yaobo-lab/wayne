package router

import (
	"wayne/internal/api/proxy"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes proxy路由
func regKubeProxyRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/", (*proxy.KubeProxyController).List),
		web.NSCtrlPost("/", (*proxy.KubeProxyController).Create),
		web.NSCtrlGet("/:name", (*proxy.KubeProxyController).Get),
		web.NSCtrlPut("/:name", (*proxy.KubeProxyController).Update),
		web.NSCtrlDelete("/:name", (*proxy.KubeProxyController).Delete),
		web.NSCtrlGet("/names", (*proxy.KubeProxyController).GetNames),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/namespaces/:namespace/:kind", r...),
		web.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/:kind", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
