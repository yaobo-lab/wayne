package permission

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type GroupController struct {
	base.APIController
}

func (c *GroupController) Prepare() {

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

// swagger:route GET /api/v1/groups permission reqListGroup
// get all group
// responses:
//
//	200: respSuccessDescription
func (c *GroupController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	gtype := c.Input().Get("type")
	if gtype != "" {
		param.Query["type__exact"] = gtype
	}

	total, err := model.GetTotal(new(model.Group), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	groups := []model.Group{}
	err = model.GetAll(new(model.Group), &groups, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(param.NewPage(total, groups))
	return
}

// swagger:route POST /api/v1/groups permission reqCreateGroup
// create group
// responses:
//
//	200: respSuccessDescription
func (c *GroupController) Create() {
	var group model.Group
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err != nil {
		c.HandleError(err)
		return
	}
	_, err = model.GroupModel.AddGroup(&group)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(group)
}

// swagger:route GET /api/v1/groups/{id} permission reqGetGroup
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *GroupController) Get() {
	id := c.GetIDFromURL()

	group, err := model.GroupModel.GetGroupById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(group)
}

// swagger:route GET /api/v1/groups/{id} permission reqUpdateGroup
// update the group
// responses:
//
//	200: respSuccessDescription
func (c *GroupController) Update() {
	var group *model.Group
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GroupModel.UpdateGroupById(group)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(group)
}

// swagger:route DELETE /api/v1/groups/{id} permission reqDeleteGroup
// delete the Group
// responses:
//
//	200: respSuccessDescription
func (c *GroupController) Delete() {
	id := c.GetIDFromURL()

	err := model.GroupModel.DeleteGroup(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(nil)
}
