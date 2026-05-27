package router

import (
	"wayne/internal/api/permission"

	"github.com/beego/beego/v2/server/web"
)

// 注册permission路由
func regPermissionRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*permission.AppUserController).List),
		web.NSCtrlPost("/", (*permission.AppUserController).Create),
		web.NSCtrlGet("/:id", (*permission.AppUserController).Get),
		web.NSCtrlPut("/:id", (*permission.AppUserController).Update),
		web.NSCtrlDelete("/:id", (*permission.AppUserController).Delete),
		web.NSCtrlGet("/permissions/:id", (*permission.AppUserController).GetPermissionByApp),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*permission.GroupController).List),
		web.NSCtrlPost("/", (*permission.GroupController).Create),
		web.NSCtrlGet("/:id", (*permission.GroupController).Get),
		web.NSCtrlPut("/:id", (*permission.GroupController).Update),
		web.NSCtrlDelete("/:id", (*permission.GroupController).Delete),
	}

	r3 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*permission.NamespaceUserController).List),
		web.NSCtrlPost("/", (*permission.NamespaceUserController).Create),
		web.NSCtrlGet("/:id", (*permission.NamespaceUserController).Get),
		web.NSCtrlPut("/:id", (*permission.NamespaceUserController).Update),
		web.NSCtrlDelete("/:id", (*permission.NamespaceUserController).Delete),
		web.NSCtrlGet("/permissions/:id", (*permission.NamespaceUserController).GetPermissionByNS),
	}

	r4 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*permission.PermissionController).List),
		web.NSCtrlPost("/", (*permission.PermissionController).Create),
		web.NSCtrlGet("/:id", (*permission.PermissionController).Get),
		web.NSCtrlPut("/:id", (*permission.PermissionController).Update),
		web.NSCtrlDelete("/:id", (*permission.PermissionController).Delete),
	}

	r5 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*permission.UserController).List),
		web.NSCtrlPost("/", (*permission.UserController).Create),
		web.NSCtrlDelete("/:id", (*permission.UserController).Delete),
		web.NSCtrlGet("/:id", (*permission.UserController).Get),
		web.NSCtrlPut("/:id", (*permission.UserController).Update),
		web.NSCtrlPut("/:id/admin", (*permission.UserController).UpdateAdmin),
		web.NSCtrlPut("/:id/resetpassword", (*permission.UserController).ResetPassword),
		web.NSCtrlGet("/names", (*permission.UserController).GetNames),
		web.NSCtrlGet("/statistics", (*permission.UserController).UserStatistics),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/users", r1...),
		web.NSNamespace("/groups", r2...),
		web.NSNamespace("/namespaces/:namespaceid([0-9]+)/users", r3...),
		web.NSNamespace("/permissions", r4...),
		web.NSNamespace("/users", r5...),
	)

	//注册路由组
	web.AddNamespace(ns)
}
