package hpa

import (
	"encoding/json"
	"fmt"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"

	v1 "k8s.io/api/autoscaling/v1"
)

type HPATplController struct {
	base.APIController
}

func (c *HPATplController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeStatefulset, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/hpas/tpls hpa reqListHPATpl
// get all HPATemplate
// responses:
//
//	200: respSuccessDescription
func (c *HPATplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	isOnline := c.GetIsOnlineFromQuery()

	hpaId := c.Input().Get("hpaId")
	if hpaId != "" {
		param.Query["hpa_id"] = hpaId
	}

	var hpaTpls []model.HPATemplate
	total, err := model.ListTemplate(&hpaTpls, param, model.TableNameHPATemplate, model.PublishTypeHPA, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range hpaTpls {
		hpaTpls[index].HPAId = tpl.HPA.Id
	}

	c.Success(param.NewPage(total, hpaTpls))
}

// swagger:route POST /api/v1/apps/{appid}/hpas/tpls hpa reqCreateHPATpl
// create HPATemplate
// responses:
//
//	200: respSuccessDescription
func (c *HPATplController) Create() {
	var hpaTpl model.HPATemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &hpaTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	err = validHPATemplate(hpaTpl.Template)
	if err != nil {
		c.HandleError(err)
		return
	}

	hpaTpl.User = c.User.Name

	_, err = model.HPATemplateModel.Add(&hpaTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(hpaTpl)
}

func validHPATemplate(hpaTplStr string) error {
	hpa := v1.HorizontalPodAutoscaler{}
	err := json.Unmarshal(hack.Slice(hpaTplStr), &hpa)
	if err != nil {
		return fmt.Errorf("hpa template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/hpas/tpls/{id} hpa reqGetHPATpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *HPATplController) Get() {
	id := c.GetIDFromURL()

	tpl, err := model.HPATemplateModel.GetById(id)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(tpl)
}

// swagger:route PUT /api/v1/apps/{appid}/hpas/tpls/{id} hpa reqUpdateHPATpl
// update the HPATemplate
// responses:
//
//	200: respSuccessDescription
func (c *HPATplController) Update() {
	id := c.GetIDFromURL()
	var tpl model.HPATemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validHPATemplate(tpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	tpl.Id = int64(id)
	err = model.HPATemplateModel.UpdateById(&tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(tpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/hpas/tpls/{id} hpa reqDeleteHPATpl
// delete the HPATemplate
// responses:
//
//	200: respSuccessDescription
func (c *HPATplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.HPATemplateModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
