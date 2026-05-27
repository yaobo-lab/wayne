package service

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/model"

	"wayne/pkg/hack"
)

// 服务模版相关操作
type ServiceTplController struct {
	base.APIController
}

func (c *ServiceTplController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Create", c.Create)
	c.Mapping("Get", c.Get)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

func (c *ServiceTplController) Prepare() {

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

// swagger:route GET /api/v1/apps/{appid}/services/tpls service reqListServiceTpl
// get all ServiceTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ServiceTplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	isOnline := c.GetIsOnlineFromQuery()

	serviceId := c.Input().Get("serviceId")
	if serviceId != "" {
		param.Query["service_id"] = serviceId
	}

	var serviceTpls []model.ServiceTemplate
	total, err := model.ListTemplate(&serviceTpls, param, model.TableNameServiceTemplate, model.PublishTypeService, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range serviceTpls {
		serviceTpls[index].ServiceId = tpl.Service.Id
	}

	c.Success(param.NewPage(total, serviceTpls))
}

// swagger:route POST /api/v1/apps/{appid}/services/tpls service reqCreateServiceTpl
// create ServiceTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ServiceTplController) Create() {
	var serviceTpl model.ServiceTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &serviceTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	err = validServiceTemplate(serviceTpl.Template)
	if err != nil {
		c.HandleError(err)
		return
	}

	serviceTpl.User = c.User.Name

	_, err = model.ServiceTplModel.Add(&serviceTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(serviceTpl)
}

func validServiceTemplate(serviceTplStr string) error {
	service := v1.Service{}
	err := json.Unmarshal(hack.Slice(serviceTplStr), &service)
	if err != nil {
		return fmt.Errorf("service template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/services/tpls/{id} service reqGetServiceTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *ServiceTplController) Get() {
	id := c.GetIDFromURL()

	serviceTpl, err := model.ServiceTplModel.GetById(id)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(serviceTpl)
}

// swagger:route PUT /api/v1/apps/{appid}/services/tpls/{id} service reqUpdateServiceTpl
// update the ServiceTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ServiceTplController) Update() {
	id := c.GetIDFromURL()
	var serviceTpl model.ServiceTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &serviceTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validServiceTemplate(serviceTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	serviceTpl.Id = int64(id)
	err = model.ServiceTplModel.UpdateById(&serviceTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(serviceTpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/services/tpls/{id} service reqDeleteServiceTpl
// delete the ServiceTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ServiceTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.ServiceTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
