package ingress

// swagger:parameters reqCreateKubeIngress
type reqCreateKubeIngress struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the ingressId
	// in: path
	IngressId string `json:"ingressId"`

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
