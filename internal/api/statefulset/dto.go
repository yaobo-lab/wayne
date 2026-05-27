package statefulset

import "wayne/internal/model"

// swagger:parameters reqListStatefulset
type reqListStatefulset struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

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

// swagger:parameters reqCreateStatefulset
type reqCreateStatefulset struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the statefulset content
	// in: body
	// required: true
	Body model.Statefulset `json:"body"`
}

// swagger:parameters reqDeleteStatefulset
type reqDeleteStatefulset struct {
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

// swagger:parameters reqGetStatefulset
type reqGetStatefulset struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateStatefulset
type reqUpdateStatefulset struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the statefulset content
	// in: body
	// required: true
	Body model.Statefulset `json:"body"`
}

// swagger:parameters reqGetNamesStatefulset
type reqGetNamesStatefulset struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersStatefulset
type reqUpdateOrdersStatefulset struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.Statefulset `json:"body"`
}

// swagger:parameters reqListStatefulsetTpl
type reqListStatefulsetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// statefulset id
	// in: query
	StatefulsetId int `json:"statefulsetId"`

	// only show online tpls,default false
	// in: query
	IsOnline bool `json:"isOnline"`

	// name filter
	// in: query
	Name string `json:"name"`

	// is deleted
	// in: query
	Deleted bool `json:"deleted"`
}

// swagger:parameters reqCreateStatefulsetTpl
type reqCreateStatefulsetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the statefulset template content
	// in: body
	// required: true
	Body model.StatefulsetTemplate `json:"body"`
}

// swagger:parameters reqUpdateStatefulsetTpl
type reqUpdateStatefulsetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the statefulset template content
	// in: body
	// required: true
	Body model.StatefulsetTemplate `json:"body"`
}

// swagger:parameters reqDeleteStatefulsetTpl
type reqDeleteStatefulsetTpl struct {
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

// swagger:parameters reqGetStatefulsetTpl
type reqGetStatefulsetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
