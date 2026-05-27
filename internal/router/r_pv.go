package router

import (
	"wayne/internal/api/pv"

	"github.com/beego/beego/v2/server/web"
)

// 注册kubernetes pv路由
func regKubePVRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/:name/clusters/:cluster", (*pv.KubePersistentVolumeController).Get),
		web.NSCtrlPut("/:name/clusters/:cluster", (*pv.KubePersistentVolumeController).Update),
		web.NSCtrlDelete("/:name/clusters/:cluster", (*pv.KubePersistentVolumeController).Delete),
		web.NSCtrlGet("/clusters/:cluster", (*pv.KubePersistentVolumeController).List),
		web.NSCtrlPost("/clusters/:cluster", (*pv.KubePersistentVolumeController).Create),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/rbd.images/clusters/:cluster", (*pv.RobinPersistentVolumeController).ListRbdImages),
		web.NSCtrlPost("/rbd.images/clusters/:cluster", (*pv.RobinPersistentVolumeController).CreateRbdImage),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/persistentvolumes", r1...),
		web.NSNamespace("/kubernetes/persistentvolumes/robin", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
