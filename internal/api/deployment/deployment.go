package deployment

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type DeploymentController struct {
	base.APIController
}

func (c *DeploymentController) Prepare() {
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
		c.CheckPermission(model.PermissionTypeDeployment, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/deployments/names deployment reqGetNamesDeployment
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	deployments, err := model.DeploymentModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(deployments)
}

// swagger:route GET /api/v1/apps/{appid}/deployments deployment reqListDeployment
// get all Deployment
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	deployment := []model.Deployment{}
	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypeDeployment, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.Deployment), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.Deployment), &deployment, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range deployment {
		deployment[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, deployment))
	return
}

// swagger:route POST /api/v1/apps/{appid}/deployments deployment reqCreateDeployment
// create Deployment
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentController) Create() {
	var deploy model.Deployment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploy)
	if err != nil {
		c.HandleError(err)
		return
	}

	deploy.User = c.User.Name
	_, err = model.DeploymentModel.Add(&deploy)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(deploy)
}

// swagger:route GET /api/v1/apps/{appid}/deployments/{id} deployment reqGetDeployment
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentController) Get() {
	id := c.GetIDFromURL()

	deploy, err := model.DeploymentModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(deploy)
}

// swagger:route PUT /api/v1/apps/{appid}/deployments/{id} deployment reqUpdateDeployment
// update the Deployment
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentController) Update() {
	id := c.GetIDFromURL()

	var deploy model.Deployment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploy)
	if err != nil {
		c.HandleError(err)
		return
	}

	deploy.Id = int64(id)
	err = model.DeploymentModel.UpdateById(&deploy)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(deploy)
}

// swagger:route PUT /api/v1/apps/{appid}/deployments/updateorders deployment reqUpdateOrdersDeployment
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentController) UpdateOrders() {
	var deploys []*model.Deployment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploys)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.DeploymentModel.UpdateOrders(deploys)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/deployments/{id} deployment reqDeleteDeployment
// delete the Deployment
// responses:
//
//	200: respSuccessDescription
func (c *DeploymentController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()
	err := model.DeploymentModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
