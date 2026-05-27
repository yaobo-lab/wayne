package statefulset

// swagger:parameters reqGetKubeStatefulset
type reqGetKubeService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the statefulset
	// in: path
	StatefulSet string `json:"statefulset"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqCreateKubeStatefulset
type reqCreateKubeStatefulset struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the statefulsetId
	// in: path
	StatefulSetId string `json:"statefulsetId"`

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
