package crd

import (
	"context"
	"encoding/json"
	"fmt"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/crd"
	"wayne/internal/model"
)

type KubeCRDController struct {
	base.APIController
}

func (c *KubeCRDController) Prepare() {

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
}

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/customresourcedefinitions crd reqListKubeCRD
// find CRD by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeCRDController) List() {
	cluster := c.Ctx.Input.Param(":cluster")
	param := c.BuildKubernetesQueryParam()
	cli := c.ApiextensionsClient(cluster)
	result, err := crd.GetCRDPage(cli, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route GET /api/v1/apps/{appid}/_proxy/clusters/{cluster}/customresourcedefinitions/{name} crd reqGetKubeCRD
// find CRD by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeCRDController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.ApiextensionsClient(cluster)

	result, err := cli.ApiextensionsV1().CustomResourceDefinitions().Get(context.Background(), name, metaV1.GetOptions{})
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route POST /api/v1/apps/{appid}/_proxy/clusters/{cluster}/customresourcedefinitions crd reqCreateKubeCRD
// create CustomResourceDefinition
// responses:
//
//	200: respSuccessDescription
func (c *KubeCRDController) Create() {
	var tpl apiextensions.CustomResourceDefinition
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.ApiextensionsClient(cluster)
	result, err := cli.ApiextensionsV1().CustomResourceDefinitions().Create(context.Background(), &tpl, metaV1.CreateOptions{})
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route PUT /api/v1/apps/{appid}/_proxy/clusters/{cluster}/customresourcedefinitions/{name} crd reqUpdateKubeCRD
// update the CustomResourceDefinition
// responses:
//
//	200: respSuccessDescription
func (c *KubeCRDController) Update() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	var tpl apiextensions.CustomResourceDefinition
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}

	if name != tpl.Name {
		c.HandleError(fmt.Errorf("name != tpl.Name"))
		return
	}

	cli := c.ApiextensionsClient(cluster)
	result, err := cli.ApiextensionsV1().CustomResourceDefinitions().Update(context.Background(), &tpl, metaV1.UpdateOptions{})
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route DELETE /api/v1/apps/{appid}/_proxy/clusters/{cluster}/customresourcedefinitions/{name} crd reqDeleteKubeCRD
// delete the CustomResourceDefinition
// responses:
//
//	200: respSuccessDescription
func (c *KubeCRDController) Delete() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.ApiextensionsClient(cluster)
	err := cli.ApiextensionsV1().CustomResourceDefinitions().Delete(context.Background(), name, metaV1.DeleteOptions{})
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")

}
