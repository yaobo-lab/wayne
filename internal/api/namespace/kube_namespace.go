package namespace

import (
	"encoding/json"
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"wayne/internal/api/base"
	"wayne/internal/k8s/client"
	common "wayne/internal/k8s/dto"
	"wayne/internal/k8s/kind/namespace"
	"wayne/internal/model"
	util "wayne/pkg"
	"wayne/pkg/hack"
	"wayne/pkg/maps"
)

type KubeNamespaceController struct {
	base.APIController
}

func (c *KubeNamespaceController) Prepare() {
	c.APIController.Prepare()
	methodActionMap := map[string]string{
		"Resources":  model.PermissionRead,
		"Statistics": model.PermissionRead,
		"Create":     model.PermissionCreate,
	}
	_, method := c.GetControllerAndAction()
	switch method {
	case "Resources", "Statistics":
		c.PreparePermission(methodActionMap, method, model.PermissionTypeNamespace)
	case "Create":
		c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeNamespace)
	}
}

// swagger:route POST /api/v1/kubernetes/namespaces/{name}/clusters/{cluster} namespace reqCreateKubeNamespace
// create the namespace
// responses:
//
//	200: respSuccessDescription
func (c *KubeNamespaceController) Create() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	tpl := new(v1.Namespace)
	tpl.Name = name

	cli := c.Client(cluster)
	// If the namespace does not exist, the value of result is nil.
	result, err := namespace.CreateNotExitNamespace(cli, tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route GET /api/v1/kubernetes/namespaces/{namespaceid}/resources namespace reqResourcesKubeNamespace
// Get namespace resource statistics
// responses:
//
//	200: respSuccessDescription
func (c *KubeNamespaceController) Resources() {
	appName := c.Input().Get("app")
	id := c.GetIntParamFromURL(":namespaceid")
	ns, err := model.NamespaceModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	var namespaceMetaData model.NamespaceMetaData
	err = json.Unmarshal(hack.Slice(ns.MetaData), &namespaceMetaData)
	if err != nil {
		c.HandleError(err)
		return
	}

	syncResourceMap := sync.Map{}
	var errs []error
	wg := sync.WaitGroup{}

	managers := client.Managers()
	managers.Range(func(key, value interface{}) bool {
		manager := value.(*client.ClusterManager)
		wg.Add(1)
		go func(m *client.ClusterManager) {
			defer wg.Done()
			clusterMetas, ok := namespaceMetaData.ClusterMetas[m.Cluster.Name]
			// can't use current cluster
			if !ok {
				return
			}

			selectorMap := map[string]string{
				util.NamespaceLabelKey: ns.Name,
			}
			if appName != "" {
				selectorMap[util.AppLabelKey] = appName
			}
			selector := labels.SelectorFromSet(selectorMap)
			resourceUsage, err := namespace.ResourcesUsageByNamespace(m.KubeClient, ns.KubeNamespace, selector.String())
			if err != nil {
				errs = append(errs, err)
				return
			}
			syncResourceMap.Store(m.Cluster.Name, common.Resource{
				Usage: &common.ResourceList{
					Cpu:    resourceUsage.Cpu / 1000,
					Memory: resourceUsage.Memory / (1024 * 1024 * 1024),
				},
				Limit: &common.ResourceList{
					Cpu:    clusterMetas.ResourcesLimit.Cpu,
					Memory: clusterMetas.ResourcesLimit.Memory,
				},
			})
		}(manager)
		return true
	})
	wg.Wait()

	if len(errs) == maps.SyncMapLen(managers) && len(errs) > 0 {
		c.HandleError(utilerrors.NewAggregate(errs))
		return
	}
	result := make(map[string]common.Resource)
	syncResourceMap.Range(func(key, value interface{}) bool {
		result[key.(string)] = value.(common.Resource)
		return true
	})
	c.Success(result)
}

// swagger:route GET /api/v1/kubernetes/namespaces/{namespaceid}/statistics namespace reqStatisticsKubeNamespace
// Get namespace resource statistics for report
// responses:
//
//	200: respSuccessDescription
func (c *KubeNamespaceController) Statistics() {
	appName := c.Input().Get("app")
	id := c.GetIntParamFromURL(":namespaceid")
	ns, err := model.NamespaceModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	var namespaceMetaData model.NamespaceMetaData
	err = json.Unmarshal(hack.Slice(ns.MetaData), &namespaceMetaData)
	if err != nil {
		c.HandleError(err)
		return
	}

	syncResourceMap := sync.Map{}
	var errs []error
	wg := sync.WaitGroup{}
	managers := client.Managers()
	managers.Range(func(key, value interface{}) bool {
		manager := value.(*client.ClusterManager)
		wg.Add(1)
		go func(m *client.ClusterManager) {
			defer wg.Done()
			selectorMap := map[string]string{
				util.NamespaceLabelKey: ns.Name,
			}
			if appName != "" {
				selectorMap[util.AppLabelKey] = appName
			}
			selector := labels.SelectorFromSet(selectorMap)
			resourceUsage, err := namespace.ResourcesOfAppByNamespace(m.KubeClient, ns.KubeNamespace, selector.String())
			if err != nil {

				errs = append(errs, err)
				return
			}
			syncResourceMap.Store(m.Cluster.Name, resourceUsage)
		}(manager)
		return true
	})

	wg.Wait()
	if len(errs) == maps.SyncMapLen(managers) && len(errs) > 0 {
		c.HandleError(utilerrors.NewAggregate(errs))
		return
	}
	result := make(map[string]*common.ResourceApp)
	syncResourceMap.Range(func(key, value interface{}) bool {
		for k, v := range value.(map[string]*common.ResourceApp) {
			if result[k] == nil {
				result[k] = v
			} else {
				result[k].Cpu += v.Cpu
				result[k].Memory += v.Memory
				result[k].PodNum += v.PodNum
			}
		}
		return true
	})
	c.Success(result)

}
