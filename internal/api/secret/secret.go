package secret

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type SecretController struct {
	base.APIController
}

func (c *SecretController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeSecret, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/secrets/names secret reqGetNamesSecret
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *SecretController) GetNames() {
	filters := make(map[string]interface{})
	deleted := c.GetDeleteFromQuery()
	filters["Deleted"] = deleted
	if c.AppId != 0 {
		filters["App__Id"] = c.AppId
	}

	secret, err := model.SecretModel.GetNames(filters)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(secret)
}

// swagger:route GET /api/v1/apps/{appid}/secrets secret reqListSecret
// get all Secret
// responses:
//
//	200: respSuccessDescription
func (c *SecretController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	secret := []model.Secret{}

	if c.AppId != 0 {
		param.Query["App__Id"] = c.AppId
	} else if !c.User.Admin {
		param.Query["App__AppUsers__User__Id__exact"] = c.User.Id
		perName := model.PermissionModel.MergeName(model.PermissionTypeSecret, model.PermissionRead)
		param.Query["App__AppUsers__Group__Permissions__Permission__Name__contains"] = perName
		param.Groupby = []string{"Id"}
	}

	total, err := model.GetTotal(new(model.Secret), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.Secret), &secret, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	for key, one := range secret {
		secret[key].AppId = one.App.Id
	}

	c.Success(param.NewPage(total, secret))
	return
}

// swagger:route POST /api/v1/apps/{appid}/secrets secret reqCreateSecret
// create Secret
// responses:
//
//	200: respSuccessDescription
func (c *SecretController) Create() {
	var secret model.Secret
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &secret)
	if err != nil {
		c.HandleError(err)
		return
	}

	secret.User = c.User.Name
	_, err = model.SecretModel.Add(&secret)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(secret)
}

// swagger:route GET /api/v1/apps/{appid}/secrets/{id} secret reqGetSecret
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *SecretController) Get() {
	id := c.GetIDFromURL()

	secret, err := model.SecretModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(secret)
}

// swagger:route PUT /api/v1/apps/{appid}/secrets/{id} secret reqUpdateSecret
// update the Secret
// responses:
//
//	200: respSuccessDescription
func (c *SecretController) Update() {
	id := c.GetIDFromURL()
	var secret model.Secret
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &secret)
	if err != nil {
		c.HandleError(err)
		return
	}

	secret.Id = int64(id)
	err = model.SecretModel.UpdateById(&secret)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(secret)
}

// swagger:route PUT /api/v1/apps/{appid}/secrets/updateorders secret reqUpdateOrdersSecret
// batch update the Orders
// responses:
//
//	200: respSuccessDescription
func (c *SecretController) UpdateOrders() {
	var secrets []*model.Secret
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &secrets)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.SecretModel.UpdateOrders(secrets)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route DELETE /api/v1/apps/{appid}/secrets/{id} secret reqDeleteSecret
// delete the Secret
// responses:
//
//	200: respSuccessDescription
func (c *SecretController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.SecretModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
