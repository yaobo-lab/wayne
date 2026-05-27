package statefulset

import (
	"encoding/json"
	"fmt"

	appV1 "k8s.io/api/apps/v1"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type StatefulsetTplController struct {
	base.APIController
}

func (c *StatefulsetTplController) Prepare() {

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

// swagger:route GET /api/v1/apps/{appid}/statefulsets/tpls statefulset reqListStatefulsetTpl
// get all StatefulsetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetTplController) List() {
	param := c.BuildQueryParam()

	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	isOnline := c.GetIsOnlineFromQuery()

	statefulsetId := c.Input().Get("statefulsetId")
	if statefulsetId != "" {
		param.Query["statefulset_id"] = statefulsetId
	}
	var tpls []model.StatefulsetTemplate
	total, err := model.ListTemplate(&tpls, param, model.TableNameStatefulsetTemplate, model.PublishTypeStatefulSet, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range tpls {
		tpls[index].StatefulsetId = tpl.Statefulset.Id
	}

	c.Success(param.NewPage(total, tpls))
}

// swagger:route POST /api/v1/apps/{appid}/statefulsets/tpls statefulset reqCreateStatefulsetTpl
// create StatefulsetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetTplController) Create() {
	var statefulsetTpl model.StatefulsetTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &statefulsetTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validStatefulsetTemplate(statefulsetTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	statefulsetTpl.User = c.User.Name
	_, err = model.StatefulsetTplModel.Add(&statefulsetTpl)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(statefulsetTpl)
}

func validStatefulsetTemplate(tpl string) error {
	statefulset := appV1.StatefulSet{}
	err := json.Unmarshal(hack.Slice(tpl), &statefulset)
	if err != nil {
		return fmt.Errorf("statefulset template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/statefulsets/tpls/{id} statefulset reqGetStatefulsetTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetTplController) Get() {
	id := c.GetIDFromURL()

	statefulsetTpl, err := model.StatefulsetTplModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(statefulsetTpl)
}

// swagger:route PUT /api/v1/apps/{appid}/statefulsets/tpls/{id} statefulset reqUpdateStatefulsetTpl
// update the StatefulsetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetTplController) Update() {
	id := c.GetIDFromURL()

	var statefulsetTpl model.StatefulsetTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &statefulsetTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validStatefulsetTemplate(statefulsetTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	statefulsetTpl.Id = int64(id)
	err = model.StatefulsetTplModel.UpdateById(&statefulsetTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(statefulsetTpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/statefulsets/tpls/{id} statefulset reqDeleteStatefulsetTpl
// delete the StatefulsetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *StatefulsetTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.StatefulsetTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
