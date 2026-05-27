package cronjob

import (
	"encoding/json"
	"fmt"

	"k8s.io/api/batch/v2alpha1"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type CronjobTplController struct {
	base.APIController
}

func (c *CronjobTplController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeCronjob, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/cronjobs/tpls cronjob reqListCronjobTpl
// get all CronjobTemplate
// responses:
//
//	200: respSuccessDescription
func (c *CronjobTplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	isOnline := c.GetIsOnlineFromQuery()

	cronjobId := c.Input().Get("cId")
	if cronjobId != "" {
		param.Query["cronjob_id"] = cronjobId
	}

	var cronjobTpls []model.CronjobTemplate
	total, err := model.ListTemplate(&cronjobTpls, param, model.TableNameCronjobTemplate, model.PublishTypeCronJob, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range cronjobTpls {
		cronjobTpls[index].CronjobId = tpl.Cronjob.Id
	}

	c.Success(param.NewPage(total, cronjobTpls))
	return
}

// swagger:route POST /api/v1/apps/{appid}/cronjobs/tpls cronjob reqCreateCronjobTpl
// create CronjobTemplate
// responses:
//
//	200: respSuccessDescription
func (c *CronjobTplController) Create() {
	var cronjobTpl model.CronjobTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cronjobTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validCronjobTemplate(cronjobTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	cronjobTpl.User = c.User.Name
	_, err = model.CronjobTplModel.Add(&cronjobTpl)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(cronjobTpl)
}

func validCronjobTemplate(templateStr string) error {
	cronjobTpl := v2alpha1.CronJob{}
	err := json.Unmarshal(hack.Slice(templateStr), &cronjobTpl)
	if err != nil {
		return fmt.Errorf("cronjobTpl template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/cronjobs/tpls/{id} cronjob reqGetCronjobTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *CronjobTplController) Get() {
	id := c.GetIDFromURL()

	cronjobTpl, err := model.CronjobTplModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(cronjobTpl)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/cronjobs/tpls/{id} cronjob reqUpdateCronjobTpl
// update the CronjobTemplate
// responses:
//
//	200: respSuccessDescription
func (c *CronjobTplController) Update() {
	id := c.GetIDFromURL()
	var cronjobTpl model.CronjobTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cronjobTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validCronjobTemplate(cronjobTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	cronjobTpl.Id = int64(id)
	err = model.CronjobTplModel.UpdateById(&cronjobTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(cronjobTpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/cronjobs/tpls/{id} cronjob reqDeleteCronjobTpl
// delete the CronjobTemplate
// responses:
//
//	200: respSuccessDescription
func (c *CronjobTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.CronjobTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
