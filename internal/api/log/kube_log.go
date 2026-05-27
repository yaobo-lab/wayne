package log

import (
	v1 "k8s.io/api/core/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/log"
	"wayne/internal/model"
	"wayne/pkg/hack"
)

type KubeLogController struct {
	base.APIController
}

func (c *KubeLogController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"List": model.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubePod)
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/podlogs/{pod}/containers/{container}/namespaces/{namespace}/clusters/{cluster} log reqListKubeLog
// pod logs
// responses:
//
//	200: respSuccessDescription
func (c *KubeLogController) List() {
	tailLines := c.GetIntParamFromQuery("tailLines")
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	pod := c.Ctx.Input.Param(":pod")
	container := c.Ctx.Input.Param(":container")
	opt := &v1.PodLogOptions{
		Container: container,
		TailLines: &tailLines,
	}
	cli := c.Client(cluster)

	result, err := log.GetLogsByPod(cli, namespace, pod, opt)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(hack.String(result))
}
