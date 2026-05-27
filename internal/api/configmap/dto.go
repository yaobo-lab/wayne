package configmap

import "wayne/internal/model"

// swagger:parameters reqListConfigMap reqListConfigMapTpl
type reqListConfigMap struct {
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

// swagger:parameters reqCreateConfigMap
type reqCreateConfigMap struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the config map content
	// in: body
	// required: true
	Body model.ConfigMap `json:"body"`
}

// swagger:parameters reqDeleteConfigMap
type reqDeleteConfigMap struct {
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

// swagger:parameters reqUpdateConfigMap
type reqUpdateConfigMap struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the config map content
	// in: body
	// required: true
	Body model.ConfigMap `json:"body"`
}

// swagger:parameters reqGetConfigMap
type reqGetConfigMap struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqGetNamesConfigMap
type reqGetNamesConfigMap struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersConfigMap
type reqUpdateOrdersConfigMap struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.ConfigMap `json:"body"`
}

// swagger:parameters reqCreateConfigMapTpl
type reqCreateConfigMapTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the config map template content
	// in: body
	// required: true
	Body model.ConfigMapTemplate `json:"body"`
}

// swagger:parameters reqUpdateConfigMapTpl
type reqUpdateConfigMapTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the config map template content
	// in: body
	// required: true
	Body model.ConfigMapTemplate `json:"body"`
}

// swagger:parameters reqDeleteConfigMapTpl
type reqDeleteConfigMapTpl struct {
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

// swagger:parameters reqGetConfigMapTpl
type reqGetConfigMapTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
