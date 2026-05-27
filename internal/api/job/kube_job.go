package job

import (
	batchv1 "k8s.io/api/batch/v1"

	"wayne/internal/api/base"
	"wayne/internal/k8s/kind/job"
	"wayne/internal/model"
)

type KubeJobController struct {
	base.APIController
}

type ClusterJob struct {
	Job     batchv1.Job `json:"kubeJob"`
	Cluster string      `json:"cluster"`
}

func (c *KubeJobController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"ListJobByCronJob": model.PermissionRead,
		"GetEvent":         model.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeJob)
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/jobs/namespaces/{namespace}/clusters/{cluster} job reqListJobByCronJobKubeJob
// find jobs by cronjob
// responses:
//
//	200: respSuccessDescription
func (c *KubeJobController) ListJobByCronJob() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	cronJob := c.Input().Get("name")
	param := c.BuildKubernetesQueryParam()
	manager := c.Manager(cluster)
	result, err := job.GetRelatedJobByCronJob(manager.KubeClient, namespace, cronJob, param)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}
