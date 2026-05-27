package router

import (
	"wayne/internal/api/pvc"

	"github.com/beego/beego/v2/server/web"
)

// 注册pvc路由
func regPVCRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*pvc.PersistentVolumeClaimController).List),
		web.NSCtrlPost("/", (*pvc.PersistentVolumeClaimController).Create),
		web.NSCtrlDelete("/:id", (*pvc.PersistentVolumeClaimController).Delete),
		web.NSCtrlPut("/:id", (*pvc.PersistentVolumeClaimController).Update),
		web.NSCtrlGet("/:id([0-9]+)", (*pvc.PersistentVolumeClaimController).Get),
		web.NSCtrlGet("/names", (*pvc.PersistentVolumeClaimController).GetNames),
		web.NSCtrlPut("/updateorders", (*pvc.PersistentVolumeClaimController).UpdateOrders),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlGet("/", (*pvc.PersistentVolumeClaimTplController).List),
		web.NSCtrlPost("/", (*pvc.PersistentVolumeClaimTplController).Create),
		web.NSCtrlPut("/:id", (*pvc.PersistentVolumeClaimTplController).Update),
		web.NSCtrlDelete("/:id", (*pvc.PersistentVolumeClaimTplController).Delete),
		web.NSCtrlGet("/:id([0-9]+)", (*pvc.PersistentVolumeClaimTplController).Get),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/apps/:appid([0-9]+)/persistentvolumeclaims", r1...),
		web.NSNamespace("/apps/:appid([0-9]+)/persistentvolumeclaims/tpls", r2...),
	)

	//注册路由组
	web.AddNamespace(ns)
}

// 注册kubernetes pvc路由
func regKubePVCRouter() {
	r1 := []web.LinkNamespace{
		web.NSCtrlPost("/:pvcId/tpls/:tplId/clusters/:cluster", (*pvc.KubePersistentVolumeClaimController).Create),
	}

	r2 := []web.LinkNamespace{
		web.NSCtrlPost("/:pvc/rbd/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).ActiveImage),
		web.NSCtrlDelete("/:pvc/rbd/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).InActiveImage),
		web.NSCtrlDelete("/:pvc/snapshot/:version/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).DeleteSnapshot),
		web.NSCtrlPut("/:pvc/snapshot/:version/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).RollbackSnapshot),
		web.NSCtrlPost("/:pvc/snapshot/:version/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).CreateSnapshot),
		web.NSCtrlGet("/:pvc/snapshot/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).ListSnapshot),
		web.NSCtrlDelete("/:pvc/snapshot/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).DeleteAllSnapshot),
		web.NSCtrlGet("/:pvc/status/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).GetPvcStatus),
		web.NSCtrlDelete("/:pvc/user/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).OfflineImageUser),
		web.NSCtrlGet("/:pvc/user/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).LoginInfo),
		web.NSCtrlGet("/:pvc/verify/namespaces/:namespace/clusters/:cluster", (*pvc.RobinPersistentVolumeClaimController).Verify),
	}

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/persistentvolumeclaims", r1...),
		web.NSNamespace("/kubernetes/apps/:appid([0-9]+)/persistentvolumeclaims/robin", r2...),
	)
	//注册路由组
	web.AddNamespace(ns)
}
