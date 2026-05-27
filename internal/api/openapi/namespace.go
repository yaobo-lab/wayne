package openapi

import (
	"net/http"

	"wayne/internal/model"
	response "wayne/internal/model/dto"
	"wayne/pkg/dto"
)

// resource info include app info and users info.
// swagger:response respresourceinfos
type respResourceInfos struct {
	// in: body
	// Required: true
	Body struct {
		dto.ResponseBase
		Resources []resourceInfo `json:"resources"`
	}
}

// swagger:route GET /openapi/v1/gateway/action/list_namespace_apps openapi ResourceInfoParam
//
// 通过给定的namespace，获取 app 信息和用户信息
//
// 因为查询范围是指定namespace，因此需要绑定 namespace apikey 使用。
//
//	Responses:
//	  200: respresourceinfo
//	  400: responseState
//	  500: responseState
func (c *OpenAPIController) ListNamespaceApps() {
	if !c.CheckoutRoutePermission(ListNamespaceApps) {
		return
	}
	ns := c.GetString("namespace")
	if ns == "" {
		c.AddErrorAndResponse("Invalid namespace parameter!", http.StatusBadRequest)
		return
	}

	namespace, err := model.NamespaceModel.GetByNameAndDeleted(ns, false)
	if err != nil {
		c.AddErrorAndResponse("No namespace exists!", http.StatusBadRequest)
		return
	}
	apps, err := model.AppModel.GetAppsByNamespaceId(namespace.Id, false)
	if err != nil {
		c.AddErrorAndResponse("Failed to get apps by namespace id!", http.StatusBadRequest)
		return
	}

	resp := new(respResourceInfos)
	for _, app := range apps {
		users := make(map[string]*response.User)

		appUsers, err := model.AppUserModel.GetUserListByAppId(app.Id)
		if err != nil {
			c.AddErrorAndResponse("Failed to get appUser list by app id!", http.StatusBadRequest)
			return
		}

		for _, usr := range appUsers {
			if users[usr.User.Name] == nil {
				u := response.User{
					Name:    usr.User.Name,
					Email:   usr.User.Email,
					Display: usr.User.Display,
					Roles:   []string{usr.Group.Name},
				}
				users[usr.User.Name] = &u
			} else {
				users[usr.User.Name].Roles = append(users[usr.User.Name].Roles, usr.Group.Name)
			}
		}

		resp.Body.Resources = append(resp.Body.Resources, resourceInfo{
			App: response.App{
				Id:          app.Id,
				Name:        app.Name,
				Namespace:   app.Namespace.Name,
				Description: app.Description,
				User:        app.User,
				Deleted:     app.Deleted,
				CreateTime:  app.CreateTime,
				UpdateTime:  app.UpdateTime,
			},
			Users: users,
		})
	}

	resp.Body.Code = http.StatusOK
	c.HandleResponse(resp.Body)
}
