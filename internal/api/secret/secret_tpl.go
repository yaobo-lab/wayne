package secret

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type SecretTplController struct {
	base.APIController
}

func (c *SecretTplController) Prepare() {

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

// swagger:route GET /api/v1/apps/{appid}/secrets/tpls secret reqListSecretTpl
// get all SecretTemplate
// responses:
//
//	200: respSuccessDescription
func (c *SecretTplController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	isOnline := c.GetIsOnlineFromQuery()

	secretId := c.Input().Get("secretId")
	if secretId != "" {
		param.Query["secret_map_id"] = secretId
	}

	var secretTpls []model.SecretTemplate
	total, err := model.ListTemplate(&secretTpls, param, model.TableNameSecretTemplate, model.PublishTypeSecret, isOnline)
	if err != nil {
		c.HandleError(err)
		return
	}
	for index, tpl := range secretTpls {
		secretTpls[index].SecretId = tpl.Secret.Id
	}

	c.Success(param.NewPage(total, secretTpls))
	return
}

// swagger:route POST /api/v1/apps/{appid}/secrets/tpls secret reqCreateSecretTpl
// create SecretTemplate
// responses:
//
//	200: respSuccessDescription
func (c *SecretTplController) Create() {
	var secretTpl model.SecretTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &secretTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validSecretTemplate(secretTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	secretTpl.User = c.User.Name
	_, err = model.SecretTplModel.Add(&secretTpl)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(secretTpl)
}

func validSecretTemplate(templateStr string) error {
	secret := v1.Secret{}
	err := json.Unmarshal(hack.Slice(templateStr), &secret)
	if err != nil {
		return fmt.Errorf("secretTpl template format error.%v", err.Error())
	}
	return nil
}

// swagger:route GET /api/v1/apps/{appid}/secrets/tpls/{id} secret reqGetSecretTpl
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *SecretTplController) Get() {
	id := c.GetIDFromURL()

	secretTpl, err := model.SecretTplModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(secretTpl)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/secrets/tpls/{id} secret reqUpdateSecretTpl
// update the SecretTemplate
// responses:
//
//	200: respSuccessDescription
func (c *SecretTplController) Update() {
	id := c.GetIDFromURL()
	var secretTpl model.SecretTemplate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &secretTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if err = validSecretTemplate(secretTpl.Template); err != nil {
		c.HandleError(err)
		return
	}

	secretTpl.Id = int64(id)
	err = model.SecretTplModel.UpdateById(&secretTpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(secretTpl)
}

// swagger:route DELETE /api/v1/apps/{appid}/secrets/tpls/{id} secret reqDeleteSecretTpl
// delete the SecretTemplate
// responses:
//
//	200: respSuccessDescription
func (c *SecretTplController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.SecretTplModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
