package configmap

// swagger:parameters reqCreateKubeConfigMap
type reqCreateKubeConfigMap struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the configMapId
	// in: path
	ConfigMapId string `json:"configMapId"`

	// the tplId
	// in: path
	TplId string `json:"tplId"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the tpl content
	// in: body
	// required: true
	Body string `json:"body"`
}
