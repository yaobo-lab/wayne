package event

import (
	"fmt"
	"net/http"

	"wayne/internal/api/base"
	api "wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/event"
	"wayne/internal/model"
	"wayne/pkg/dto"
	common "wayne/pkg/dto"
)

type KubeEventController struct {
	base.APIController
}

func (c *KubeEventController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"List": model.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubePod)
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/events/namespaces/{namespace}/clusters/{cluster} events reqListKubeEvent
// Get Pod Event by resource type and name
// responses:
//
//	200: respSuccessDescription
func (c *KubeEventController) List() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	resourceType := c.Input().Get("type")
	resourceName := c.Input().Get("name")
	param := c.BuildKubernetesQueryParam()
	manager := c.Manager(cluster)
	var result *common.Page
	var err error
	switch resourceType {
	case api.ResourceNameCronJob:
		result, err = event.GetPodsEventByCronJobPage(manager.KubeClient, namespace, resourceName, param)
	default:
		err = &dto.ErrorResult{
			Code: http.StatusBadRequest,
			Msg:  fmt.Sprintf("Unsupported resource type (%s). ", resourceType),
		}
	}
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}
