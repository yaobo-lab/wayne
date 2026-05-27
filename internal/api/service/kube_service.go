package service

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/service"
	"wayne/internal/model"
)

type KubeServiceController struct {
	base.APIController
}

func (c *KubeServiceController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Create": model.PermissionCreate,
		"Get":    model.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeService)
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/services/{service}/detail/namespaces/{namespace}/clusters/{cluster} service reqGetKubeService
// find Deployment by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeServiceController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":service")
	manager := c.Manager(cluster)
	serviceDetail, err := service.GetServiceDetail(manager.Client, manager.CacheFactory, namespace, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(serviceDetail)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/services/{serviceId}/tpls/{tplId}/clusters/{cluster} service reqCreateKubeService
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeServiceController) Create() {
	serviceId := c.GetIntParamFromURL(":serviceId")
	tplId := c.GetIntParamFromURL(":tplId")
	var kubeService v1.Service

	appId := c.GetIntParamFromURL(":appid")
	app, err := model.AppModel.GetById(appId)
	if err != nil {
		c.HandleError(err)
		return
	}

	if app == nil || app.Id == 0 {
		c.HandleError(fmt.Errorf("appId is empty:%d", appId))
		return
	}

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &kubeService)
	if err != nil {
		c.HandleError(err)
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Client(cluster)
	namespaceModel, err := model.NamespaceModel.GetNamespaceByAppId(c.AppId)
	if err != nil {
		c.HandleError(err)
		return
	}

	clusterModel, err := model.ClusterModel.GetParsedMetaDataByName(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	// add service predeploy
	servicePreDeploy(&kubeService, clusterModel, namespaceModel)

	// 发布资源到k8s平台
	_, err = service.CreateOrUpdateService(cli, &kubeService)

	if err != nil {
		c.HandleError(err)
		return
	}

	// 添加发布状态
	publishStatus := model.PublishStatus{
		ResourceId:  int64(serviceId),
		TemplateId:  int64(tplId),
		Type:        model.PublishTypeService,
		Cluster:     cluster,
		AppId:       app.Id,
		NamespaceId: app.Namespace.Id,
	}

	err = model.PublishStatusModel.Publish(&publishStatus)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success("ok")
}

func servicePreDeploy(kubeService *v1.Service, cluster *model.Cluster, namespace *model.Namespace) {
	preDefinedAnnotationMap := make(map[string]string)

	annotationResult := make(map[string]string, 0)
	// user defined
	for k, v := range kubeService.Annotations {
		preDefinedAnnotationMap[k] = v
	}
	// cluster defined, overwrite user defined
	for k, v := range cluster.MetaDataObj.ServiceAnnotations {
		preDefinedAnnotationMap[k] = v
	}
	// namespace defined, overwrite cluster and user defined
	for k, v := range namespace.MetaDataObj.ServiceAnnotations {
		preDefinedAnnotationMap[k] = v
	}
	for k, v := range preDefinedAnnotationMap {
		annotationResult[k] = v
	}

	kubeService.Annotations = annotationResult
}
