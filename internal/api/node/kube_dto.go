package node

// swagger:parameters reqGetKubeNode reqDeleteKubeNode reqGetLabelsKubeNode
type reqGetKubeNode struct {
	// the node name
	// in: path
	Name string `json:"name"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqUpdateKubeNode reqAddLabelKubeNode reqDeleteLabelKubeNode reqAddLabelsKubeNode reqDeleteLabelsKubeNode reqSetTaintKubeNode reqDeleteTaintKubeNode
type reqUpdateKubeNode struct {
	// the node name
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

// swagger:parameters reqListKubeNode
type reqListKubeNode struct {
	// the cluster name
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqNodeStatisticsKubeNode
type reqNodeStatisticsKubeNode struct {
	// the cluster name
	// in: query
	Cluster string `json:"cluster"`
}
