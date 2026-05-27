package base

import (
	"fmt"
	"strconv"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"

	"wayne/internal/k8s/client"
	"wayne/internal/model"
)

type APIController struct {
	LoggedInController
	NamespaceId int64
	AppId       int64
}

func (c *APIController) Prepare() {

	c.LoggedInController.Prepare()

	namespaceId, _ := strconv.Atoi(c.Ctx.Input.Param(":namespaceid"))
	if namespaceId < 0 {
		c.HandleError(fmt.Errorf("namespaceId (%d) can't less than 0.", namespaceId))
		return
	}

	c.NamespaceId = int64(namespaceId)

	appId, _ := strconv.Atoi(c.Ctx.Input.Param(":appid"))
	if appId < 0 {
		c.HandleError(fmt.Errorf("appId (%d) can't less than 0.", appId))
		return
	}

	c.AppId = int64(appId)
	if c.NamespaceId == 0 && c.AppId != 0 {
		app, err := model.AppModel.GetById(c.AppId)
		if err != nil {
			c.HandleError(fmt.Errorf("Get app by id error."))
			return
		}
		c.NamespaceId = app.Namespace.Id
	}
}

func (c *APIController) PreparePermission(methodActionMap map[string]string, method string, permissionType string) {
	action, ok := methodActionMap[method]
	if !ok {
		return
	}
	c.CheckPermission(permissionType, action)
}

/*
 * 检查资源权限
 */
func (c *APIController) CheckPermission(perType string, perAction string) {

	// 如果用户是admin，跳过permission检查
	if c.User.Admin {
		return
	}

	if c.NamespaceId <= 0 {
		c.AbortForbidden("Permission error NamespaceId is 0")
		return
	}

	if c.AppId <= 0 {
		c.AbortForbidden("Permission error AppId is 0")
		return
	}

	perName := model.PermissionModel.MergeName(perType, perAction)

	// 检查namespace的操作权限
	_, err := model.NamespaceUserModel.GetOneByPermission(c.NamespaceId, c.User.Id, perName)
	if err != nil {
		c.AbortForbidden(fmt.Sprintf("User (%s) does not have current namespace (%d) permissions (%s).", c.User.Name, c.NamespaceId, perName))
		return
	}

	// 检查App的操作权限
	_, err = model.AppUserModel.GetOneByPermission(c.AppId, c.User.Id, perName)
	if err == nil {
		return
	}
}

func (c *APIController) ApiextensionsClient(cluster string) *clientset.Clientset {
	manager, err := client.Manager(cluster)

	if err != nil {
		c.HandleError(err)
		return nil
	}

	cli, err := clientset.NewForConfig(manager.Config)
	if err != nil {
		c.HandleError(err)
		return nil
	}
	return cli
}

func (c *APIController) Client(cluster string) *kubernetes.Clientset {
	kubeClient, err := client.Client(cluster)
	if err != nil {
		c.HandleError(err)
		return nil
	}
	return kubeClient
}

func (c *APIController) Manager(cluster string) *client.ClusterManager {
	kubeManager, err := client.Manager(cluster)
	if err != nil {
		c.HandleError(err)
		return nil
	}
	return kubeManager
}

func (c *APIController) KubeClient(cluster string) client.ResourceHandler {
	kubeManager, err := client.Manager(cluster)
	if err != nil {
		c.HandleError(err)
		return nil
	}
	return kubeManager.KubeClient
}

func (c *APIController) Success(data interface{}) {
	c.ResultHandlerController.Success(data)
}

// Abort stops controller handler and show the error data， e.g. Prepare
func (c *APIController) AbortForbidden(msg string) {
	c.ResultHandlerController.AbortForbidden(msg)
}

func (c *APIController) AbortUnauthorized(msg string) {
	c.ResultHandlerController.AbortUnauthorized(msg)
}

// Handle return http code and body normally, need return
func (c *APIController) HandleError(err error) {
	c.ResultHandlerController.AbortInternalServerError(err.Error())
}
