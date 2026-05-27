package namespace

import "wayne/internal/model"

// swagger:parameters reqListNamespace
type reqListNamespace struct {
	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// name filter
	// in: query
	Name string `json:"name"`

	// is deleted, default list all
	// in: query
	Deleted bool `json:"deleted"`
}

// swagger:parameters reqCreateNamespace
type reqCreateNamespace struct {
	// the namespace content
	// in: body
	// required: true
	Body model.Namespace `json:"body"`
}

// swagger:parameters reqDeleteNamespace
type reqDeleteNamespace struct {
	// the id you want to delete
	// in: path
	Id string `json:"id"`

	// is logical deletion,default true
	// in: query
	Logical bool `json:"logical"`
}

// swagger:parameters reqUpdateNamespace
type reqUpdateNamespace struct {
	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the deployment content
	// in: body
	// required: true
	Body model.Namespace `json:"body"`
}

// swagger:parameters reqGetNamespace
type reqGetNamespace struct {
	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqStatisticsNamespace reqHistoryNamespace
type reqStatisticsNamespace struct {
	// the id
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqInitDefaultNamespace
type reqInitDefaultNamespace struct {
}

// swagger:parameters reqGetNamesNamespace
type reqGetNamesNamespace struct {
	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}
