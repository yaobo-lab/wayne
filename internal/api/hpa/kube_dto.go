package hpa

// swagger:parameters reqCreateKubeHPA
type reqCreateKubeHPA struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the hpaId
	// in: path
	HpaId string `json:"hpaId"`

	// the tplId
	// in: path
	TplId string `json:"tplId"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}
