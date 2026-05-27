package router

import (
	"wayne/internal/api/openapi"

	"github.com/beego/beego/v2/server/web"
)

// 注册openapi路由
func regOpenAPIRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/get_app_deploys", (*openapi.OpenAPIController).ListAppDeploys),
		web.NSCtrlGet("/get_deployment_detail", (*openapi.OpenAPIController).GetDeploymentDetail),
		web.NSCtrlGet("/get_deployment_status", (*openapi.OpenAPIController).GetDeploymentStatus),
		web.NSCtrlGet("/get_latest_deployment_tpl", (*openapi.OpenAPIController).GetLatestDeploymentTpl),
		web.NSCtrlGet("/get_resource_info", (*openapi.OpenAPIController).GetResourceInfo),
		web.NSCtrlGet("/list_namespace_apps", (*openapi.OpenAPIController).ListNamespaceApps),
		web.NSCtrlGet("/restart_deployment", (*openapi.OpenAPIController).RestartDeployment),
		web.NSCtrlGet("/scale_deployment", (*openapi.OpenAPIController).ScaleDeployment),
		web.NSCtrlGet("/upgrade_deployment", (*openapi.OpenAPIController).UpgradeDeployment),
	}

	ns := web.NewNamespace("/openapi/v1",
		web.NSNamespace("/gateway/action", r...),
	)

	//注册路由组
	web.AddNamespace(ns)
}
