package app

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

// 操作App相关资源
type AppController struct {
	base.APIController
}

func (c *AppController) Prepare() {
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
		c.CheckPermission(model.PermissionTypeApp, perAction)
	}
}

// swagger:route GET /api/v1/apps/statistics app reqAppStatisticsApp
// app count statistics
// responses:
//
//	200: respSuccessDescription
func (c *AppController) AppStatistics() {
	param := c.BuildQueryParam()
	total, err := model.GetTotal(new(model.App), param)
	if err != nil {
		c.HandleError(err)
		return
	}
	details, err := model.AppModel.GetAppCountGroupByNamespace()
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(model.AppStatistics{Total: total, Details: details})
}

// swagger:route GET /api/v1/namespaces/{namespaceid}/apps/names app reqGetNamesApp
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *AppController) GetNames() {
	deleted := c.GetDeleteFromQuery()

	apps, err := model.AppModel.GetNames(deleted)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(apps)
}

// swagger:route GET /api/v1/namespaces/{namespaceid}/apps app reqListApp
// get all app
// responses:
//
//	200: respSuccessDescription
func (c *AppController) List() {
	param := c.BuildQueryParam()
	if c.NamespaceId != 0 {
		param.Query["namespace_id"] = c.NamespaceId
	}

	starred := c.GetBoolParamFromQueryWithDefault("starred", false)

	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	total, err := model.AppModel.Count(param, starred, c.User.Id)
	if err != nil {
		c.HandleError(err)
		return
	}

	apps, err := model.AppModel.List(param, starred, c.User.Id)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(param.NewPage(total, apps))
	return
}

// swagger:route POST /api/v1/namespaces/{namespaceid}/apps app reqCreateApp
// create app
// responses:
//
//	200: respSuccessDescription
func (c *AppController) Create() {
	var app model.App
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &app)
	if err != nil {
		c.HandleError(err)
		return
	}

	app.User = c.User.Name
	_, err = model.AppModel.Add(&app)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(app)
}

// swagger:route GET /api/v1/namespaces/{namespaceid}/apps/{id} app reqGetApp
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *AppController) Get() {
	id := c.GetIDFromURL()

	app, err := model.AppModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(app)
	return
}

// swagger:route PUT /api/v1/namespaces/{namespaceid}/apps/{id} app reqUpdateApp
// update the App
// responses:
//
//	200: respSuccessDescription
func (c *AppController) Update() {
	id := c.GetIDFromURL()
	var app model.App
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &app)
	if err != nil {
		c.HandleError(err)
		return
	}

	app.Id = int64(id)
	err = model.AppModel.UpdateById(&app)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(app)
}

// swagger:route DELETE /api/v1/namespaces/{namespaceid}/apps/{id} app reqDeleteApp
// delete the App
// responses:
//
//	200: respSuccessDescription
func (c *AppController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.AppModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
