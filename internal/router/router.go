// Package routers k8s APIs
//
// k8s 管理平台  api
// Terms Of Service:
//
//	BasePath: /
//	Version: 2.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- api_token:
//	- api_key:
//
//	SecurityDefinitions:
//	api_token:
//	     type: apiKey
//	     name: Authorization
//	     in: header
//	api_key:
//	     type: apiKey
//	     name: apiKey
//	     in: query
//
// swagger:meta
package router

import (
	"path"
	kpod "wayne/internal/api/pod"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/plugins/cors"
)

func Setup() {
	// Beego注解路由代码生成规则和程序运行路径相关，需要改写一下避免产生不一致的文件名
	if beego.BConfig.RunMode == "dev" && path.Base(beego.AppPath) == "_build" {
		beego.AppPath = path.Join(path.Dir(beego.AppPath), "src/backend")
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	//pod 容器内部
	beego.Handler("/ws/pods/exec/*", kpod.CreateAttachHandler("/ws/pods/exec"), true)

	regApikeyRouter()
	regAppRouter()
	regAppStarredRouter()
	regAuthRouter()
	regClusterRouter()
	regConfigRouter()
	regConfigMapRouter()
	regCronJobRouter()
	regDaemonSetRouter()
	regDeploymentRouter()
	regHPARouter()
	regIngressRouter()
	regNamespaceRouter()
	regOpenAPIRouter()
	regPermissionRouter()
	regPublishStatusRouter()
	regPVCRouter()
	regSecretRouter()
	regStatefulSetRouter()
	regServiceRouter()

	//k8s
	regKubeConfigMapRouter()
	regKubeCRDRouter()
	regKubeCronjobRouter()
	regKubeDaemonSetRouter()
	regKubeDeploymentRouter()
	regKubeEventRouter()
	regKubeHPARouter()
	regKubeIngressRouter()
	regKubeJobRouter()
	regKubeLogRouter()
	regKubeNamespaceRouter()
	regKubeNodeRouter()
	regKubePodRouter()
	regKubeProxyRouter()
	regKubePVRouter()
	regKubePVCRouter()
	regKubeSecretRouter()
	regKubeServiceRouter()
	regKubeStatefulSetRouter()
}
