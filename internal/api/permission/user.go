package permission

import (
	"encoding/json"

	"wayne/internal/api/base"
	"wayne/internal/model"
)

type UserController struct {
	base.APIController
}

func (c *UserController) Prepare() {

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
	case "ResetPassword":
		perAction = model.PermissionUpdate
	case "UpdateAdmin":
		perAction = model.PermissionUpdate
	}
	if perAction != "" && !c.User.Admin {
		c.AbortForbidden("operation need admin permission.")
	}
}

// swagger:route GET /api/v1/users/statistics permission reqUserStatisticsUser
// user count statistics
// responses:
//
//	200: respSuccessDescription
func (c *UserController) UserStatistics() {
	param := c.BuildQueryParam()

	total, err := model.GetTotal(new(model.User), param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(model.UserStatistics{Total: total})
}

// swagger:route GET /api/v1/users permission reqListUser
// get all user
// responses:
//
//	200: respSuccessDescription
func (c *UserController) List() {
	param := c.BuildQueryParam()
	name := c.Input().Get("name")
	if name != "" {
		param.Query["name__contains"] = name
	}
	total, err := model.GetTotal(new(model.User), param)
	if err != nil {
		c.HandleError(err)
		return
	}

	users := []model.User{}
	err = model.GetAll(new(model.User), &users, param)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(param.NewPage(total, users))
}

// swagger:route GET /api/v1/users/names permission reqGetNamesUser
// get all id and names
// responses:
//
//	200: respSuccessDescription
func (c *UserController) GetNames() {
	users, err := model.UserModel.GetNames()
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(users)
}

// swagger:route POST /api/v1/users permission reqCreateUser
// create user
// responses:
//
//	200: respSuccessDescription
func (c *UserController) Create() {
	var user model.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.HandleError(err)
		return
	}
	_, err = model.UserModel.AddUser(&user)

	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(user)
}

// swagger:route GET /api/v1/users/{id} permission reqGetUser
// find Object by id
// responses:
//
//	200: respSuccessDescription
func (c *UserController) Get() {
	id := c.GetIDFromURL()

	user, err := model.UserModel.GetUserById(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(user)
}

// swagger:route PUT /api/v1/users/{id}/resetpassword permission reqResetPasswordUser
// update the user admin
// responses:
//
//	200: respSuccessDescription
func (c *UserController) ResetPassword() {
	var user *struct {
		Id       int64  `json:"id"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.UserModel.ResetUserPassword(user.Id, user.Password)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(user)
}

// swagger:route PUT /api/v1/users/{id}/admin permission reqUpdateAdminUser
// update the user admin
// responses:
//
//	200: respSuccessDescription
func (c *UserController) UpdateAdmin() {
	var user *model.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.UserModel.UpdateUserAdmin(user)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(user)
}

// swagger:route PUT /api/v1/users/{id} permission reqUpdateUser
// update the user
// responses:
//
//	200: respSuccessDescription
func (c *UserController) Update() {
	var user *model.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.HandleError(err)
		return
	}

	err = model.UserModel.UpdateUserById(user)
	if err != nil {
		c.HandleError(err)
		return
	}

	c.Success(user)
}

// swagger:route DELETE /api/v1/users/{id} permission reqDeleteUser
// delete the User
// responses:
//
//	200: respSuccessDescription
func (c *UserController) Delete() {
	id := c.GetIDFromURL()

	err := model.UserModel.DeleteUser(int64(id))
	if err != nil {
		c.HandleError(err)
		return
	}
	c.Success(nil)
}
