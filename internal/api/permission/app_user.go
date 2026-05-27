package permission

import (
	"encoding/json"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/mitchellh/mapstructure"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

// 操作AppUser相关资源
type AppUserController struct {
	base.APIController
}

func (c *AppUserController) Prepare() {

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
		c.CheckPermission(model.PermissionTypeAppUser, perAction)
	}
}

// swagger:route GET /api/v1/apps/{appid}/users permission reqListAppUser
// get all appUser
// responses:
//
//	200: respSuccessDescription
func (c *AppUserController) List() {
	param := c.BuildQueryParam()

	if c.AppId != 0 {
		param.Query["App__Id__exact"] = c.AppId
	}

	userId := c.Input().Get("userId")
	if userId != "" {
		param.Query["User__Id__exact"] = userId
	}
	userName := c.Input().Get("userName")
	if userName != "" {
		param.Query["User__Name__contains"] = userName
	}
	param.Groupby = []string{"App", "User"}

	total, err := model.GetTotal(new(model.AppUser), param)
	if err != nil {
		c.HandleError(err)
		return
	}
	appUsers := []model.AppUser{}

	err = model.GetAll(new(model.AppUser), &appUsers, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	// 获取这批用户的group列表
	err = model.AppUserModel.SetGroupsName(appUsers)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(param.NewPage(total, appUsers))
}

// swagger:route POST /api/v1/apps/{appid}/users permission reqCreateAppUser
// create appUser
// responses:
//
//	200: respSuccessDescription
func (c *AppUserController) Create() {
	var appUser model.AppUser
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &appUser)
	if err != nil {
		c.HandleError(err)
		return
	}

	oneGroup := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroup != "" {
		groupsFlag = false
	}

	// 检查该app所属的namespace，是否配置了group
	app, err := model.AppModel.GetById(appUser.App.Id)
	if err != nil {
		c.HandleError(err)
		return
	}
	_, err = model.NamespaceUserModel.GetByNamespaceIdAndUserId(app.Namespace.Id, appUser.User.Id)
	if err == orm.ErrNoRows {
		c.AbortForbidden("User not in namespace.")
	} else if err != nil {
		c.HandleError(err)
		return
	}

	_, err = model.AppUserModel.Add(&appUser, groupsFlag)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(appUser)
}

// swagger:route GET /api/v1/apps/{appid}/users/{id} permission reqGetAppUser
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *AppUserController) Get() {
	id := c.GetIDFromURL()

	oneGroup := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroup != "" {
		groupsFlag = false
	}
	ns, err := model.AppUserModel.GetById(int64(id), groupsFlag)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(ns)
	return
}

// swagger:route GET /api/v1/apps/{appid}/users/permissions/{id} permission reqGetPermissionByAppAppUser
// get PerApp by appId
// responses:
//
//	200: respSuccessDescription
func (c *AppUserController) GetPermissionByApp() {
	id := c.GetIDFromURL()

	appPers, err := model.AppUserModel.GetAllPermission(int64(id), c.User.Id)
	if err != nil {
		c.HandleError(err)
		return
	}

	app, err := model.AppModel.GetById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	nsPers, err := model.NamespaceUserModel.GetAllPermission(app.Namespace.Id, c.User.Id)
	if err != nil {
		c.HandleError(err)
		return
	}

	var ret model.TypePermission
	mapPer := make(map[string]map[string]bool)
	for _, permission := range appPers {
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
	return
}

// swagger:route PUT /api/v1/apps/{appid}/users/{id} permission reqUpdateAppUser
// update the AppUser
// responses:
//
//	200: respSuccessDescription
func (c *AppUserController) Update() {
	id := c.GetIDFromURL()
	oneGroup := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroup != "" {
		groupsFlag = false
	}
	var appUser model.AppUser
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &appUser)
	if err != nil {
		c.HandleError(err)
		return
	}

	appUser.Id = int64(id)
	err = model.AppUserModel.UpdateById(&appUser, groupsFlag)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(appUser)
}

// swagger:route DELETE /api/v1/apps/{appid}/users/{id} permission reqDeleteAppUser
// delete the AppUser
// responses:
//
//	200: respSuccessDescription
func (c *AppUserController) Delete() {
	id := c.GetIDFromURL()
	oneGroup := c.Input().Get("oneGroup")
	groupsFlag := true
	if oneGroup != "" {
		groupsFlag = false
	}

	err := model.AppUserModel.DeleteById(int64(id), groupsFlag)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
