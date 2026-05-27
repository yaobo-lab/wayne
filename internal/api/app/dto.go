package app

import "wayne/internal/model"

// swagger:parameters reqListApp
type reqListApp struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// is starred app.default not star
	// in: query
	Starred bool `json:"starred"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// name filter
	// in: query
	Name string `json:"name"`

	// is deleted, default list all.
	// in: query
	Deleted bool `json:"deleted"`
}

// swagger:parameters reqCreateApp
type reqCreateApp struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the app content
	// in: body
	// required: true
	Body struct {
		model.App
	} `json:"body"`
}

// swagger:parameters reqUpdateApp
type reqUpdateApp struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the id you want to update
	// in: path
	Id int `json:"id"`

	// the body
	// in: body
	// required: true
	Body struct {
		model.App
	} `json:"body"`
}

// swagger:parameters reqDeleteApp
type reqDeleteApp struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the id you want to delete
	// in: path
	Id int `json:"id"`

	// is logical deletion,default true
	// in: query
	Logical bool `json:"logical"`
}

// swagger:parameters reqGetApp
type reqGetApp struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the id you want to get
	// in: path
	Id int `json:"id"`
}

// swagger:parameters reqGetNamesApp
type reqGetNamesApp struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// is deleted,default false.
	// in: query
	Deleted bool `json:"deleted"`
}

// swagger:parameters reqAppStatisticsApp
type reqAppStatisticsApp struct {
}
