package pvc

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type PersistentVolumeClaimTplController struct {
	base.APIController
}

func (c *PersistentVolumeClaimTplController) Prepare() {

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

// swagger:route GET /api/v1/apps/{appid}/persistentvolumeclaims/tpls pvc reqListPersistentVolumeClaimTpl
// get all PersistentVolumeClaimTemplate
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimTplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	isOnline := c.GetIsOnlineFromQuery()

	pvcId := c.Input().Get("pvcId")
	if pvcId != "" {
		param.Query["persistent_volume_claim_id"] = pvcId
	}

	var tpls []model.PersistentVolumeClaimTemplate
	total, err := model.ListTemplate(&tpls, param, model.TableNamePersistentVolumeClaimTemplate, model.PublishTypePersistentVolumeClaim, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range tpls {
		tpls[index].PersistentVolumeClaimId = tpl.PersistentVolumeClaim.Id
	}

	c.Success(param.NewPage(total, tpls))
}

// swagger:route POST /api/v1/apps/{appid}/persistentvolumeclaims/tpls pvc reqCreatePersistentVolumeClaimTpl
// create PersistentVolumeClaimTemplate
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimTplController) Create() {
	var tpl model.PersistentVolumeClaimTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validPersistentVolumeClaimTemplate(tpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	tpl.User = c.User.Name
	_, err = model.PersistentVolumeClaimTplModel.Add(&tpl)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(tpl)
}

func validPersistentVolumeClaimTemplate(templateStr string) error {
	tpl := v1.PersistentVolumeClaim{}
	err := json.Unmarshal(hack.Slice(templateStr), &tpl)
	if err != nil {
		return fmt.Errorf("tpl template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/persistentvolumeclaims/tpls/{id} pvc reqGetPersistentVolumeClaimTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimTplController) Get() {
	id := c.GetIDFromURL()

	tpl, err := model.PersistentVolumeClaimTplModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(tpl)
}

// swagger:route PUT /api/v1/apps/{appid}/persistentvolumeclaims/tpls/{id} pvc reqUpdatePersistentVolumeClaimTpl
// update the PersistentVolumeClaimTemplate
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimTplController) Update() {
	id := c.GetIDFromURL()
	var tpl model.PersistentVolumeClaimTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validPersistentVolumeClaimTemplate(tpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	tpl.Id = int64(id)
	err = model.PersistentVolumeClaimTplModel.UpdateById(&tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(tpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/persistentvolumeclaims/tpls/{id} pvc reqDeletePersistentVolumeClaimTpl
// delete the PersistentVolumeClaimTemplate
// responses:
//
//	200: respSuccessDescription
func (c *PersistentVolumeClaimTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.PersistentVolumeClaimTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
