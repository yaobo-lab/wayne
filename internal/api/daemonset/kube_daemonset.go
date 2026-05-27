package daemonset

import (
	"encoding/json"
	"fmt"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/daemonset"
	"wayne/internal/model"
	"wayne/pkg/hack"

	appsV1 "k8s.io/api/apps/v1"
)

type KubeDaemonSetController struct {
	base.APIController
}

func (c *KubeDaemonSetController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Get":    model.PermissionRead,
		"Delete": model.PermissionDelete,
		"Create": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeDaemonSet)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/daemonsets/{daemonSetId}/tpls/{tplId}/clusters/{cluster} daemonset reqCreateKubeDaemonSet
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeDaemonSetController) Create() {
	daemonSetId := c.GetIntParamFromURL(":daemonSetId")
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

	var kubeDaemonSet appsV1.DaemonSet
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &kubeDaemonSet)
	if err != nil {
		c.HandleError(err)
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Client(cluster)
	namespaceModel, err := getNamespace(c.AppId)
	if err != nil {
		c.HandleError(err)
		return
	}
	clusterModel, err := model.ClusterModel.GetParsedMetaDataByName(cluster)
	if err != nil {
		c.HandleError(err)
		return
	}
	DaemonSetModel, err := model.DaemonSetModel.GetParseMetaDataById(int64(daemonSetId))
	if err != nil {
		c.HandleError(err)
		return
	}
	daemonSetPreDeploy(&kubeDaemonSet, DaemonSetModel, clusterModel, namespaceModel)

	// 发布资源到k8s平台
	_, err = daemonset.CreateOrUpdateDaemonSet(cli, &kubeDaemonSet)
	if err != nil {
		c.HandleError(err)
		return
	}

	// 添加发布状态
	publishStatus := model.PublishStatus{
		ResourceId:  daemonSetId,
		TemplateId:  tplId,
		Type:        model.PublishTypeDaemonSet,
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

func getNamespace(appId int64) (*model.Namespace, error) {
	app, err := model.AppModel.GetById(appId)
	if err != nil {
		return nil, err
	}

	ns, err := model.NamespaceModel.GetById(app.Namespace.Id)
	if err != nil {
		return nil, err
	}
	var namespaceMetaData model.NamespaceMetaData
	err = json.Unmarshal(hack.Slice(ns.MetaData), &namespaceMetaData)
	if err != nil {
		return nil, err
	}
	ns.MetaDataObj = namespaceMetaData
	return ns, nil
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/daemonsets/{daemonSet}/namespaces/{namespace}/clusters/{cluster} daemonset reqGetKubeDaemonSet
// find DaemonSet by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeDaemonSetController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":daemonSet")
	manager := c.Manager(cluster)

	result, err := daemonset.GetDaemonSetDetail(manager.Client, manager.CacheFactory, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
	return
}

// swagger:route DELETE /api/v1/kubernetes/apps/{appid}/daemonsets/{daemonSet}/namespaces/{namespace}/clusters/{cluster} daemonset reqDeleteKubeDaemonSet
// delete the DaemonSet
// responses:
//
//	200: respSuccessDescription
func (c *KubeDaemonSetController) Delete() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":daemonSet")
	cli := c.Client(cluster)

	err := daemonset.DeleteDaemonSet(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}
