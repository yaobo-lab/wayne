package cronjob

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type CronjobController struct {
	base.APIController
}

func (c *CronjobController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeCronjob, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/cronjobs/names cronjob reqGetNamesCronjob
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *CronjobController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	cronjobs, err := model.CronjobModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(cronjobs)
}

// swagger:route GET /api/v1/apps/{appid}/cronjobs cronjob reqListCronjob
// get all Cronjob
// responses:
//
//	200: respSuccessDescription
func (c *CronjobController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	cronjob := []model.Cronjob{}

	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypeCronjob, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.Cronjob), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.Cronjob), &cronjob, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range cronjob {
		cronjob[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, cronjob))
	return
}

// swagger:route POST /api/v1/apps/{appid}/cronjobs cronjob reqCreateCronjob
// create Cronjob
// responses:
//
//	200: respSuccessDescription
func (c *CronjobController) Create() {
	var cronjob model.Cronjob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cronjob)
	if err != nil {
		c.HandleError(err)
		return
	}

	cronjob.User = c.User.Name
	_, err = model.CronjobModel.Add(&cronjob)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(cronjob)
}

// swagger:route GET /api/v1/apps/{appid}/cronjobs/{id} cronjob reqGetCronjob
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *CronjobController) Get() {
	id := c.GetIDFromURL()

	cronjob, err := model.CronjobModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(cronjob)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/cronjobs/{id} cronjob reqUpdateCronjob
// update the Cronjob
// responses:
//
//	200: respSuccessDescription
func (c *CronjobController) Update() {
	id := c.GetIDFromURL()
	var cronjob model.Cronjob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cronjob)
	if err != nil {
		c.HandleError(err)
		return
	}

	cronjob.Id = int64(id)
	err = model.CronjobModel.UpdateById(&cronjob)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(cronjob)
}

// swagger:route PUT /api/v1/apps/{appid}/cronjobs/updateorders cronjob reqUpdateOrdersCronjob
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *CronjobController) UpdateOrders() {
	var cronjobs []*model.Cronjob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cronjobs)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.CronjobModel.UpdateOrders(cronjobs)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/cronjobs/{id} cronjob reqDeleteCronjob
// delete the Cronjob
// responses:
//
//	200: respSuccessDescription
func (c *CronjobController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.CronjobModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
