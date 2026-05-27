package log

// swagger:parameters reqListKubeLog
type reqListKubeLog struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the pod name
	// in: path
	Pod string `json:"pod"`

	// the container name
	// in: path
	Container string `json:"container"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// log tail lines.
	// in: query
	// required: true
	TailLines int `json:"tailLines"`
}
