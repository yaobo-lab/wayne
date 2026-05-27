package router

import (
	"wayne/internal/api/secret"

	"github.com/beego/beego/v2/server/web"
)

// 注册secret路由
func regSecretRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*secret.SecretController).List),
		web.NSCtrlPost("/", (*secret.SecretController).Create),
		web.NSCtrlDelete("/:id", (*secret.SecretController).Delete),
		web.NSCtrlPut("/:id", (*secret.SecretController).Update),
		web.NSCtrlGet("/:id([0-9]+)", (*secret.SecretController).Get),
		web.NSCtrlGet("/names", (*secret.SecretController).GetNames),
		web.NSCtrlPut("/updateorders", (*secret.SecretController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*secret.SecretTplController).List),
		web.NSCtrlPost("/", (*secret.SecretTplController).Create),
		web.NSCtrlPut("/:id", (*secret.SecretTplController).Update),
		web.NSCtrlDelete("/:id", (*secret.SecretTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*secret.SecretTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/secrets", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/secrets/tpls", r2...),
	)

	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes secret路由
func regKubeSecretRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlPost("/:secretId/tpls/:tplId/clusters/:cluster", (*secret.KubeSecretController).Create),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/secrets", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
