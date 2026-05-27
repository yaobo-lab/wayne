package permission

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"

	"wayne/pkg/logger"
)

type PermissionController struct {
	base.APIController
}

func (c *PermissionController) Prepare() {

	c.APIController.Prepare()

	perAction := ""
	_, method := c.GetControllerAndAction()
	switch method {
	case "Create":
		perAction = model.PermissionCreate
	case "Update":
		perAction = model.PermissionUpdate
	case "Delete":
		perAction = model.PermissionDelete
	}
	if perAction != "" && !c.User.Admin {
		c.AbortForbidden("operation need admin permission.")
	}
}

// swagger:route GET /api/v1/permissions permission reqListPermission
// get all permission
// responses:
//
//	200: respSuccessDescription
func (c *PermissionController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	total, err := model.GetTotal(new(model.Permission), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	permissions := []model.Permission{}
	err = model.GetAll(new(model.Permission), &permissions, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(param.NewPage(total, permissions))
}

// swagger:route GET /api/v1/permissions/{id} permission reqGetPermission
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *PermissionController) Get() {
	id := c.GetIDFromURL()

	permissions, err := model.PermissionModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(permissions)
}

// swagger:route POST /api/v1/permissions permission reqCreatePermission
// create permission
// responses:
//
//	200: respSuccessDescription
func (c *PermissionController) Create() {
	var permission model.Permission
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &permission)
	if err != nil {
		c.HandleError(err)
		return
	}
	_, err = model.PermissionModel.Add(&permission)

	if err != nil {
		logger.Errorf("create error.%v", err.Error())
		c.HandleError(err)
		return
	}
	c.Success(permission)
}

// swagger:route PUT /api/v1/permissions/{id} permission reqUpdatePermission
// update the permission
// responses:
//
//	200: respSuccessDescription
func (c *PermissionController) Update() {
	var permission *model.Permission
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &permission)
	if err != nil {
		c.HandleError(err)
		return
	}

	id := c.GetIDFromURL()
	permission.Id = int64(id)

	err = model.PermissionModel.UpdateById(permission)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(permission)
}

// swagger:route DELETE /api/v1/permissions/{id} permission reqDeletePermission
// delete the Permission
// responses:
//
//	200: respSuccessDescription
func (c *PermissionController) Delete() {
	id := c.GetIDFromURL()

	err := model.PermissionModel.Delete(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(nil)
}
