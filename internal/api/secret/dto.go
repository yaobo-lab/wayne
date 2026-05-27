package secret

import "wayne/internal/model"

// swagger:parameters reqListSecret
type reqListSecret struct {
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

// swagger:parameters reqCreateSecret
type reqCreateSecret struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the secret content
	// in: body
	// required: true
	Body model.Secret `json:"body"`
}

// swagger:parameters reqDeleteSecret
type reqDeleteSecret struct {
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

// swagger:parameters reqUpdateSecret
type reqUpdateSecret struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the secret content
	// in: body
	// required: true
	Body model.Secret `json:"body"`
}

// swagger:parameters reqGetSecret
type reqGetSecret struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqGetNamesSecret
type reqGetNamesSecret struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersSecret
type reqUpdateOrdersSecret struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.Secret `json:"body"`
}

// swagger:parameters reqListSecretTpl
type reqListSecretTpl struct {
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
	SecretId int `json:"secretId"`

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

// swagger:parameters reqCreateSecretTpl
type reqCreateSecretTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the secret template content
	// in: body
	// required: true
	Body model.SecretTemplate `json:"body"`
}

// swagger:parameters reqUpdateSecretTpl
type reqUpdateSecretTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the secret template content
	// in: body
	// required: true
	Body model.SecretTemplate `json:"body"`
}

// swagger:parameters reqDeleteSecretTpl
type reqDeleteSecretTpl struct {
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

// swagger:parameters reqGetSecretTpl
type reqGetSecretTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
