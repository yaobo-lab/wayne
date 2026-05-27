package service

// swagger:parameters reqGetKubeService
type reqGetKubeService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the service
	// in: path
	Service string `json:"service"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqCreateKubeService
type reqCreateKubeService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the serviceId
	// in: path
	ServiceId string `json:"serviceId"`

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
