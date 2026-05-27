package event

// swagger:parameters reqListKubeEvent
type reqListKubeEvent struct {
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

	// the query type. deployments, statefulsets, daemonsets,cronjobs
	// in: query
	// required: true
	Type string `json:"type"`

	// the query resource name.
	// in: query
	// required: true
	Name string `json:"name"`
}
