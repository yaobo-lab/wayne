package router

import (
	"wayne/internal/api/deployment"

	"github.com/beego/beego/v2/server/web"
)

// 注册deployment路由
func regDeploymentRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*deployment.DeploymentController).List),
		web.NSCtrlPost("/", (*deployment.DeploymentController).Create),
		web.NSCtrlDelete("/:id([0-9]+)", (*deployment.DeploymentController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*deployment.DeploymentController).Get),
		web.NSCtrlPut("/:id([0-9]+)", (*deployment.DeploymentController).Update),
		web.NSCtrlGet("/names", (*deployment.DeploymentController).GetNames),
		web.NSCtrlPut("/updateorders", (*deployment.DeploymentController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*deployment.DeploymentTplController).List),
		web.NSCtrlPost("/", (*deployment.DeploymentTplController).Create),
		web.NSCtrlPut("/:id", (*deployment.DeploymentTplController).Update),
		web.NSCtrlDelete("/:id", (*deployment.DeploymentTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*deployment.DeploymentTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/deployments", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/deployments/tpls", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes deployment路由
func regKubeDeploymentRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/:deployment/detail/namespaces/:namespace/clusters/:cluster", (*deployment.KubeDeploymentController).Get),
		web.NSCtrlDelete("/:deployment/namespaces/:namespace/clusters/:cluster", (*deployment.KubeDeploymentController).Delete),
		web.NSCtrlPost("/:deployment/namespaces/:namespace/clusters/:cluster/updatescale", (*deployment.KubeDeploymentController).UpdateScale),
		web.NSCtrlPost("/:deploymentId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster", (*deployment.KubeDeploymentController).Create),
		web.NSCtrlGet("/namespaces/:namespace/clusters/:cluster", (*deployment.KubeDeploymentController).List),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/deployments", r...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
