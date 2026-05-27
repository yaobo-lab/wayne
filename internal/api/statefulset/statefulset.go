package statefulset

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type StatefulsetController struct {
	base.APIController
}

func (c *StatefulsetController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeStatefulset, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/statefulsets/names statefulset reqGetNamesStatefulset
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	statefulsets, err := model.StatefulsetModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(statefulsets)
}

// swagger:route GET /api/v1/apps/{appid}/statefulsets statefulset reqListStatefulset
// get all Statefulset
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	var statefulsets []model.Statefulset
	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	}

	total, err := model.GetTotal(new(model.Statefulset), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.Statefulset), &statefulsets, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range statefulsets {
		statefulsets[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, statefulsets))
	return
}

// swagger:route POST /api/v1/apps/{appid}/statefulsets statefulset reqCreateStatefulset
// create Statefulset
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetController) Create() {
	var statefulset model.Statefulset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &statefulset)
	if err != nil {
		c.HandleError(err)
		return
	}

	statefulset.User = c.User.Name
	_, err = model.StatefulsetModel.Add(&statefulset)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(statefulset)
}

// swagger:route GET /api/v1/apps/{appid}/statefulsets/{id} statefulset reqGetStatefulset
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetController) Get() {
	id := c.GetIDFromURL()

	statefulset, err := model.StatefulsetModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(statefulset)
}

// swagger:route PUT /api/v1/apps/{appid}/statefulsets/{id} statefulset reqUpdateStatefulset
// update the Statefulset
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetController) Update() {
	id := c.GetIDFromURL()

	var statefulset model.Statefulset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &statefulset)
	if err != nil {
		c.HandleError(err)
		return
	}

	statefulset.Id = int64(id)
	err = model.StatefulsetModel.UpdateById(&statefulset)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(statefulset)
}

// swagger:route PUT /api/v1/apps/{appid}/statefulsets/updateorders statefulset reqUpdateOrdersStatefulset
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetController) UpdateOrders() {
	var statefulsets []*model.Statefulset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &statefulsets)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.StatefulsetModel.UpdateOrders(statefulsets)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/statefulsets/{id} statefulset reqDeleteStatefulset
// delete the Statefulset
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.StatefulsetModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
