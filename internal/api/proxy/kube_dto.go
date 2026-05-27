package proxy

// swagger:parameters reqListNamespaceKubeProxy
type reqListNamespaceKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the kind
	// in: path
	Kind string `json:"kind"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// the labelSelector for list e.g. filter=app=infra-wayne,wayne-app=infra
	// in: query
	Filter string `json:"filter"`

	// labelSelector, ex. labelSelector=name=test
	// in: query
	LabelSelector string `json:"labelSelector"`

	// column sorted by, ex. sortby=-id, '-' representation desc, and sortby=id representation asc
	// in: query
	SortBy string `json:"sortby"`
}

// swagger:parameters reqListKubeProxy
type reqListKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the kind
	// in: path
	Kind string `json:"kind"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// the labelSelector for list e.g. filter=app=infra-wayne,wayne-app=infra
	// in: query
	Filter string `json:"filter"`

	// labelSelector, ex. labelSelector=name=test
	// in: query
	LabelSelector string `json:"labelSelector"`

	// column sorted by, ex. sortby=-id, '-' representation desc, and sortby=id representation asc
	// in: query
	SortBy string `json:"sortby"`
}

// swagger:parameters reqCreateNamespaceKubeProxy
type reqCreateNamespaceKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the namespace
	// in: path
	Namespace string `json:"namespace"`

	// the kind
	// in: path
	Kind string `json:"kind"`

	// the kubernetes resource
	// in: body
	Resource string `json:"resource"`
}

// swagger:parameters reqCreateKubeProxy
type reqCreateKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the kind
	// in: path
	Kind string `json:"kind"`

	// the kubernetes resource
	// in: body
	Resource string `json:"resource"`
}

// swagger:parameters reqGetNamespaceKubeProxy
type reqGetNamespaceKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`

	// the resource name
	// in: path
	Name string `json:"name"`
}

// swagger:parameters reqGetKubeProxy
type reqGetKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`

	// the resource name
	// in: path
	Name string `json:"name"`
}

// swagger:parameters reqUpdateNamespaceKubeProxy
type reqUpdateNamespaceKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`

	// the resource name
	// in: path
	Name string `json:"name"`

	// the kubernetes resource
	// in: body
	Resource string `json:"resource"`
}

// swagger:parameters reqUpdateKubeProxy
type reqUpdateKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`

	// the resource name
	// in: path
	Name string `json:"name"`

	// the kubernetes resource
	// in: body
	Resource string `json:"resource"`
}

// swagger:parameters reqDeleteNamespaceKubeProxy
type reqDeleteNamespaceKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`

	// the resource name
	// in: path
	Name string `json:"name"`

	// force to delete the resource from etcd.
	// in: body
	Force bool `json:"force"`
}

// swagger:parameters reqDeleteKubeProxy
type reqDeleteKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`

	// the resource name
	// in: path
	Name string `json:"name"`

	// force to delete the resource from etcd.
	// in: body
	Force bool `json:"force"`
}

// swagger:parameters reqGetNamesNamespaceKubeProxy
type reqGetNamesNamespaceKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`
}

// swagger:parameters reqGetNamesKubeProxy
type reqGetNamesKubeProxy struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster name
	// in: path
	Cluster string `json:"cluster"`

	// the resource kind
	// in: path
	Kind string `json:"kind"`
}
