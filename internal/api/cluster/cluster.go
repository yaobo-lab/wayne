package cluster

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

// 集群相关操作
type ClusterController struct {
	base.APIController
}

func (c *ClusterController) Prepare() {
	c.APIController.Prepare()
	perAction := ""
	_, method := c.GetControllerAndAction()
	switch method {
	case "Create":
		perAction = model.PermissionCreate
	case "Update":
		perAction = model.PermissionUpdate
	case "Delete":
		perAction = model.PermissionDelete
	}
	if perAction != "" && !c.User.Admin {
		c.AbortForbidden("operation need admin permission.")
	}
}

// swagger:route GET /api/v1/clusters/names cluster reqGetNamesCluster
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *ClusterController) GetNames() {
	deleted := c.GetDeleteFromQuery()

	services, err := model.ClusterModel.GetNames(deleted)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(services)
}

// swagger:route POST /api/v1/clusters cluster reqCreateCluster
// create Cluster
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ClusterController) Create() {
	var cluster model.Cluster
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster)
	if err != nil {
		c.HandleError(err)
		return
	}
	cluster.User = c.User.Name

	objectid, err := model.ClusterModel.Add(&cluster)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(objectid)
}

// swagger:route PUT /api/v1/clusters/{name} cluster reqUpdateCluster
// update the object
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ClusterController) Update() {
	name := c.Ctx.Input.Param(":name")

	var cluster model.Cluster
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster)
	if err != nil {
		c.HandleError(err)
		return
	}
	cluster.Name = name
	err = model.ClusterModel.UpdateByName(&cluster)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(cluster)
}

// swagger:route GET /api/v1/clusters/{name} cluster reqGetCluster
// find Object by objectid
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ClusterController) Get() {
	name := c.Ctx.Input.Param(":name")

	cluster, err := model.ClusterModel.GetByName(name)
	if err != nil {
		c.HandleError(err)
		return
	}
	// 非admin用户不允许查看kubeconfig配置
	if !c.User.Admin {
		cluster.KubeConfig = ""
	}
	c.Success(cluster)
}

// swagger:route GET /api/v1/clusters cluster reqListCluster
// get all objects
// responses:
//
//	200: respSuccessDescription
func (c *ClusterController) List() {
	param := c.BuildQueryParam()
	clusters := []model.Cluster{}

	total, err := model.GetTotal(new(model.Cluster), param)
	if err != nil {
		c.HandleError(err)
		return
	}
	err = model.GetAll(new(model.Cluster), &clusters, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	// 非admin用户不允许查看kubeconfig配置
	if !c.User.Admin {
		for i := range clusters {
			clusters[i].KubeConfig = ""
		}
	}

	c.Success(param.NewPage(total, clusters))
	return
}

// swagger:route DELETE /api/v1/clusters/{name} cluster reqDeleteCluster
// delete the cluster
// responses:
//
//	200: respSuccessDescription
//	403: respFailureDescription
func (c *ClusterController) Delete() {
	name := c.Ctx.Input.Param(":name")

	logical := c.GetLogicalFromQuery()

	err := model.ClusterModel.DeleteByName(name, logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
