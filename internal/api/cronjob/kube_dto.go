package cronjob

// swagger:parameters reqGetKubeCronjob reqDeleteKubeCronjob
type reqGetKubeCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cronjob
	// in: path
	Cronjob string `json:"cronjob"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqCreateKubeCronjob
type reqCreateKubeCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cronjobId
	// in: path
	CronjobId string `json:"cronjobId"`

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

// swagger:parameters reqSuspendKubeCronjob
type reqSuspendKubeCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cronjobId
	// in: path
	CronjobId string `json:"cronjobId"`

	// the tplId
	// in: path
	TplId string `json:"tplId"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`
}
