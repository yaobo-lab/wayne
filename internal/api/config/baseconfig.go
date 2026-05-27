package config

import (
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/adapter"

	"wayne/internal/api/base"
	"wayne/internal/model"
	util "wayne/pkg"
	common "wayne/pkg/dto"
)

type BaseConfigController struct {
	beego.Controller
}

// swagger:route GET /api/v1/configs/base config reqListBaseBaseConfig
// get base config
// responses:
//
//	200: respSuccessDescription
func (c *BaseConfigController) ListBase() {
	configMap := make(map[string]interface{})
	configMap["enableDBLogin"] = beego.AppConfig.DefaultBool("EnableDBLogin", false)
	configMap["appLabelKey"] = util.AppLabelKey
	configMap["namespaceLabelKey"] = util.NamespaceLabelKey
	configMap["oauth2Login"] = parseAuthEnabled("auth.oauth2")
	configMap["enableApiKeys"] = beego.AppConfig.DefaultBool("EnableApiKeys", false)

	var configs []model.Config
	err := model.GetAll(new(model.Config), &configs, &common.QueryParam{
		PageSize: 1000,
		PageNo:   1,
	})
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	for _, conf := range configs {
		configMap[string(conf.Name)] = conf.Value
	}
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = base.Result{Data: configMap}
	c.ServeJSON()
}

func parseAuthEnabled(name string) bool {
	enabled := false
	enabledSection, err := beego.AppConfig.GetSection(name)
	if err == nil {
		enabledBool, err := strconv.ParseBool(enabledSection["enabled"])
		if err == nil {
			enabled = enabledBool
		}
	}
	return enabled
}
