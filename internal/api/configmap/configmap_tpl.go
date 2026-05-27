package configmap

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type ConfigMapTplController struct {
	base.APIController
}

// 重写 prepare
func (c *ConfigMapTplController) Prepare() {
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
		c.CheckPermission(model.PermissionTypeConfigMap, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/configmaps/tpls configmap reqListConfigMapTpl
// get all ConfigMapTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapTplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	isOnline := c.GetIsOnlineFromQuery()

	configmapId := c.Input().Get("cId")
	if configmapId != "" {
		param.Query["config_map_id"] = configmapId
	}
	var configMapTpls []model.ConfigMapTemplate
	total, err := model.ListTemplate(&configMapTpls, param, model.TableNameConfigMapTemplate, model.PublishTypeConfigMap, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range configMapTpls {
		configMapTpls[index].ConfigMapId = tpl.ConfigMap.Id
	}

	c.Success(param.NewPage(total, configMapTpls))
	return
}

// swagger:route POST /api/v1/apps/{appid}/configmaps/tpls configmap reqCreateConfigMapTpl
// create ConfigMapTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapTplController) Create() {
	var configMapTpl model.ConfigMapTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &configMapTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validConfigMapTemplate(configMapTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	configMapTpl.User = c.User.Name
	_, err = model.ConfigMapTplModel.Add(&configMapTpl)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(configMapTpl)
}

func validConfigMapTemplate(templateStr string) error {
	configMapTpl := v1.ConfigMap{}
	err := json.Unmarshal(hack.Slice(templateStr), &configMapTpl)
	if err != nil {
		return fmt.Errorf("configMapTpl template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/configmaps/tpls/{id} configmap reqGetConfigMapTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapTplController) Get() {
	id := c.GetIDFromURL()

	configMapTpl, err := model.ConfigMapTplModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(configMapTpl)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/configmaps/tpls/{id} configmap reqUpdateConfigMapTpl
// update the ConfigMapTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapTplController) Update() {
	id := c.GetIDFromURL()
	var configMapTpl model.ConfigMapTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &configMapTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validConfigMapTemplate(configMapTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	configMapTpl.Id = int64(id)
	err = model.ConfigMapTplModel.UpdateById(&configMapTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(configMapTpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/configmaps/tpls/{id} configmap reqDeleteConfigMapTpl
// delete the ConfigMapTemplate
// responses:
//
//	200: respSuccessDescription
func (c *ConfigMapTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.ConfigMapTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
