package daemonset

// swagger:parameters reqGetKubeDaemonSet reqDeleteKubeDaemonSet
type reqGetKubeDaemonSet struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the daemonSet
	// in: path
	DaemonSet string `json:"daemonSet"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqCreateKubeDaemonSet
type reqCreateKubeDaemonSet struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the daemonSetId
	// in: path
	DaemonSetId string `json:"daemonSetId"`

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
