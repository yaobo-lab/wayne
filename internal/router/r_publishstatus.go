package router

import (
	"wayne/internal/api/publishstatus"

	"github.com/beego/beego/v2/server/web"
)

// 注册publishstatus路由
func regPublishStatusRouter() {
	r := []web.LinkNamespace{
		web.NSCtrlGet("/", (*publishstatus.PublishStatusController).List),
		web.NSCtrlDelete("/:id", (*publishstatus.PublishStatusController).Delete),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/publishstatus", r...),
	)

	//注册路由组
	web.AddNamespace(ns)
}
