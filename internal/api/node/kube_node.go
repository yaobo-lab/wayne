package node

import (
	"encoding/json"
	"fmt"
	"sync"

	v1 "k8s.io/api/core/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"wayne/internal/api/base"
	"wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/node"
	"wayne/internal/model"

	"wayne/pkg/logger"
)

type KubeNodeController struct {
	base.APIController
}

type Label struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type LabelSet struct {
	Labels []Label
}

func (c *KubeNodeController) Prepare() {

	c.APIController.Prepare()
	methodActionMap := map[string]string{
		"NodeStatistics": model.PermissionRead,
		"List":           model.PermissionRead,
		"Get":            model.PermissionRead,
		"Update":         model.PermissionUpdate,
		"Delete":         model.PermissionDelete,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, model.PermissionTypeKubeNode)
}

// swagger:route GET /api/v1/kubernetes/nodes/statistics node reqNodeStatisticsKubeNode
// kubernetes statistics
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) NodeStatistics() {
	cluster := c.Input().Get("cluster")
	total := 0
	countSyncMap := sync.Map{}
	countMap := make(map[string]int)
	if cluster == "" {
		managers := client.Managers()
		var errs []error
		wg := sync.WaitGroup{}

		managers.Range(func(key, value interface{}) bool {
			manager := value.(*client.ClusterManager)
			clu := key.(string)
			wg.Add(1)
			go func(clu string, mang *client.ClusterManager) {
				defer wg.Done()
				count, err := node.GetNodeCounts(mang.CacheFactory)
				if err != nil {
					errs = append(errs, err)
				}
				total += count
				countSyncMap.Store(clu, count)
			}(clu, manager)
			return true
		})

		wg.Wait()
		if len(errs) > 0 {
			c.HandleError(utilerrors.NewAggregate(errs))
			return
		}
		countSyncMap.Range(func(key, value interface{}) bool {
			countMap[key.(string)] = value.(int)
			return true
		})
	} else {

		manager, err := client.Manager(cluster)
		if err != nil {
			c.HandleError(err)
			return
		}

		count, err := node.GetNodeCounts(manager.CacheFactory)
		if err != nil {
			c.HandleError(err)
			return
		}
		total += count

	}

	c.Success(node.NodeStatistics{Total: total, Details: countMap})
}

// swagger:route GET /api/v1/kubernetes/nodes/clusters/{cluster} node reqListKubeNode
// list nodes
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) List() {
	cluster := c.Ctx.Input.Param(":cluster")
	manager := c.Manager(cluster)
	result, err := node.ListNode(manager.CacheFactory)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route GET /api/v1/kubernetes/nodes/{name}/clusters/{cluster} node reqGetKubeNode
// find node by cluster
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) Get() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)

	result, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route PUT /api/v1/kubernetes/nodes/{name}/clusters/{cluster} node reqUpdateKubeNode
// update the Node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) Update() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	var tpl v1.Node
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	if name != tpl.Name {
		c.HandleError(fmt.Errorf("name != tpl.Name"))
		return
	}

	cli := c.Client(cluster)
	result, err := node.UpdateNode(cli, &tpl)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(result)
}

// swagger:route DELETE /api/v1/kubernetes/nodes/{name}/clusters/{cluster} node reqDeleteKubeNode
// delete the Node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) Delete() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)
	err := node.DeleteNode(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success("ok!")
}

// swagger:route GET /api/v1/kubernetes/nodes/{name}/clusters/{cluster}/labels node reqGetLabelsKubeNode
// get labels of a node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) GetLabels() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)

	result, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	labels := result.ObjectMeta.Labels

	c.Success(labels)
}

