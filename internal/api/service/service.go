package service

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type ServiceController struct {
	base.APIController
}

func (c *ServiceController) URLMapping() {
	c.Mapping("GetNames", c.GetNames)
	c.Mapping("List", c.List)
	c.Mapping("Create", c.Create)
	c.Mapping("Get", c.Get)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

func (c *ServiceController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeService, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/services/names service reqGetNamesService
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *ServiceController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()

	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	services, err := model.ServiceModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(services)
}

// swagger:route GET /api/v1/apps/{appid}/services service reqListService
// get all Service
// responses:
//
//	200: respSuccessDescription
func (c *ServiceController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	service := []model.Service{}
	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypeService, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.Service), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.Service), &service, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range service {
		service[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, service))
}

// swagger:route POST /api/v1/apps/{appid}/services service reqCreateService
// create Service
// responses:
//
//	200: respSuccessDescription
func (c *ServiceController) Create() {
	var service model.Service
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &service)
	if err != nil {
		c.HandleError(err)
		return
	}

	service.User = c.User.Name
	_, err = model.ServiceModel.Add(&service)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(service)
}

// swagger:route GET /api/v1/apps/{appid}/services/{id} service reqGetService
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *ServiceController) Get() {
	id := c.GetIDFromURL()

	service, err := model.ServiceModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(service)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/services/{id} service reqUpdateService
// update the Service
// responses:
//
//	200: respSuccessDescription
func (c *ServiceController) Update() {
	id := c.GetIDFromURL()
	var service model.Service
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &service)
	if err != nil {
		c.HandleError(err)
		return
	}

	service.Id = int64(id)
	err = model.ServiceModel.UpdateById(&service)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(service)
}

// swagger:route PUT /api/v1/apps/{appid}/services/updateorders service reqUpdateOrdersService
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *ServiceController) UpdateOrders() {
	var services []*model.Service
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &services)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.ServiceModel.UpdateOrders(services)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/services/{id} service reqDeleteService
// delete the Service
// responses:
//
//	200: respSuccessDescription
func (c *ServiceController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()
	err := model.ServiceModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
