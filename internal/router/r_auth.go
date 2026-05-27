package router

import (
	"wayne/internal/api/auth"

	"github.com/beego/beego/v2/server/web"
)

// 注册auth路由
func regAuthRouter() {
	web.CtrlGet("/currentuser", (*auth.AuthController).CurrentUser)
	web.CtrlPost("/login/:type/?:name", (*auth.AuthController).Login)
	web.CtrlGet("/login/:type/?:name", (*auth.AuthController).Login)
	web.CtrlGet("/logout", (*auth.AuthController).Logout)
}
