package daemonset

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type DaemonSetController struct {
	base.APIController
}

func (c *DaemonSetController) Prepare() {
	c.APIController.Prepare()
	perAction := ""
	_, method := c.GetControllerAndAction()
	switch method {
	case "Get", "List", "GetNames":
		perAction = model.PermissionRead
	case "Create":
		perAction = model.PermissionCreate
	case "Update":
		perAction = model.PermissionUpdate
	case "Delete":
		perAction = model.PermissionDelete
	}
	if perAction != "" {
		c.CheckPermission(model.PermissionTypeDaemonSet, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/daemonsets/names daemonset reqGetNamesDaemonSet
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	daemonSets, err := model.DaemonSetModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(daemonSets)
}

// swagger:route GET /api/v1/apps/{appid}/daemonsets daemonset reqListDaemonSet
// get all DaemonSet
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	var daemonSets []model.DaemonSet
	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	}

	total, err := model.GetTotal(new(model.DaemonSet), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.DaemonSet), &daemonSets, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range daemonSets {
		daemonSets[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, daemonSets))
	return
}

// swagger:route POST /api/v1/apps/{appid}/daemonsets daemonset reqCreateDaemonSet
// create DaemonSet
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetController) Create() {
	var daemonSet model.DaemonSet
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &daemonSet)
	if err != nil {
		c.HandleError(err)
		return
	}

	daemonSet.User = c.User.Name
	_, err = model.DaemonSetModel.Add(&daemonSet)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(daemonSet)
}

// swagger:route GET /api/v1/apps/{appid}/daemonsets/{id} daemonset reqGetDaemonSet
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetController) Get() {
	id := c.GetIDFromURL()

	daemonSet, err := model.DaemonSetModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(daemonSet)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/daemonsets/{id} daemonset reqUpdateDaemonSet
// update the DaemonSet
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetController) Update() {
	id := c.GetIDFromURL()

	var daemonSet model.DaemonSet
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &daemonSet)
	if err != nil {
		c.HandleError(err)
		return
	}

	daemonSet.Id = int64(id)
	err = model.DaemonSetModel.UpdateById(&daemonSet)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(daemonSet)
}

// swagger:route PUT /api/v1/apps/{appid}/daemonsets/updateorders daemonset reqUpdateOrdersDaemonSet
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetController) UpdateOrders() {
	var daemonSets []*model.DaemonSet
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &daemonSets)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.DaemonSetModel.UpdateOrders(daemonSets)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/daemonsets/{id} daemonset reqDeleteDaemonSet
// delete the DaemonSet
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.DaemonSetModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
