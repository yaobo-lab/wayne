package proxy

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"wayne/internal/api/base"
	api "wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/proxy"
	"wayne/internal/model"
)

type KubeProxyController struct {
	base.APIController
}

func (c *KubeProxyController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Get":      model.PermissionRead,
		"List":     model.PermissionRead,
		"GetNames": model.PermissionRead,
		"Create":   model.PermissionCreate,
		"Update":   model.PermissionUpdate,
		"Delete":   model.PermissionDelete,
	}
	_, method := c.GetControllerAndAction()
	kind := c.Ctx.Input.Param(":kind")
	resourceMap, ok := api.KindToResourceMap[kind]
	if !ok {
		c.HandleError(fmt.Errorf("Request resource kind (%s) not supported!", kind))
		return
	}

	c.PreparePermission(methodActionMap, method, fmt.Sprintf("KUBE%s", strings.ToUpper(resourceMap.GroupVersionResourceKind.Kind)))
}

// swagger:route GET /api/v1/apps/{appid}/_proxy/namespaces/{namespace}/clusters/{cluster}/{kind}/{name} proxy reqCreateNamespaceKubeProxy
// Find Object by name
// responses:
//   200: respSuccessDescription

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/{kind}/{name} proxy reqCreateKubeProxy
// Find Object by name
// responses:
//
//	200: respSuccessDescription
func (c *KubeProxyController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":name")
	kind := c.Ctx.Input.Param(":kind")
	kubeClient := c.KubeClient(cluster)
	result, err := kubeClient.Get(kind, namespace, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)

}

// swagger:route GET /api/v1/apps/{appid}/_proxy/namespaces/{namespace}/clusters/{cluster}/{kind}/names proxy reqGetNamesNamespaceKubeProxy
// get all names
// responses:
//   200: respSuccessDescription

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/{kind}/names proxy reqGetNamesKubeProxy
// get all names
// responses:
//
//	200: respSuccessDescription
func (c *KubeProxyController) GetNames() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	kind := c.Ctx.Input.Param(":kind")
	kubeClient := c.KubeClient(cluster)
	result, err := proxy.GetNames(kubeClient, kind, namespace)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route GET /api/v1/apps/{appid}/_proxy/namespaces/{namespace}/clusters/{cluster}/{kind} proxy reqListNamespaceKubeProxy
// List Objects
// responses:
//   200: respSuccessDescription

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/{kind} proxy reqListKubeProxy
// List Objects
// responses:
//
//	200: respSuccessDescription
func (c *KubeProxyController) List() {
	param := c.BuildKubernetesQueryParam()
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	kind := c.Ctx.Input.Param(":kind")
	kubeClient := c.KubeClient(cluster)
	result, err := proxy.GetPage(kubeClient, kind, namespace, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)

}

// swagger:route POST /api/v1/apps/{appid}/_proxy/namespaces/{namespace}/clusters/{cluster}/{kind} proxy reqCreateNamespaceKubeProxy
// Create the resource
// responses:
//   200: respSuccessDescription

// swagger:route POST /api/v1/apps/{appid}/_proxy/clusters/{cluster}/{kind} proxy reqCreateKubeProxy
// Create the resource
// responses:
//
//	200: respSuccessDescription
func (c *KubeProxyController) Create() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	kind := c.Ctx.Input.Param(":kind")
	var object runtime.Unknown
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &object)
	if err != nil {
		c.HandleError(err)
		return
	}

	kubeClient := c.KubeClient(cluster)
	result, err := kubeClient.Create(kind, namespace, &object)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route PUT /api/v1/apps/{appid}/_proxy/namespaces/{namespace}/clusters/{cluster}/{kind}/{name} proxy reqUpdateNamespaceKubeProxy
// Update the resource
// responses:
//   200: respSuccessDescription

// swagger:route PUT /api/v1/apps/{appid}/_proxy/clusters/{cluster}/{kind}/{name} proxy reqUpdateKubeProxy
// Update the resource
// responses:
//
//	200: respSuccessDescription
func (c *KubeProxyController) Update() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":name")
	kind := c.Ctx.Input.Param(":kind")
	var object runtime.Unknown
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &object)
	if err != nil {
		c.HandleError(err)
		return
	}
	kubeClient := c.KubeClient(cluster)
	result, err := kubeClient.Update(kind, namespace, name, &object)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route DELETE /api/v1/apps/{appid}/_proxy/namespaces/{namespace}/clusters/{cluster}/{kind}/{name} proxy reqDeleteNamespaceKubeProxy
// delete the resource
// responses:
//   200: respSuccessDescription

// swagger:route DELETE /api/v1/apps/{appid}/_proxy/clusters/{cluster}/{kind}/{name} proxy reqDeleteKubeProxy
// delete the resource
// responses:
//
//	200: respSuccessDescription
func (c *KubeProxyController) Delete() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	name := c.Ctx.Input.Param(":name")
	kind := c.Ctx.Input.Param(":kind")
	force := c.Input().Get("force")
	defaultPropagationPolicy := meta_v1.DeletePropagationBackground
	defaultDeleteOptions := meta_v1.DeleteOptions{
		PropagationPolicy: &defaultPropagationPolicy,
	}
	if force != "" {
		forceBool, err := strconv.ParseBool(force)
		if err != nil {
			c.HandleError(err)
			return
		}
		if forceBool {
			var gracePeriodSeconds int64 = 0
			defaultDeleteOptions.GracePeriodSeconds = &gracePeriodSeconds
		}
	}
	kubeClient := c.KubeClient(cluster)
	err := kubeClient.Delete(kind, namespace, name, &defaultDeleteOptions)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}
