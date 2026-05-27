package pv

// swagger:parameters reqGetKubePersistentVolume reqDeleteKubePersistentVolume
type reqGetKubePersistentVolume struct {
	// the name
	// in: path
	Name string `json:"name"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqUpdateKubePersistentVolume
type reqUpdateKubePersistentVolume struct {
	// the name
	// in: path
	Name string `json:"name"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}

// swagger:parameters reqListKubePersistentVolume reqListRbdImagesRobinPersistentVolume
type reqListKubePersistentVolume struct {
	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqCreateKubePersistentVolume reqCreateRbdImageRobinPersistentVolume
type reqCreateKubePersistentVolume struct {
	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}
