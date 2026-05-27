package cluster

// swagger:parameters reqCreateCluster
type reqCreateCluster struct {
	// the cluster content
	// in: body
	Body string `json:"body"`
}

// swagger:parameters reqListCluster
type reqListCluster struct {
	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`
}

// swagger:parameters reqUpdateCluster
type reqUpdateCluster struct {
	// the name you want to update
	// in: path
	Name string `json:"name"`

	// the cluster content
	// in: body
	Body string `json:"body"`
}

// swagger:parameters reqGetCluster
type reqGetCluster struct {
	// the name you want to update
	// in: path
	Name string `json:"name"`
}

// swagger:parameters reqDeleteCluster
type reqDeleteCluster struct {
	// the name you want to delete
	// in: path
	Name string `json:"name"`

	// is logical deletion,default true
	// in: query
	Logical bool `json:"logical"`
}

// swagger:parameters reqGetNamesCluster
type reqGetNamesCluster struct {
	// is deleted,default false.
	// in: query
	Deleted bool `json:"deleted"`
}
