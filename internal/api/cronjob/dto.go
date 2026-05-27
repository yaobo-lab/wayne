package cronjob

import "wayne/internal/model"

// swagger:parameters reqListCronjob reqListCronjobTpl
type reqListCronjob struct {
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

// swagger:parameters reqCreateCronjob
type reqCreateCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cronjob content
	// in: body
	// required: true
	Body model.Cronjob `json:"body"`
}

// swagger:parameters reqDeleteCronjob
type reqDeleteCronjob struct {
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

// swagger:parameters reqUpdateCronjob
type reqUpdateCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the cronjob content
	// in: body
	// required: true
	Body model.Cronjob `json:"body"`
}

// swagger:parameters reqGetCronjob
type reqGetCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqGetNamesCronjob
type reqGetNamesCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersCronjob
type reqUpdateOrdersCronjob struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.Cronjob `json:"body"`
}

// swagger:parameters reqCreateCronjobTpl
type reqCreateCronjobTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the cronjob template content
	// in: body
	// required: true
	Body model.CronjobTemplate `json:"body"`
}

// swagger:parameters reqUpdateCronjobTpl
type reqUpdateCronjobTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the cronjob template content
	// in: body
	// required: true
	Body model.CronjobTemplate `json:"body"`
}

// swagger:parameters reqDeleteCronjobTpl
type reqDeleteCronjobTpl struct {
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

// swagger:parameters reqGetCronjobTpl
type reqGetCronjobTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
