package hpa

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"

	"wayne/pkg/logger"
)

type HPAController struct {
	base.APIController
}

func (c *HPAController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeHPA, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/hpas/names hpa reqGetNamesHPA
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *HPAController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	hpas, err := model.HPAModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(hpas)
}

// swagger:route GET /api/v1/apps/{appid}/hpas hpa reqListHPA
// get all HPA
// responses:
//
//	200: respSuccessDescription
func (c *HPAController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	hpas := []model.HPA{}
	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypeHPA, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.HPA), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.HPA), &hpas, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range hpas {
		hpas[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, hpas))
}

// swagger:route POST /api/v1/apps/{appid}/hpas hpa reqCreateHPA
// create HPA
// responses:
//
//	200: respSuccessDescription
func (c *HPAController) Create() {
	var hpa model.HPA
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &hpa)
	if err != nil {
		c.HandleError(err)
		return
	}

	hpa.User = c.User.Name
	_, err = model.HPAModel.Add(&hpa)

	if err != nil {
		logger.Errorf("create error.%v", err.Error())
		c.HandleError(err)
		return
	}
	c.Success(hpa)
}

// swagger:route GET /api/v1/apps/{appid}/hpas/{id} hpa reqGetHPA
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *HPAController) Get() {
	id := c.GetIDFromURL()

	hpa, err := model.HPAModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(hpa)
}

// swagger:route PUT /api/v1/apps/{appid}/hpas/{id} hpa reqUpdateHPA
// update the HPA
// responses:
//
//	200: respSuccessDescription
func (c *HPAController) Update() {
	id := c.GetIDFromURL()
	var hpa model.HPA
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &hpa)
	if err != nil {
		c.HandleError(err)
		return
	}

	hpa.Id = int64(id)
	err = model.HPAModel.UpdateById(&hpa)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(hpa)
}

// swagger:route PUT /api/v1/apps/{appid}/hpas/updateorders hpa reqUpdateOrdersHPA
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *HPAController) UpdateOrders() {
	var hpas []*model.HPA
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &hpas)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.HPAModel.UpdateOrders(hpas)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}

// swagger:route DELETE /api/v1/apps/{appid}/hpas/{id} hpa reqDeleteHPA
// delete the HPA
// responses:
//
//	200: respSuccessDescription
func (c *HPAController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.HPAModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
