package service

import "wayne/internal/model"

// swagger:parameters reqListService
type reqListService struct {
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

// swagger:parameters reqCreateService
type reqCreateService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the secret content
	// in: body
	// required: true
	Body model.Service `json:"body"`
}

// swagger:parameters reqDeleteService
type reqDeleteService struct {
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

// swagger:parameters reqUpdateService
type reqUpdateService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the secret content
	// in: body
	// required: true
	Body model.Service `json:"body"`
}

// swagger:parameters reqGetService
type reqGetService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqGetNamesService
type reqGetNamesService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersService
type reqUpdateOrdersService struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.Service `json:"body"`
}

// swagger:parameters reqListServiceTpl
type reqListServiceTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// secret id
	// in: query
	ServiceId int `json:"secretId"`

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

// swagger:parameters reqCreateServiceTpl
type reqCreateServiceTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the secret template content
	// in: body
	// required: true
	Body model.ServiceTemplate `json:"body"`
}

// swagger:parameters reqUpdateServiceTpl
type reqUpdateServiceTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the secret template content
	// in: body
	// required: true
	Body model.ServiceTemplate `json:"body"`
}

// swagger:parameters reqDeleteServiceTpl
type reqDeleteServiceTpl struct {
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

// swagger:parameters reqGetServiceTpl
type reqGetServiceTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
