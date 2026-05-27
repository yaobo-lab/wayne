package daemonset

import "wayne/internal/model"

// swagger:parameters reqListDaemonSet
type reqListDaemonSet struct {
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

// swagger:parameters reqCreateDaemonSet
type reqCreateDaemonSet struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the daemonset content
	// in: body
	// required: true
	Body model.DaemonSet `json:"body"`
}

// swagger:parameters reqDeleteDaemonSet
type reqDeleteDaemonSet struct {
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

// swagger:parameters reqGetDaemonSet
type reqGetDaemonSet struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateDaemonSet
type reqUpdateDaemonSet struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the daemonset content
	// in: body
	// required: true
	Body model.DaemonSet `json:"body"`
}

// swagger:parameters reqGetNamesDaemonSet
type reqGetNamesDaemonSet struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersDaemonSet
type reqUpdateOrdersDaemonSet struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.DaemonSet `json:"body"`
}

// swagger:parameters reqListDaemonSetTpl
type reqListDaemonSetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// daemonSet id
	// in: query
	DaemonSetId int `json:"daemonSetId"`

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

// swagger:parameters reqCreateDaemonSetTpl
type reqCreateDaemonSetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the daemonset template content
	// in: body
	// required: true
	Body model.DaemonSetTemplate `json:"body"`
}

// swagger:parameters reqUpdateDaemonSetTpl
type reqUpdateDaemonSetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the daemonset template content
	// in: body
	// required: true
	Body model.DaemonSetTemplate `json:"body"`
}

// swagger:parameters reqDeleteDaemonSetTpl
type reqDeleteDaemonSetTpl struct {
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

// swagger:parameters reqGetDaemonSetTpl
type reqGetDaemonSetTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
