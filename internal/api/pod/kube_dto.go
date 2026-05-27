package pod

// swagger:parameters reqTerminalKubePod
type reqTerminalKubePod struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the pod
	// in: path
	Pod string `json:"pod"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the cmd you want to exec.
	// in: query
	// required: true
	Cmd string `json:"cmd"`

	// the container name.
	// in: query
	// required: true
	Container string `json:"container"`
}

// swagger:parameters reqListKubePod
type reqListKubePod struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// the query type. deployment, statefulset, daemonSet, job, pod
	// in: query
	// required: true
	Type string `json:"type"`

	// the query resource name.
	// in: query
	Name string `json:"name"`
}

// swagger:parameters reqPodStatisticsKubePod
type reqPodStatisticsKubePod struct {
	// the cluster
	// in: query
	Cluster string `json:"cluster"`
}