// swagger:route POST /api/v1/kubernetes/nodes/{name}/clusters/{cluster}/label node reqAddLabelKubeNode
// add a label for a node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) AddLabel() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)
	var label Label
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &label)
	if err != nil {
		c.HandleError(err)
		return
	}

	nodeInfo, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	if len(nodeInfo.ObjectMeta.Labels) == 0 {
		nodeInfo.ObjectMeta.Labels = map[string]string{label.Key: label.Value}
	} else {
		nodeInfo.ObjectMeta.Labels[label.Key] = label.Value
	}

	newNode, err := node.UpdateNode(cli, nodeInfo)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(newNode)
}

// swagger:route DELETE /api/v1/kubernetes/nodes/{name}/clusters/{cluster}/label node reqDeleteLabelKubeNode
// delete a label of the node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) DeleteLabel() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)
	var label Label
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &label)
	if err != nil {
		c.HandleError(err)
		return
	}

	nodeInfo, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	if _, ok := nodeInfo.ObjectMeta.Labels[label.Key]; ok {
		delete(nodeInfo.ObjectMeta.Labels, label.Key)
	} else {
		return
	}

	newNode, err := node.UpdateNode(cli, nodeInfo)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(newNode)
}

// swagger:route POST /api/v1/kubernetes/nodes/{name}/clusters/{cluster}/labels node reqAddLabelsKubeNode
// Add labels in bulk for node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) AddLabels() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)
	var labels LabelSet
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &labels)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	if len(result.ObjectMeta.Labels) == 0 {
		result.ObjectMeta.Labels = make(map[string]string)
	}

	for _, label := range labels.Labels {
		result.ObjectMeta.Labels[label.Key] = label.Value
	}

	newNode, err := node.UpdateNode(cli, result)
	if err != nil {

		c.HandleError(err)
		return
	}
	c.Success(newNode)
}

// swagger:route DELETE /api/v1/kubernetes/nodes/{name}/clusters/{cluster}/labels node reqDeleteLabelsKubeNode
// Delete node labels in batches
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) DeleteLabels() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)
	var labels LabelSet
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &labels)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	for _, label := range labels.Labels {
		if _, ok := result.ObjectMeta.Labels[label.Key]; ok {
			delete(result.ObjectMeta.Labels, label.Key)
		} else {
			logger.Errorf("delete failed use the label key:(%s)", label.Key)
			return
		}
	}

	newNode, err := node.UpdateNode(cli, result)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(newNode)
}

// swagger:route POST /api/v1/kubernetes/nodes/{name}/clusters/{cluster}/taint node reqSetTaintKubeNode
// set taint for a node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) SetTaint() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)

	var taint v1.Taint
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &taint)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	taints := result.Spec.Taints
	if len(taints) == 0 {
		taints = []v1.Taint{}
		taints = append(taints, taint)
	} else {
		taints = append(taints, taint)
	}
	result.Spec.Taints = taints

	newNode, err := node.UpdateNode(cli, result)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(newNode)
}

// swagger:route DELETE /api/v1/kubernetes/nodes/{name}/clusters/{cluster}/taint node reqDeleteTaintKubeNode
// delete a taint from node
// responses:
//
//	200: respSuccessDescription
func (c *KubeNodeController) DeleteTaint() {
	cluster := c.Ctx.Input.Param(":cluster")
	name := c.Ctx.Input.Param(":name")
	cli := c.Client(cluster)

	var taint v1.Taint
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &taint)
	if err != nil {
		c.HandleError(err)
		return
	}

	result, err := node.GetNodeByName(cli, name)
	if err != nil {
		c.HandleError(err)
		return
	}
	taints := result.Spec.Taints
	if len(taints) == 0 {
		logger.Errorf("delete failed,the taint does not exist ")
		return
	}
	newTaints := []v1.Taint{}
	for _, t := range taints {
		if t.Key != taint.Key || t.Value != taint.Value || t.Effect != taint.Effect {
			newTaints = append(newTaints, t)
		}
	}
	result.Spec.Taints = newTaints

	newNode, err := node.UpdateNode(cli, result)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(newNode)
}
