package pv

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/pv"
	"wayne/internal/model"
)

type KubePersistentVolumeController struct {
	base.APIController
}

func (c *KubePersistentVolumeController) Prepare() {

	c.APIController.Prepare()
	methodActionMap := map[string]string{
		"List":   model.PermissionRead,
		"Get":    model.PermissionRead,
		"Create": model.PermissionCreate,
		"Update": model.PermissionUpdate,
		"Delete": model.PermissionDelete,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubePersistentVolume)

}

// swagger:route GET /api/v1/kubernetes/persistentvolumes/clusters/{cluster} pv reqListKubePersistentVolume
// find pv by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubePersistentVolumeController) List() {
	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Client(cluster)
	result, err := pv.ListPersistentVolume(cli, metaV1.ListOptions{})
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route GET /api/v1/kubernetes/persistentvolumes/{name}/clusters/{cluster} pv reqGetKubePersistentVolume
// find pv by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubePersistentVolumeController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)

	result, err := pv.GetPersistentVolumeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route POST /api/v1/kubernetes/persistentvolumes/clusters/{cluster} pv reqCreateKubePersistentVolume
// create PersistentVolume
// responses:
//
//	200: respSuccessDescription
func (c *KubePersistentVolumeController) Create() {
	var pvTpl v1.PersistentVolume
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pvTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	cluster := c.Ctx.Input.Param(":cluster")
	cli := c.Client(cluster)
	result, err := pv.CreatePersistentVolume(cli, &pvTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route PUT /api/v1/kubernetes/persistentvolumes/{name}/clusters/{cluster} pv reqUpdateKubePersistentVolume
// update the PersistentVolume
// responses:
//
//	200: respSuccessDescription
func (c *KubePersistentVolumeController) Update() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	var pvTpl v1.PersistentVolume
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pvTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if name != pvTpl.Name {
		c.HandleError(fmt.Errorf(" name != pvTpl.Name"))
		return
	}

	cli := c.Client(cluster)
	result, err := pv.UpdatePersistentVolume(cli, &pvTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route DELETE /api/v1/kubernetes/persistentvolumes/{name}/clusters/{cluster} pv reqDeleteKubePersistentVolume
// delete the PersistentVolume
// responses:
//
//	200: respSuccessDescription
func (c *KubePersistentVolumeController) Delete() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)
	err := pv.DeletePersistentVolume(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")

}
