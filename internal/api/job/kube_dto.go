package job

// swagger:parameters reqListJobByCronJobKubeJob
type reqListJobByCronJobKubeJob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the namespace
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

	// the cronjob name.
	// in: query
	// required: true
	Name string `json:"name"`
}
