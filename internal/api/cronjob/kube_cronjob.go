package cronjob

import (
	"encoding/json"
	"fmt"

	"k8s.io/api/batch/v1beta1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/cronjob"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type KubeCronjobController struct {
	base.APIController
}

func (c *KubeCronjobController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Get":     model.PermissionRead,
		"Delete":  model.PermissionDelete,
		"Create":  model.PermissionCreate,
		"Suspend": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeCronJob)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/cronjobs/{cronjobId}/tpls/{tplId}/clusters/{cluster}/suspend cronjob reqSuspendKubeCronjob
// Suspend CronJob
// responses:
//
//	200: respSuccessDescription
func (c *KubeCronjobController) Suspend() {
	cronjobId := c.GetIntParamFromURL(":cronjobId")
	tplId := c.GetIntParamFromURL(":tplId")
	cluster := c.Ctx.Input.Param(":cluster")
	appId := c.GetIntParamFromURL(":appid")

	cli := c.Client(cluster)

	namespaceModel, err := getNamespace(c.AppId)
	if err != nil {
		c.HandleError(err)
		return
	}
	cronjobModel, err := model.CronjobModel.GetParseMetaDataById(int64(cronjobId))
	if err != nil {
		c.HandleError(err)
		return
	}

	// 更新Suspend状态为挂起
	err = cronjob.SuspendCronjob(cli, cronjobModel.Name, namespaceModel.KubeNamespace)

	if err != nil {
		c.HandleError(err)
		return
	}
	err = addDeployStatus(appId, cronjobId, tplId, cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success("ok")
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/cronjobs/{cronjobId}/tpls/{tplId}/clusters/{cluster} cronjob reqCreateKubeCronjob
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeCronjobController) Create() {
	cronjobId := c.GetIntParamFromURL(":cronjobId")
	tplId := c.GetIntParamFromURL(":tplId")
	appId := c.GetIntParamFromURL(":appid")

	var kubeCronJob v1beta1.CronJob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCronJob)
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
	cronjobModel, err := model.CronjobModel.GetParseMetaDataById(int64(cronjobId))
	if err != nil {
		c.HandleError(err)
		return
	}

	cronjobPreDeploy(&kubeCronJob, cronjobModel, clusterModel, namespaceModel)

	// 发布资源到k8s平台
	_, err = cronjob.CreateOrUpdateCronjob(cli, &kubeCronJob)

	if err != nil {
		c.HandleError(err)
		return
	}

	err = addDeployStatus(appId, cronjobId, tplId, cluster)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok")

}

func addDeployStatus(appId, resourceId, tplId int64, cluster string) error {

	app, err := model.AppModel.GetById(appId)
	if err != nil {
		return err
	}

	if app == nil || app.Id == 0 {
		return fmt.Errorf("appId is empty:%d", appId)
	}

	// 添加发布状态
	publishStatus := model.PublishStatus{
		ResourceId:  resourceId,
		TemplateId:  tplId,
		Type:        model.PublishTypeCronJob,
		Cluster:     cluster,
		AppId:       app.Id,
		NamespaceId: app.Namespace.Id,
	}
	err = model.PublishStatusModel.Publish(&publishStatus)
	if err != nil {
		return err
	}
	return nil
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

// swagger:route GET /api/v1/kubernetes/apps/{appid}/cronjobs/{cronjob}/namespaces/{namespace}/clusters/{cluster} cronjob reqGetKubeCronjob
// find Cronjob by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeCronjobController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":cronjob")

	cli := c.Client(cluster)

	result, err := cronjob.GetCronjobDetail(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route DELETE /api/v1/kubernetes/apps/{appid}/cronjobs/{cronjob}/namespaces/{namespace}/clusters/{cluster} cronjob reqDeleteKubeCronjob
// delete the Cronjob
// responses:
//
//	200: respSuccessDescription
func (c *KubeCronjobController) Delete() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":cronjob")
	cli := c.Client(cluster)

	err := cronjob.DeleteCronjob(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}
