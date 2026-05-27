package pvc

// swagger:parameters reqCreateKubePersistentVolumeClaim
type reqCreateKubePersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the pvcId
	// in: path
	PvcId string `json:"pvcId"`

	// the tplId
	// in: path
	TplId string `json:"tplId"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the tpl content
	// in: body
	// required: true
	Body string `json:"body"`
}

// swagger:parameters reqActiveImageRobinPersistentVolumeClaim reqInActiveImageRobinPersistentVolumeClaim reqListSnapshotRobinPersistentVolumeClaim reqDeleteAllSnapshotRobinPersistentVolumeClaim reqGetPvcStatusRobinPersistentVolumeClaim reqOfflineImageUserRobinPersistentVolumeClaim reqLoginInfoRobinPersistentVolumeClaim reqVerifyRobinPersistentVolumeClaim
type reqActiveImageRobinPersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the pvc
	// in: path
	Pvc string `json:"pvc"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqDeleteSnapshotRobinPersistentVolumeClaim reqRollbackSnapshotRobinPersistentVolumeClaim reqCreateSnapshotRobinPersistentVolumeClaim
type reqDeleteSnapshotRobinPersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the pvc
	// in: path
	Pvc string `json:"pvc"`

	// the version
	// in: path
	Version string `json:"version"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}
