package secret

import (
	"encoding/json"
	"fmt"

	kapi "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/secret"
	"wayne/internal/model"
)

type KubeSecretController struct {
	base.APIController
}

func (c *KubeSecretController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Create": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeSecret)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/secrets/{secretId}/tpls/{tplId}/clusters/{cluster} secret reqCreateKubeSecret
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeSecretController) Create() {
	secretId := c.GetIntParamFromURL(":secretId")
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

	var kubeSecret kapi.Secret
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &kubeSecret)
	if err != nil {
		c.HandleError(err)
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Client(cluster)

	// 发布资源到k8s平台
	_, err = secret.CreateOrUpdateSecret(cli, &kubeSecret)
	if err != nil {
		c.HandleError(err)
		return
	}

	// 添加发布状态
	publishStatus := model.PublishStatus{
		ResourceId:  int64(secretId),
		TemplateId:  int64(tplId),
		Type:        model.PublishTypeSecret,
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
