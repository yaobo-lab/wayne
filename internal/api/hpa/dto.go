package hpa

import "wayne/internal/model"

// swagger:parameters reqListHPA
type reqListHPA struct {
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

// swagger:parameters reqCreateHPA
type reqCreateHPA struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the hpa content
	// in: body
	// required: true
	Body model.HPA `json:"body"`
}

// swagger:parameters reqDeleteHPA
type reqDeleteHPA struct {
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

// swagger:parameters reqGetHPA
type reqGetHPA struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateHPA
type reqUpdateHPA struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the hpa content
	// in: body
	// required: true
	Body model.HPA `json:"body"`
}

// swagger:parameters reqGetNamesHPA
type reqGetNamesHPA struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersHPA
type reqUpdateOrdersHPA struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.HPA `json:"body"`
}

// swagger:parameters reqListHPATpl
type reqListHPATpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// hpa id
	// in: query
	HPAId int `json:"hpaId"`

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

// swagger:parameters reqCreateHPATpl
type reqCreateHPATpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the hpa template content
	// in: body
	// required: true
	Body model.HPATemplate `json:"body"`
}

// swagger:parameters reqGetHPATpl
type reqGetHPATpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateHPATpl
type reqUpdateHPATpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the hpa template content
	// in: body
	// required: true
	Body model.HPATemplate `json:"body"`
}

// swagger:parameters reqDeleteHPATpl
type reqDeleteHPATpl struct {
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
