package hpa

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/hpa"
	"wayne/internal/model"

	autoscaling "k8s.io/api/autoscaling/v1"
)

type KubeHPAController struct {
	base.APIController
}

func (c *KubeHPAController) Prepare() {
	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"Create": model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeHorizontalPodAutoscaler)
}

// swagger:route POST /api/v1/kubernetes/apps/{appid}/hpas/{hpaId}/tpls/{tplId}/clusters/{cluster} hpa reqCreateKubeHPA
// responses:
//
//	200: respSuccessDescription
func (c *KubeHPAController) Create() {
	hpaId := c.GetIntParamFromURL(":hpaId")
	tplId := c.GetIntParamFromURL(":tplId")
	var kubeHPA autoscaling.HorizontalPodAutoscaler
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &kubeHPA)
	if err != nil {
		c.HandleError(err)
		return
	}
	clusterName := c.Ctx.Input.Param(":cluster")
	k8sClient := c.Client(clusterName)

	_, err = hpa.CreateOrUpdateHPA(k8sClient, &kubeHPA)
	if err != nil {
		c.HandleError(err)
		return
	}

	publishStatus := model.PublishStatus{
		ResourceId: int64(hpaId),
		TemplateId: int64(tplId),
		Type:       model.PublishTypeHPA,
		Cluster:    clusterName,
	}
	err = model.PublishStatusModel.Publish(&publishStatus)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
