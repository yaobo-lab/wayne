package pod

import (
	"sync"

	"wayne/internal/api/base"
	"wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/pod"
	"wayne/internal/model"

	"wayne/pkg/logger"
)

type KubePodController struct {
	base.APIController
}

func (c *KubePodController) Prepare() {

	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"PodStatistics": model.PermissionRead,
		"List":          model.PermissionRead,
		"Terminal":      model.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubePod)
}

// swagger:route GET /api/v1/kubernetes/pods/statistics pod reqPodStatisticsKubePod
// kubernetes statistics
// responses:
//
//	200: respSuccessDescription
func (c *KubePodController) PodStatistics() {
	cluster := c.Input().Get("cluster")
	total := 0
	countSyncMap := sync.Map{}
	countMap := make(map[string]int)
	if cluster == "" {
		managers := client.Managers()
		wg := sync.WaitGroup{}

		managers.Range(func(key, value interface{}) bool {
			manager := value.(*client.ClusterManager)
			wg.Add(1)
			go func(manager *client.ClusterManager) {
				defer wg.Done()
				count, err := pod.GetPodCounts(manager.CacheFactory)
				if err != nil {
					logger.Errorf("get pod counts error.", key, err)
					return
				}
				total += count
				countSyncMap.Store(manager.Cluster.Name, count)
			}(manager)
			return true
		})

		wg.Wait()
		countSyncMap.Range(func(key, value interface{}) bool {
			countMap[key.(string)] = value.(int)
			return true
		})

	} else {
		manager, err := client.Manager(cluster)
		if err == nil {
			count, err := pod.GetPodCounts(manager.CacheFactory)
			if err != nil {
				c.HandleError(err)
				return
			}
			total += count
		} else {
			c.HandleError(err)
			return
		}

	}

	c.Success(pod.PodStatistics{Total: total, Details: countMap})
}

// swagger:route GET /api/v1/kubernetes/apps/{appid}/pods/namespaces/{namespace}/clusters/{cluster} pod reqListKubePod
// find pods by resource type
// responses:
//
//	200: respSuccessDescription
func (c *KubePodController) List() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	resourceType := c.Input().Get("type")
	resourceName := c.Input().Get("name")
	param := c.BuildKubernetesQueryParam()
	manager := c.Manager(cluster)
	result, err := pod.GetPodListPageByType(manager.KubeClient, namespace, resourceName, resourceType, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}
