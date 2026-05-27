package crd

// swagger:parameters reqListKubeCRD
type reqListKubeCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`
}

// swagger:parameters reqCreateKubeCRD
type reqCreateKubeCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}

// swagger:parameters reqGetKubeCRD reqDeleteKubeCRD
type reqGetKubeCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the name
	// in: path
	Name string `json:"name"`
}

// swagger:parameters reqUpdateKubeCRD
type reqUpdateKubeCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the name
	// in: path
	Name string `json:"name"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}

// swagger:parameters reqListNamespaceKubeCustomCRD
type reqListNamespaceKubeCustomCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the group
	// in: path
	Group string `json:"group"`

	// the version
	// in: path
	Version string `json:"version"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the kind
	// in: path
	Kind string `json:"kind"`
}

// swagger:parameters reqListKubeCustomCRD
type reqListKubeCustomCRD struct {
	// The appid
	// in: path
	AppId string `json:"appid"`

	// The cluster
	// in: path
	Cluster string `json:"cluster"`

	// The group
	// in: path
	Group string `json:"group"`

	// The version
	// in: path
	Version string `json:"version"`

	// The kind
	// in: path
	Kind string `json:"kind"`
}

// swagger:parameters reqCreateNamespaceKubeCustomCRD
type reqCreateNamespaceKubeCustomCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the group
	// in: path
	Group string `json:"group"`

	// the version
	// in: path
	Version string `json:"version"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the kind
	// in: path
	Kind string `json:"kind"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}

// swagger:parameters reqCreateKubeCustomCRD
type reqCreateKubeCustomCRD struct {
	// The appid
	// in: path
	AppId string `json:"appid"`

	// The cluster
	// in: path
	Cluster string `json:"cluster"`

	// The group
	// in: path
	Group string `json:"group"`

	// The version
	// in: path
	Version string `json:"version"`

	// The kind
	// in: path
	Kind string `json:"kind"`

	// The body
	// in: body
	// required: true
	Body string `json:"body"`
}

// swagger:parameters reqGetNamespaceKubeCustomCRD reqDeleteNamespaceKubeCustomCRD
type reqGetNamespaceKubeCustomCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the group
	// in: path
	Group string `json:"group"`

	// the version
	// in: path
	Version string `json:"version"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the kind
	// in: path
	Kind string `json:"kind"`

	// the name
	// in: path
	Name string `json:"name"`
}

// swagger:parameters reqGetKubeCustomCRD reqDeleteKubeCustomCRD
type reqGetKubeCustomCRD struct {
	// The appid
	// in: path
	AppId string `json:"appid"`

	// The cluster
	// in: path
	Cluster string `json:"cluster"`

	// The group
	// in: path
	Group string `json:"group"`

	// The version
	// in: path
	Version string `json:"version"`

	// The kind
	// in: path
	Kind string `json:"kind"`

	// the name
	// in: path
	Name string `json:"name"`
}

// swagger:parameters reqUpdateNamespaceKubeCustomCRD
type reqUpdateNamespaceKubeCustomCRD struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cluster
	// in: path
	Cluster string `json:"cluster"`

	// the group
	// in: path
	Group string `json:"group"`

	// the version
	// in: path
	Version string `json:"version"`

	// the namespace name
	// in: path
	Namespace string `json:"namespace"`

	// the kind
	// in: path
	Kind string `json:"kind"`

	// the name
	// in: path
	Name string `json:"name"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}

// swagger:parameters reqUpdateKubeCustomCRD
type reqUpdateKubeCustomCRD struct {
	// The appid
	// in: path
	AppId string `json:"appid"`

	// The cluster
	// in: path
	Cluster string `json:"cluster"`

	// The group
	// in: path
	Group string `json:"group"`

	// The version
	// in: path
	Version string `json:"version"`

	// The kind
	// in: path
	Kind string `json:"kind"`

	// the name
	// in: path
	Name string `json:"name"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}
