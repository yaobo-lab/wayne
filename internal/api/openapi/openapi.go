package openapi

import (
	"fmt"
	"net/http"

	"crypto/rsa"
	"wayne/internal/api/base"
	"wayne/internal/model"
)

const (
	GetPodInfoAction             = "GET_POD_INFO"
	GetPodInfoFromIPAction       = "GET_POD_INFO_FROM_IP"
	GetResourceInfoAction        = "GET_RESOURCE_INFO"
	GetDeploymentStatusAction    = "GET_DEPLOYMENT_STATUS"
	UpgradeDeploymentAction      = "UPGRADE_DEPLOYMENT"
	ScaleDeploymentAction        = "SCALE_DEPLOYMENT"
	RestartDeploymentAction      = "RESTART_DEPLOYMENT"
	GetDeploymentDetailAction    = "GET_DEPLOYMENT_DETAIL"
	GetLatestDeploymentTplAction = "GET_LATEST_DEPLOYMENT_TPL"
	GetPodListAction             = "GET_POD_LIST"

	ListNamespaceUsers = "LIST_NAMESPACE_USERS"
	ListNamespaceApps  = "LIST_NAMESPACE_APPS"
	ListAppDeploys     = "List_APP_DEPLOYS"

	PermissionPrefix = "OPENAPI_"
)

var (
	RsaPrivateKey *rsa.PrivateKey
	RsaPublicKey  *rsa.PublicKey
)

type OpenAPIController struct {
	base.APIKeyController
}

func (c *OpenAPIController) Prepare() {
	c.APIKeyController.Prepare()
}

func (c *OpenAPIController) CheckoutRoutePermission(action string) bool {
	permission := false
	for _, p := range c.APIKey.Group.Permissions {
		if p.Name == PermissionPrefix+action {
			permission = true
		}
	}
	if !permission {
		c.AddErrorAndResponse(fmt.Sprintf("APIKey does not have the following permission: %s", PermissionPrefix+action), http.StatusUnauthorized)
		return false
	}
	return true
}

func (c *OpenAPIController) CheckDeploymentPermission(deployment string) bool {
	if c.APIKey.Type == model.NamespaceAPIKey {
		d, err := model.DeploymentModel.GetByName(deployment)
		if err != nil {
			c.AddErrorAndResponse(err.Error(), http.StatusBadRequest)
			return false
		}
		app, _ := model.AppModel.GetById(d.AppId)
		if app.Namespace.Id != c.APIKey.ResourceId {
			c.AddErrorAndResponse(fmt.Sprintf("APIKey does not have permission to operate request resource: %s", deployment), http.StatusUnauthorized)
			return false
		}
	}
	if c.APIKey.Type == model.ApplicationAPIKey {
		deploy, err := model.DeploymentModel.GetByName(deployment)
		if err != nil {
			c.AddErrorAndResponse(err.Error(), http.StatusBadRequest)
			return false
		}
		if deploy.AppId != c.APIKey.ResourceId {
			c.AddErrorAndResponse(fmt.Sprintf("APIKey does not have permission to operate request resource(deployment): %s", deployment), http.StatusUnauthorized)
			return false
		}
	}
	return true
}

func (c *OpenAPIController) CheckNamespacePermission(namespace string) bool {
	if c.APIKey.Type == model.NamespaceAPIKey {
		ns, err := model.NamespaceModel.GetByName(namespace)
		if err != nil {
			c.AddErrorAndResponse(err.Error(), http.StatusBadRequest)
			return false
		}
		if ns.Deleted == true {
			c.AddErrorAndResponse(fmt.Sprintf("The requested namespace has been offline: %s", namespace), http.StatusBadRequest)
			return false
		}
		if ns.Id != c.APIKey.ResourceId {
			c.AddErrorAndResponse(fmt.Sprintf("APIKey does not have permission to operate request resource(namespace): %s", namespace), http.StatusUnauthorized)
			return false
		}
	}
	return true
}
