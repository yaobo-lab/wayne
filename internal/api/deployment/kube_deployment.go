package deployment

import (
	"encoding/json"
	"fmt"
	"net/http"

	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"

	"wayne/internal/api/base"
	"wayne/internal/api/utils"
	"wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/deployment"
	"wayne/internal/k8s/kind/namespace"
	"wayne/internal/model"
	util "wayne/pkg"
	"wayne/pkg/dto"
)

type KubeDeploymentController struct {
	base.APIController
}

type Replica struct {
	Num int32
}

func (c *KubeDeploymentController) Prepare() {
	c.APIController.Prepare()
	methodActionMap := map[string]string{
		"List":   model.PermissionRead,
		"Get":    model.PermissionRead,
		"Delete": model.PermissionDelete,
		"Create": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeDeployment)
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/deployments/namespaces/{namespace}/clusters/{cluster} deployment reqListKubeDeployment
// get all deployment
// responses:
//
//	200: respSuccessDescription
func (c *KubeDeploymentController) List() {

	param := c.BuildQueryParam()
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	manager := c.Manager(cluster)

	result, err := deployment.GetDeploymentPage(manager.CacheFactory, namespace, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/deployments/{deploymentId}/tpls/{tplId}/clusters/{cluster} deployment reqCreateKubeDeployment
// deploy tpl
// responses:
//
//	200: respSuccessDescription
func (c *KubeDeploymentController) Create() {

	deploymentId := c.GetIntParamFromURL(":deploymentId")

	tplId := c.GetIntParamFromURL(":tplId")

	var kubeDeployment appsV1.Deployment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &kubeDeployment)

	if err != nil {
		c.HandleError(err)
		return
	}

	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Manager(cluster)

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

	deploymentModel, err := model.DeploymentModel.GetParseMetaDataById(int64(deploymentId))
	if err != nil {
		c.HandleError(err)
		return
	}

	//处理合并元数据
	utils.DeploymentPreDeploy(&kubeDeployment, deploymentModel, clusterModel, namespaceModel)

	err = checkResourceAvailable(namespaceModel, cli.KubeClient, &kubeDeployment, cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	// 发布资源到k8s平台
	_, err = deployment.CreateOrUpdateDeployment(cli.Client, &kubeDeployment)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.PublishStatusModel.Add(deploymentId, tplId, cluster, model.PublishTypeDeployment)
	// 添加发布状态
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.DeploymentModel.Update(*kubeDeployment.Spec.Replicas, deploymentModel, cluster)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success("ok")
}

func checkResourceAvailable(ns *model.Namespace, cli client.ResourceHandler, kubeDeployment *appsV1.Deployment, cluster string) error {

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

	requestResourceList, err := deployment.GetDeploymentResource(cli, kubeDeployment)
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

// swagger:route GET /api/v1/kubernetes/apps/{appid}/deployments/{deployment}/detail/namespaces/{namespace}/clusters/{cluster} deployment reqGetKubeDeployment
// find Deployment by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeDeploymentController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":deployment")
	manager := c.Manager(cluster)
	result, err := deployment.GetDeploymentDetail(manager.Client, manager.CacheFactory, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route DELETE /api/v1/kubernetes/apps/{appid}/deployments/{deployment}/namespaces/{namespace}/clusters/{cluster} deployment reqDeleteKubeDeployment
// delete the Deployment
// responses:
//
//	200: respSuccessDescription
func (c *KubeDeploymentController) Delete() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":deployment")
	cli := c.Client(cluster)

	err := deployment.DeleteDeployment(cli, name, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success("ok!")
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/deployments/{deployment}/namespaces/{namespace}/clusters/{cluster}/updatescale deployment reqUpdateScaleKubeDeployment
// Update the number of replica for target deployment
// responses:
//
//	200: respSuccessDescription
func (c *KubeDeploymentController) UpdateScale() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":deployment")
	cli := c.Client(cluster)

	var replica Replica
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &replica)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = deployment.UpdateScale(cli, name, namespace, replica.Num)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}
