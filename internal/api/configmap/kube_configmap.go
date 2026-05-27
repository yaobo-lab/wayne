package configmap

import (
	"encoding/json"
	"fmt"

	kapi "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/configmap"
	"wayne/internal/model"
)

type KubeConfigMapController struct {
	base.APIController
}

func (c *KubeConfigMapController) Prepare() {
	c.APIController.Prepare()
	methodActionMap := map[string]string{
		"Create": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeConfigMap)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/configmaps/{configMapId}/tpls/{tplId}/clusters/{cluster} configmap reqCreateKubeConfigMap
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeConfigMapController) Create() {

	configMapId := c.GetIntParamFromURL(":configMapId")
	tplId := c.GetIntParamFromURL(":tplId")

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

	var kubeConfigMap kapi.ConfigMap
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &kubeConfigMap)
	if err != nil {
		c.HandleError(err)
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Client(cluster)

	// 发布资源到k8s平台
	_, err = configmap.CreateOrUpdateConfigMap(cli, &kubeConfigMap)
	if err != nil {
		c.HandleError(err)
		return
	}

	// 添加发布状态
	publishStatus := model.PublishStatus{
		ResourceId:  int64(configMapId),
		TemplateId:  int64(tplId),
		Type:        model.PublishTypeConfigMap,
		Cluster:     cluster,
		AppId:       appId,
		NamespaceId: app.Namespace.Id,
	}

	err = model.PublishStatusModel.Publish(&publishStatus)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok")

}
