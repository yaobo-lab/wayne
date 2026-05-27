package statefulset

import (
	"encoding/json"
	"fmt"
	"net/http"

	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"

	"wayne/internal/api/base"
	"wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/namespace"
	"wayne/internal/k8s/kind/statefulset"
	"wayne/internal/model"
	util "wayne/pkg"
	"wayne/pkg/dto"
	"wayne/pkg/hack"

	"wayne/pkg/logger"
)

type KubeStatefulsetController struct {
	base.APIController
}

func (c *KubeStatefulsetController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Create": model.PermissionCreate,
		"Get":    model.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeStatefulSet)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/statefulsets/{statefulsetId}/tpls/{tplId}/clusters/{cluster} statefulset reqCreateKubeStatefulset
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeStatefulsetController) Create() {
	statefulsetId := c.GetIntParamFromURL(":statefulsetId")
	tplId := c.GetIntParamFromURL(":tplId")
	appId := c.GetIntParamFromURL(":appid")

	var kubeStatefulset appsV1.StatefulSet
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &kubeStatefulset)
	if err != nil {
		c.HandleError(err)
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Manager(cluster)

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
	statefulsetModel, err := model.StatefulsetModel.GetParseMetaDataById(int64(statefulsetId))
	if err != nil {
		c.HandleError(err)
		return
	}
	statefulsetPreDeploy(&kubeStatefulset, statefulsetModel, clusterModel, namespaceModel)

	err = checkResourceAvailable(namespaceModel, cli.KubeClient, &kubeStatefulset, cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	// 发布资源到k8s平台
	_, err = statefulset.CreateOrUpdateStatefulset(cli.Client, &kubeStatefulset)
	if err != nil {
		c.HandleError(err)
		return
	}
	err = addDeployStatus(appId, statefulsetId, tplId, cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = updateMetadata(*kubeStatefulset.Spec.Replicas, statefulsetModel, cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success("ok")

}

func addDeployStatus(appId, statefulsetId, tplId int64, cluster string) error {

	app, err := model.AppModel.GetById(appId)
	if err != nil {
		return err
	}

	if app == nil || app.Id == 0 {
		return fmt.Errorf("appId is empty:%d", appId)
	}

	// 添加发布状态
	publishStatus := model.PublishStatus{
		ResourceId:  statefulsetId,
		TemplateId:  tplId,
		Type:        model.PublishTypeStatefulSet,
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

func checkResourceAvailable(ns *model.Namespace, cli client.ResourceHandler, kubeStatefulset *appsV1.StatefulSet, cluster string) error {
	// this namespace can't use current cluster.
	clusterMetas, ok := ns.MetaDataObj.ClusterMetas[cluster]
	if !ok {
		return &dto.ErrorResult{
			Code:    http.StatusForbidden,
			SubCode: http.StatusForbidden,
			Msg:     fmt.Sprintf("Current namespace (%s) can't use current cluster (%s).Please contact administrator. ", ns.Name, cluster),
		}
	}

	// check resources
	selector := labels.SelectorFromSet(map[string]string{
		util.NamespaceLabelKey: ns.Name,
	})
	namespaceResourceUsed, err := namespace.ResourcesUsageByNamespace(cli, ns.KubeNamespace, selector.String())

	requestResourceList, err := statefulset.GetStatefulsetResource(cli, kubeStatefulset)
	if err != nil {
		return err
	}

	if clusterMetas.ResourcesLimit.Memory != 0 &&
		clusterMetas.ResourcesLimit.Memory-(namespaceResourceUsed.Memory+requestResourceList.Memory)/(1024*1024*1024) < 0 {
		return &dto.ErrorResult{
			Code:    http.StatusForbidden,
			SubCode: base.ErrorSubCodeInsufficientResource,
			Msg:     fmt.Sprintf("request namespace resource (memory:%dGi) is not enough for this deploy", requestResourceList.Memory/(1024*1024*1024)),
		}
	}

	if clusterMetas.ResourcesLimit.Cpu != 0 &&
		clusterMetas.ResourcesLimit.Cpu-(namespaceResourceUsed.Cpu+requestResourceList.Cpu)/1000 < 0 {
		return &dto.ErrorResult{
			Code:    http.StatusForbidden,
			SubCode: base.ErrorSubCodeInsufficientResource,
			Msg:     fmt.Sprintf("request namespace resource (cpu:%d) is not enough for this deploy", requestResourceList.Cpu/1000),
		}
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

func updateMetadata(replicas int32, statefulset *model.Statefulset, cluster string) (err error) {
	statefulset.MetaDataObj.Replicas[cluster] = replicas
	newMetaData, err := json.Marshal(&statefulset.MetaDataObj)
	if err != nil {
		logger.Errorf("statefulset metadata marshal error.%v", err)
		return
	}
	statefulset.MetaData = string(newMetaData)
	err = model.StatefulsetModel.UpdateById(statefulset)
	if err != nil {
		logger.Errorf("statefulset metadata update error.%v", err)
	}
	return
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/statefulsets/{statefulset}/namespaces/{namespace}/clusters/{cluster} statefulset reqGetKubeStatefulset
// find Statefulset by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeStatefulsetController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":statefulset")
	manager := c.Manager(cluster)
	result, err := statefulset.GetStatefulsetDetail(manager.Client, manager.CacheFactory, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}
