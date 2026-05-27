package apikey

import (
	"wayne/internal/model"
	common "wayne/pkg/dto"
)

// swagger:parameters reqListAppApiKey
type reqListAppApiKey struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// is deleted, default list all.
	// in: query
	Deleted bool `json:"deleted"`
}

// swagger:parameters reqListNamespaceApiKey
type reqListNamespaceApiKey struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// is deleted, default list all.
	// in: query
	Deleted bool `json:"deleted"`
}

// success
// swagger:response
type respListAPIKey struct {
	// in: body
	Body struct {
		Data struct {
			common.Page
			List []model.APIKey `json:"list"`
		} `json:"data"`
	}
}

// swagger:parameters reqCreateAppApiKey
type reqCreateAppApiKey struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the APIKey content
	// in: body
	// required: true
	Body struct {
		model.APIKey
	} `json:"body"`
}

// swagger:parameters reqCreateNamespaceApiKey
type reqCreateNamespaceApiKey struct {
	// the namespaceid
	// in: path
	NamespaceId int `json:"namespaceid"`

	// the APIKey content
	// in: body
	// required: true
	Body struct {
		model.APIKey
	} `json:"body"`
}

// swagger:parameters reqUpdateAppApiKey
type reqUpdateAppApiKey struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the body
	// in: body
	// required: true
	Body struct {
		model.APIKey
	} `json:"body"`
}

// swagger:parameters reqUpdateNamespaceApiKey
type reqUpdateNamespaceApiKey struct {
	// the namespaceid
	// in: path
	NamespaceId int `json:"namespaceid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the body
	// in: body
	// required: true
	Body struct {
		model.APIKey
	} `json:"body"`
}

// swagger:parameters reqDeleteAppApiKey
type reqDeleteAppApiKey struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to delete
	// in: path
	Id string `json:"id"`

	// is logical deletion,default true
	// in: query
	Logical bool `json:"logical"`
}

// swagger:parameters reqDeleteNamespaceApiKey
type reqDeleteNamespaceApiKey struct {
	// the namespaceid
	// in: path
	NamespaceId int `json:"namespaceid"`

	// the id you want to delete
	// in: path
	Id string `json:"id"`

	// is logical deletion,default true
	// in: query
	Logical bool `json:"logical"`
}

// swagger:parameters reqGetAppApiKey
type reqGetAppApiKey struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqGetNamespaceApiKey
type reqGetNamespaceApiKey struct {
	// the namespaceid
	// in: path
	NamespaceId int `json:"namespaceid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
