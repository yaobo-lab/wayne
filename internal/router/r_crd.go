package router

import (
	"wayne/internal/api/crd"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes crd路由
func regKubeCRDRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*crd.KubeCRDController).List),
		web.NSCtrlPost("/", (*crd.KubeCRDController).Create),
		web.NSCtrlGet("/:name", (*crd.KubeCRDController).Get),
		web.NSCtrlPut("/:name", (*crd.KubeCRDController).Update),
		web.NSCtrlDelete("/:name", (*crd.KubeCRDController).Delete),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*crd.KubeCustomCRDController).List),
		web.NSCtrlPost("/", (*crd.KubeCustomCRDController).Create),
		web.NSCtrlGet("/:name", (*crd.KubeCustomCRDController).Get),
		web.NSCtrlPut("/:name", (*crd.KubeCustomCRDController).Update),
		web.NSCtrlDelete("/:name", (*crd.KubeCustomCRDController).Delete),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/customresourcedefinitions", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/apis/:group/:version/namespaces/:namespace/:kind", r2...),
		web.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/apis/:group/:version/:kind", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
