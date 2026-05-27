package deployment

// swagger:parameters reqGetKubeDeployment reqDeleteKubeDeployment
type reqGetKubeDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the deployment
	// in: path
	Deployment string `json:"deployment"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqUpdateScaleKubeDeployment
type reqUpdateScaleKubeDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the deployment
	// in: path
	Deployment string `json:"deployment"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// number of replica
	// in: body
	// required: true
	Replica int `json:"replica"`
}

// swagger:parameters reqCreateKubeDeployment
type reqCreateKubeDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the deploymentId
	// in: path
	DeploymentId string `json:"deploymentId"`

	// the tplId
	// in: path
	TplId string `json:"tplId"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the tpl content
	// in: body
	// required: true
	Body int `json:"body"`
}

// swagger:parameters reqListKubeDeployment
type reqListKubeDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// column filter, ex. filter=name=test
	// in: query
	Filter string `json:"filter"`

	// column sorted by, ex. sortby=-id, '-' representation desc, and sortby=id representation asc
	// in: query
	SortBy string `json:"sortby"`
}
