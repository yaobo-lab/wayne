package pvc

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type PersistentVolumeClaimController struct {
	base.APIController
}

func (c *PersistentVolumeClaimController) Prepare() {

	c.APIController.Prepare()

	perAction := ""
	_, method := c.GetControllerAndAction()
	switch method {
	case "Get", "List":
		perAction = model.PermissionRead
	case "Create":
		perAction = model.PermissionCreate
	case "Update":
		perAction = model.PermissionUpdate
	case "Delete":
		perAction = model.PermissionDelete
	}
	if perAction != "" {
		c.CheckPermission(model.PermissionTypePersistentVolumeClaim, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/persistentvolumeclaims/names pvc reqGetNamesPersistentVolumeClaim
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	pvcs, err := model.PersistentVolumeClaimModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(pvcs)
}

// swagger:route GET /api/v1/apps/{appid}/persistentvolumeclaims pvc reqListPersistentVolumeClaim
// get all PersistentVolumeClaim
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	pvc := []model.PersistentVolumeClaim{}

	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypePersistentVolumeClaim, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.PersistentVolumeClaim), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.PersistentVolumeClaim), &pvc, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range pvc {
		pvc[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, pvc))
	return
}

// swagger:route POST /api/v1/apps/{appid}/persistentvolumeclaims pvc reqCreatePersistentVolumeClaim
// create PersistentVolumeClaim
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimController) Create() {
	var pvc model.PersistentVolumeClaim
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pvc)
	if err != nil {
		c.HandleError(err)
		return
	}

	pvc.User = c.User.Name
	_, err = model.PersistentVolumeClaimModel.Add(&pvc)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(pvc)
}

// swagger:route GET /api/v1/apps/{appid}/persistentvolumeclaims/{id} pvc reqGetPersistentVolumeClaim
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimController) Get() {
	id := c.GetIDFromURL()

	pvc, err := model.PersistentVolumeClaimModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(pvc)
}

// swagger:route PUT /api/v1/apps/{appid}/persistentvolumeclaims/{id} pvc reqUpdatePersistentVolumeClaim
// update the PersistentVolumeClaim
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimController) Update() {
	id := c.GetIDFromURL()
	var pvc model.PersistentVolumeClaim
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pvc)
	if err != nil {
		c.HandleError(err)
		return
	}

	pvc.Id = int64(id)
	err = model.PersistentVolumeClaimModel.UpdateById(&pvc)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(pvc)
}

// swagger:route PUT /api/v1/apps/{appid}/persistentvolumeclaims/updateorders pvc reqUpdateOrdersPersistentVolumeClaim
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimController) UpdateOrders() {
	var persistentVolumeClaims []*model.PersistentVolumeClaim
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &persistentVolumeClaims)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.PersistentVolumeClaimModel.UpdateOrders(persistentVolumeClaims)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/persistentvolumeclaims/{id} pvc reqDeletePersistentVolumeClaim
// delete the PersistentVolumeClaim
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.PersistentVolumeClaimModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
