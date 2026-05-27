package deployment

import (
	"encoding/json"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type DeploymentTplController struct {
	base.APIController
}

func (c *DeploymentTplController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeDeployment, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/deployments/tpls deployment reqListDeploymentTpl
// get all DeploymentTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentTplController) List() {
	param := c.BuildQueryParam()

	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	isOnline := c.GetIsOnlineFromQuery()

	deploymentId := c.Input().Get("deploymentId")
	if deploymentId != "" {
		param.Query["deployment_id"] = deploymentId
	}
	var deploymentTpls []model.DeploymentTemplate
	total, err := model.ListTemplate(&deploymentTpls, param, model.TableNameDeploymentTemplate, model.PublishTypeDeployment, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range deploymentTpls {
		deploymentTpls[index].DeploymentId = tpl.Deployment.Id
	}

	c.Success(param.NewPage(total, deploymentTpls))
	return
}

// swagger:route POST /api/v1/apps/{appid}/deployments/tpls deployment reqCreateDeploymentTpl
// create DeploymentTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentTplController) Create() {
	var deployTpl model.DeploymentTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &deployTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validDeploymentTemplate(deployTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	deployTpl.User = c.User.Name
	_, err = model.DeploymentTplModel.Add(&deployTpl)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(deployTpl)
}

func validDeploymentTemplate(deployStr string) error {
	deployment := appsv1.Deployment{}
	err := json.Unmarshal(hack.Slice(deployStr), &deployment)
	if err != nil {
		return fmt.Errorf("deployment template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/deployments/tpls/{id} deployment reqGetDeploymentTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentTplController) Get() {
	id := c.GetIDFromURL()

	deployTpl, err := model.DeploymentTplModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(deployTpl)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/deployments/tpls/{id} deployment reqUpdateDeploymentTpl
// update the DeploymentTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentTplController) Update() {
	id := c.GetIDFromURL()

	var deployTpl model.DeploymentTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &deployTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validDeploymentTemplate(deployTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	deployTpl.Id = int64(id)
	err = model.DeploymentTplModel.UpdateById(&deployTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(deployTpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/deployments/tpls/{id} deployment reqDeleteDeploymentTpl
// delete the DeploymentTemplate
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.DeploymentTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
