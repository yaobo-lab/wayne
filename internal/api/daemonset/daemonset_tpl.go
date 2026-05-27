package daemonset

import (
	"encoding/json"
	"fmt"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"

	appsV1 "k8s.io/api/apps/v1"
)

type DaemonSetTplController struct {
	base.APIController
}

func (c *DaemonSetTplController) Prepare() {
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
		c.CheckPermission(model.PermissionTypeDaemonSet, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/daemonsets/tpls daemonset reqListDaemonSetTpl
// get all DaemonSetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetTplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	isOnline := c.GetIsOnlineFromQuery()

	daemonSetId := c.Input().Get("daemonSetId")
	if daemonSetId != "" {
		param.Query["daemon_set_id"] = daemonSetId
	}
	var tpls []model.DaemonSetTemplate
	total, err := model.ListTemplate(&tpls, param, model.TableNameDaemonSetTemplate, model.PublishTypeDaemonSet, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range tpls {
		tpls[index].DaemonSetId = tpl.DaemonSet.Id
	}

	c.Success(param.NewPage(total, tpls))
	return
}

// swagger:route POST /api/v1/apps/{appid}/daemonsets/tpls daemonset reqCreateDaemonSetTpl
// create DaemonSetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetTplController) Create() {
	var tpl model.DaemonSetTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validDaemonSetTemplate(tpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	tpl.User = c.User.Name
	_, err = model.DaemonSetTplModel.Add(&tpl)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(tpl)
}

func validDaemonSetTemplate(tpl string) error {
	daemonSet := appsV1.DaemonSet{}
	err := json.Unmarshal(hack.Slice(tpl), &daemonSet)
	if err != nil {
		return fmt.Errorf("daemonSet template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/daemonsets/tpls/{id} daemonset reqGetDaemonSetTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetTplController) Get() {
	id := c.GetIDFromURL()

	tpl, err := model.DaemonSetTplModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(tpl)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/daemonsets/tpls/{id} daemonset reqUpdateDaemonSetTpl
// update the DaemonSetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetTplController) Update() {
	id := c.GetIDFromURL()

	var tpl model.DaemonSetTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validDaemonSetTemplate(tpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	tpl.Id = int64(id)
	err = model.DaemonSetTplModel.UpdateById(&tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(tpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/daemonsets/tpls/{id} daemonset reqDeleteDaemonSetTpl
// delete the DaemonSetTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DaemonSetTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.DaemonSetTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
