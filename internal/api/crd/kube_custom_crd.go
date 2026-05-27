package crd

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/runtime"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/crd"
	"wayne/internal/model"
)

type KubeCustomCRDController struct {
	base.APIController

	cluster   string
	namespace string
	group     string
	kind      string
	version   string
	name      string
}

func (c *KubeCustomCRDController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"List":   model.PermissionRead,
		"Get":    model.PermissionRead,
		"Create": model.PermissionCreate,
		"Update": model.PermissionUpdate,
		"Delete": model.PermissionDelete,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeCustomResourceDefinition)

	// build params
	c.cluster = c.Ctx.Input.Param(":cluster")
	c.namespace = c.Ctx.Input.Param(":namespace")
	c.group = c.Ctx.Input.Param(":group")
	c.kind = c.Ctx.Input.Param(":kind")
	c.version = c.Ctx.Input.Param(":version")
	c.name = c.Ctx.Input.Param(":name")
}

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/namespaces/{namespace}/{kind} crd reqListNamespaceKubeCustomCRD
// find CRD by cluster
// responses:
//   200: respSuccessDescription

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/{kind} crd reqListKubeCustomCRD
// find CRD by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeCustomCRDController) List() {
	param := c.BuildKubernetesQueryParam()
	cli := c.Client(c.cluster)
	result, err := crd.GetCustomCRDPage(cli, c.group, c.version, c.kind, c.namespace, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/namespaces/{namespace}/{kind}/{name} crd reqGetNamespaceKubeCustomCRD
// find CRD by cluster
// responses:
//   200: respSuccessDescription

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/{kind}/{name} crd reqGetKubeCustomCRD
// find CRD by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeCustomCRDController) Get() {
	cli := c.Client(c.cluster)
	result, err := crd.GetCustomCRD(cli, c.group, c.version, c.kind, c.namespace, c.name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route POST /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/namespaces/{namespace}/{kind} crd reqCreateNamespaceKubeCustomCRD
// create CustomResourceDefinition
// responses:
//   200: respSuccessDescription

// swagger:route POST /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/{kind} crd reqCreateKubeCustomCRD
// create CustomResourceDefinition
// responses:
//
//	200: respSuccessDescription
func (c *KubeCustomCRDController) Create() {
	cli := c.Client(c.cluster)
	result, err := crd.CreatCustomCRD(cli, c.group, c.version, c.kind, c.namespace, c.Ctx.Input.RequestBody)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route PUT /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/namespaces/{namespace}/{kind}/{name} crd reqUpdateNamespaceKubeCustomCRD
// update the CustomResourceDefinition
// responses:
//   200: respSuccessDescription

// swagger:route PUT /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/{kind}/{name} crd reqUpdateKubeCustomCRD
// update the CustomResourceDefinition
// responses:
//
//	200: respSuccessDescription
func (c *KubeCustomCRDController) Update() {
	cli := c.Client(c.cluster)
	var object runtime.Unknown
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &object)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := crd.UpdateCustomCRD(cli, c.group, c.version, c.kind, c.namespace, c.name, &object)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route DELETE /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/namespaces/{namespace}/{kind}/{name} crd reqDeleteNamespaceKubeCustomCRD
// delete the CustomResourceDefinition
// responses:
//   200: respSuccessDescription

// swagger:route DELETE /api/v1/apps/{appid}/_proxy/clusters/{cluster}/apis/{group}/{version}/{kind}/{name} crd reqDeleteKubeCustomCRD
// delete the CustomResourceDefinition
// responses:
//
//	200: respSuccessDescription
func (c *KubeCustomCRDController) Delete() {
	cli := c.Client(c.cluster)
	err := crd.DeleteCustomCRD(cli, c.group, c.version, c.kind, c.namespace, c.name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}
