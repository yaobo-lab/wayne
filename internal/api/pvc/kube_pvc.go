package pvc

import (
	"encoding/json"
	"fmt"

	kapi "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/pvc"
	"wayne/internal/model"
)

type KubePersistentVolumeClaimController struct {
	base.APIController
}

func (c *KubePersistentVolumeClaimController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Create": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubePersistentVolumeClaim)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/persistentvolumeclaims/{pvcId}/tpls/{tplId}/clusters/{cluster} pvc reqCreateKubePersistentVolumeClaim
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubePersistentVolumeClaimController) Create() {
	pvcId := c.GetIntParamFromURL(":pvcId")
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

	var kubePersistentVolumeClaim kapi.PersistentVolumeClaim
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &kubePersistentVolumeClaim)
	if err != nil {
		c.HandleError(err)
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Client(cluster)

	// 发布资源到k8s平台
	_, err = pvc.CreateOrUpdatePersistentVolumeClaim(cli, &kubePersistentVolumeClaim)

	if err != nil {
		c.HandleError(err)
		return
	}

	// 添加发布状态
	publishStatus := model.PublishStatus{
		ResourceId:  int64(pvcId),
		TemplateId:  int64(tplId),
		Type:        model.PublishTypePersistentVolumeClaim,
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
