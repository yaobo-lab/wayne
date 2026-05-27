package configmap

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type ConfigMapController struct {
	base.APIController
}

func (c *ConfigMapController) Prepare() {
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
		c.CheckPermission(model.PermissionTypeConfigMap, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/configmaps/names configmap reqGetNamesConfigMap
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	configMaps, err := model.ConfigMapModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(configMaps)
}

// swagger:route GET /api/v1/apps/{appid}/configmaps configmap reqListConfigMap
// get all ConfigMap
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	configMap := []model.ConfigMap{}

	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypeConfigMap, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.ConfigMap), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.ConfigMap), &configMap, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range configMap {
		configMap[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, configMap))
	return
}

// swagger:route POST /api/v1/apps/{appid}/configmaps configmap reqCreateConfigMap
// create ConfigMap
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapController) Create() {
	var configmap model.ConfigMap
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &configmap)
	if err != nil {
		c.HandleError(err)
		return
	}

	configmap.User = c.User.Name
	_, err = model.ConfigMapModel.Add(&configmap)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(configmap)
}

// swagger:route GET /api/v1/apps/{appid}/configmaps/{id} configmap reqGetConfigMap
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapController) Get() {
	id := c.GetIDFromURL()

	configmap, err := model.ConfigMapModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(configmap)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/configmaps/{id} configmap reqUpdateConfigMap
// update the ConfigMap
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapController) Update() {
	id := c.GetIDFromURL()
	var configmap model.ConfigMap
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &configmap)
	if err != nil {
		c.HandleError(err)
		return
	}

	configmap.Id = int64(id)
	err = model.ConfigMapModel.UpdateById(&configmap)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(configmap)
}

// swagger:route PUT /api/v1/apps/{appid}/configmaps/updateorders configmap reqUpdateOrdersConfigMap
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapController) UpdateOrders() {
	var configMaps []*model.ConfigMap
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &configMaps)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.ConfigMapModel.UpdateOrders(configMaps)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/configmaps/{id} configmap reqDeleteConfigMap
// delete the ConfigMap
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.ConfigMapModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
