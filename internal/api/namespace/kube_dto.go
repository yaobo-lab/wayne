package namespace

// swagger:parameters reqCreateKubeNamespace
type reqCreateKubeNamespace struct {
	// the name
	// in: path
	Name string `json:"name"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqResourcesKubeNamespace reqStatisticsKubeNamespace
type reqResourcesKubeNamespace struct {
	// the namespace id
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the app Name
	// in: query
	App string `json:"app"`
}
