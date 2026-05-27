package namespace

import (
	"encoding/json"
	"strconv"

	"wayne/internal/api/base"
	"wayne/internal/model"
	common "wayne/pkg/dto"
)

type NamespaceController struct {
	base.APIController
}

func (c *NamespaceController) Prepare() {

	c.APIController.Prepare()

	perAction := ""
	_, method := c.GetControllerAndAction()
	switch method {
	case "Get", "List", "GetNames", "GetHistory":
		perAction = model.PermissionRead
	case "Create":
		perAction = model.PermissionCreate
	case "Update", "Migrate":
		perAction = model.PermissionUpdate
	case "Delete":
		perAction = model.PermissionDelete
	}
	if perAction != "" && !c.User.Admin {
		c.AbortForbidden("operation need admin permission.")
	}
}

// swagger:route GET /api/v1/namespaces/names namespace reqGetNamesNamespace
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) GetNames() {
	deleted := c.GetDeleteFromQuery()

	namespaces, err := model.NamespaceModel.GetNames(deleted)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(namespaces)
}

// swagger:route GET /api/v1/namespaces namespace reqListNamespace
// get all namespaces
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}

	total, err := model.GetTotal(new(model.Namespace), param)
	if err != nil {
		c.HandleError(err)
		return
	}
	namespaces := []model.Namespace{}

	err = model.GetAll(new(model.Namespace), &namespaces, param)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(param.NewPage(total, namespaces))
	return
}

// swagger:route POST /api/v1/namespaces namespace reqCreateNamespace
// create namespace
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) Create() {
	var ns model.Namespace
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ns)
	if err != nil {
		c.HandleError(err)
		return
	}
	ns.User = c.User.Name
	_, err = model.NamespaceModel.Add(&ns)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(ns)
}

// swagger:route GET /api/v1/namespaces/init namespace reqInitDefaultNamespace
// init default namespace
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) InitDefault() {
	if c.User.Admin == true {
		err := model.NamespaceModel.InitNamespace()
		if err != nil {
			c.HandleError(err)
			return
		}
	} else {
		c.AbortForbidden("Only admin can init.")
	}
	namespaces := []model.Namespace{}
	c.Success(namespaces)
	return
}

// swagger:route GET /api/v1/namespaces/{id} namespace reqGetNamespace
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) Get() {
	id := c.GetIDFromURL()

	ns, err := model.NamespaceModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(ns)
	return
}

// swagger:route PUT /api/v1/namespaces/{id} namespace reqUpdateNamespace
// update the namespace
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) Update() {
	id := c.GetIDFromURL()
	var ns model.Namespace
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ns)
	if err != nil {
		c.HandleError(err)
		return
	}
	ns.Id = int64(id)
	err = model.NamespaceModel.UpdateById(&ns)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(ns)
}

// swagger:route POST /api/v1/namespaces/{id} namespace reqDeleteNamespace
// delete the Namespace
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) Delete() {
	id := c.GetIDFromURL()
	logical := c.GetLogicalFromQuery()

	err := model.NamespaceModel.DeleteById(int64(id), logical)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}

// swagger:route GET /api/v1/namespaces/{id}/statistics namespace reqStatisticsNamespace
// The app resource count statistics
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceController) Statistics() {
	namespaceId := c.GetIDFromURL()
	appId, _ := strconv.ParseInt(c.Input().Get("app_id"), 10, 64)
	param := &common.QueryParam{
		Query: map[string]interface{}{
			"deleted":            false,
			"App__Namespace__Id": namespaceId,
		},
	}
	if appId != 0 {
		param.Query["App__Id"] = appId
	}
	resources := []string{
		model.TableNameDeployment,
		model.TableNameStatefulset,
		model.TableNameDaemonSet,
		model.TableNameCronjob,
		model.TableNameService,
		model.TableNameConfigMap,
		model.TableNameSecret,
		model.TableNamePersistentVolumeClaim,
	}

	var result = make(map[string]int64, 0)

	for _, resource := range resources {
		count, err := model.GetTotal(resource, param)
		if err != nil {
			c.HandleError(err)
			return
		}
		kubeApiType := model.TableToKubeApiTypeMap[resource]
		result[string(kubeApiType)] = count
	}

	c.Success(result)
}
