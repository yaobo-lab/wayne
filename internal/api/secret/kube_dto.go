package secret

// swagger:parameters reqCreateKubeSecret
type reqCreateKubeSecret struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the secretId
	// in: path
	SecretId string `json:"secretId"`

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
