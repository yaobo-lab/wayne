package apikey

import (
	"encoding/json"
	"time"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/dgrijalva/jwt-go"

	"wayne/internal/api/base"
	"wayne/internal/model"
	rsakey "wayne/pkg"
)

type ApiKeyController struct {
	base.APIController
}

func (c *ApiKeyController) Prepare() {

	c.APIController.Prepare()
	if beego.AppConfig.String("EnableApiKeys") != "true" {
		c.AbortForbidden("APIKey is not enabled. Please contact the administrator.")
	}

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
		c.CheckPermission(model.PermissionTypeAPIKey, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/apikeys apikey reqListAppApiKey
// get all
// responses:
//   200: respSuccessDescription

// swagger:route GET /api/v1/namespaces/{namespaceid}/apikeys apikey reqListNamespaceApiKey
// get all
// responses:
//
//	200: respSuccessDescription
func (c *ApiKeyController) List() {
	param := c.BuildQueryParam()

	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	apiKeyType := c.Input().Get("type")
	if apiKeyType != "" {
		param.Query["type"] = apiKeyType
	}

	resourceId := c.Input().Get("resourceId")
	if resourceId != "" {
		param.Query["ResourceId"] = resourceId
	}

	var apiKeys []model.APIKey

	param.Relate = "Group"
	total, err := model.GetTotal(new(model.APIKey), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.GetAll(new(model.APIKey), &apiKeys, param)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(param.NewPage(total, apiKeys))
	return
}

// swagger:route GET /api/v1/apps/{appid}/apikeys/{id} apikey reqGetAppApiKey
// find Object by id
// responses:
//   200: respSuccessDescription

// swagger:route GET /api/v1/namespaces/{namespaceid}/apikeys/{id} apikey reqGetNamespaceApiKey
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *ApiKeyController) Get() {
	id := c.GetIDFromURL()
	app, err := model.ApiKeyModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(app)
	return
}

// swagger:route PUT /api/v1/apps/{appid}/apikeys/{id} apikey reqUpdateAppApiKey
// update the App
// responses:
//   200: respSuccessDescription

// swagger:route PUT /api/v1/namespaces/{namespaceid}/apikeys/{id} apikey reqUpdateNamespaceApiKey
// update the App
// responses:
//
//	200: respSuccessDescription
func (c *ApiKeyController) Update() {
	id := c.GetIDFromURL()
	var apiKey model.APIKey
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &apiKey)
	if err != nil {
		c.HandleError(err)
		return
	}

	apiKey.Id = int64(id)
	err = model.ApiKeyModel.UpdateById(&apiKey)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(apiKey)
}

// swagger:route POST /api/v1/apps/{appid}/apikeys apikey reqCreateAppApiKey
// create APIKey
// responses:
//   200: respSuccessDescription

// swagger:route POST /api/v1/namespaces/{namespaceid}/apikeys apikey reqCreateNamespaceApiKey
// create APIKey
// responses:
//
//	200: respSuccessDescription
func (c *ApiKeyController) Create() {
	var apiKey model.APIKey
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &apiKey)
	if err != nil {
		c.HandleError(err)
		return
	}

	apiKey.User = c.User.Name
	_, err = model.ApiKeyModel.Add(&apiKey)

	if err != nil {
		c.HandleError(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		// 签发者
		"iss": "wayne",
		// 签发时间
		"iat": apiKey.CreateTime.Unix(),
		"exp": apiKey.CreateTime.Add(time.Duration(apiKey.ExpireIn) * time.Second).Unix(),
		"aud": apiKey.Id,
	})

	apiToken, err := token.SignedString(rsakey.RsaPrivateKey)
	if err != nil {
		c.HandleError(err)
		return
	}

	apiKey.Token = apiToken

	err = model.ApiKeyModel.UpdateById(&apiKey)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(apiKey)
}

// swagger:route DELETE /api/v1/apps/{appid}/apikeys/{id} apikey reqDeleteAppApiKey
// delete the APIKey
// responses:
//   200: respSuccessDescription

// swagger:route DELETE /api/v1/namespaces/{namespaceid}/apikeys/{id} apikey reqDeleteNamespaceApiKey
// delete the APIKey
// responses:
//
//	200: respSuccessDescription
func (c *ApiKeyController) Delete() {
	id := c.GetIDFromURL()

	logical := c.GetLogicalFromQuery()

	err := model.ApiKeyModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
