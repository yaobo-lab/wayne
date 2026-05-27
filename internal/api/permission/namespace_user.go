package permission

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

// 操作NamespaceUser相关资源
type NamespaceUserController struct {
	base.APIController
}

func (c *NamespaceUserController) Prepare() {

	c.APIController.Prepare()

	perAction := ""
	_, method := c.GetControllerAndAction()
	switch method {
	case "Get", "List":
		perAction = model.PermissionRead
	case "Create":
		perAction = model.PermissionCreate
	case "Update":
		perAction = model.PermissionUpdate
	case "Delete":
		perAction = model.PermissionDelete
	}
	if perAction != "" {
		c.CheckPermission(model.PermissionTypeNamespaceUser, perAction)
	}
}

// swagger:route GET /api/v1/namespaces/{namespaceid}/users permission reqListNamespaceUser
// get all namespaceUser
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceUserController) List() {
	param := c.BuildQueryParam()

	if c.NamespaceId != 0 {
		param.Query["Namespace__Id__exact"] = c.NamespaceId
	}

	userId := c.Input().Get("userId")
	if userId != "" {
		param.Query["User__Id__exact"] = userId
	}
	userName := c.Input().Get("userName")
	if userName != "" {
		param.Query["User__Name__contains"] = userName
	}
	param.Groupby = []string{"Namespace", "User"}

	total, err := model.GetTotal(new(model.NamespaceUser), param)
	if err != nil {
		c.HandleError(err)
		return
	}
	namespaceUsers := []model.NamespaceUser{}

	err = model.GetAll(new(model.NamespaceUser), &namespaceUsers, param)
	if err != nil {
		c.HandleError(err)
		return
	}

	// 获取这批用户的group列表
	err = model.NamespaceUserModel.SetGroupsName(namespaceUsers)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(param.NewPage(total, namespaceUsers))
}

// swagger:route POST /api/v1/namespaces/{namespaceid}/users permission reqCreateNamespaceUser
// create namespaceUser
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceUserController) Create() {
	var namespaceUser model.NamespaceUser
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &namespaceUser)
	if err != nil {
		c.HandleError(err)
		return
	}

	oneGroup := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroup != "" {
		groupsFlag = false
	}

	_, err = model.NamespaceUserModel.Add(&namespaceUser, groupsFlag)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(namespaceUser)
}

// swagger:route GET /api/v1/namespaces/{namespaceid}/users/{id} permission reqGetNamespaceUser
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceUserController) Get() {
	id := c.GetIDFromURL()

	oneGroups := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroups != "" {
		groupsFlag = false
	}

	ns, err := model.NamespaceUserModel.GetById(int64(id), groupsFlag)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(ns)
}

// swagger:route PUT /api/v1/namespaces/{namespaceid}/users/{id} permission reqUpdateNamespaceUser
// update the NamespaceUser
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceUserController) Update() {
	id := c.GetIDFromURL()
	var namespaceUser model.NamespaceUser
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &namespaceUser)
	if err != nil {
		c.HandleError(err)
		return
	}
	namespaceUser.Id = int64(id)

	oneGroup := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroup != "" {
		groupsFlag = false
	}
	err = model.NamespaceUserModel.UpdateById(&namespaceUser, groupsFlag)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(namespaceUser)
}

// swagger:route DELETE /api/v1/namespaces/{namespaceid}/users/{id} permission reqDeleteNamespaceUser
// delete the NamespaceUser
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceUserController) Delete() {
	id := c.GetIDFromURL()

	oneGroup := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroup != "" {
		groupsFlag = false
	}

	err := model.NamespaceUserModel.DeleteById(int64(id), groupsFlag)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}

// swagger:route GET /api/v1/namespaces/{namespaceid}/users/permissions/{id} permission reqGetPermissionByNSNamespaceUser
// get PerNS by nsId
// responses:
//
//	200: respSuccessDescription
func (c *NamespaceUserController) GetPermissionByNS() {
	id := c.GetIDFromURL()

	nsPers, err := model.NamespaceUserModel.GetAllPermission(int64(id), c.User.Id)
	if err != nil {
		c.HandleError(err)
		return
	}

	var ret model.TypePermission
	mapPer := make(map[string]map[string]bool)
	for _, permission := range nsPers {
		paction, ptype, err := model.PermissionModel.SplitName(permission.Name)
		if err != nil {
			c.HandleError(err)
			return
		}
		_, ok := mapPer[ptype]
		if ok != true {
			mapPer[ptype] = make(map[string]bool)
		}
		mapPer[ptype][paction] = true
	}

	if err = mapstructure.Decode(mapPer, &ret); err != nil {
		c.HandleError(err)
		return
	}

	c.Success(ret)
}
