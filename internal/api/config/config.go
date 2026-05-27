package config

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/adapter"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type ConfigController struct {
	base.APIController
}

// swagger:route GET /api/v1/configs/system config reqListSystemConfig
// get system config
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ConfigController) ListSystem() {
	appConfig, err := beego.AppConfig.GetSection("default")
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(appConfig)
}

// swagger:route GET /api/v1/configs config reqListConfig
// get system config
// responses:
//
//	200: respSuccessDescription
func (c *ConfigController) List() {
	param := c.BuildQueryParam()
	var configs []model.Config

	total, err := model.GetTotal(new(model.Config), param)
	if err != nil {
		c.HandleError(err)
		return
	}
	err = model.GetAll(new(model.Config), &configs, param)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(param.NewPage(total, configs))
}

// swagger:route POST /api/v1/configs config reqCreateConfig
// get system config
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ConfigController) Create() {
	var config model.Config
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &config)
	if err != nil {
		c.HandleError(err)
		return
	}

	id, err := model.ConfigModel.Add(&config)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(id)
}

// swagger:route PUT /api/v1/configs/{id} config reqUpdateConfig
// update the object
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ConfigController) Update() {
	id := c.GetIDFromURL()

	var config model.Config
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &config)
	if err != nil {
		c.HandleError(err)
		return
	}
	config.Id = int64(id)
	err = model.ConfigModel.UpdateById(&config)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(config)
}

// swagger:route GET /api/v1/configs/{id} config reqGetConfig
// find Object by objectid
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ConfigController) Get() {
	id := c.GetIDFromURL()

	config, err := model.ConfigModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(config)
}

// swagger:route DELETE /api/v1/configs/{id} config reqDeleteConfig
// delete the app
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ConfigController) Delete() {
	id := c.GetIDFromURL()

	err := model.ConfigModel.DeleteById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
