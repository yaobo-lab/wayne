package ingress

import "wayne/internal/model"

// swagger:parameters reqListIngress
type reqListIngress struct {
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

// swagger:parameters reqCreateIngress
type reqCreateIngress struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the ingress content
	// in: body
	// required: true
	Body model.Ingress `json:"body"`
}

// swagger:parameters reqDeleteIngress
type reqDeleteIngress struct {
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

// swagger:parameters reqGetIngress
type reqGetIngress struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateIngress
type reqUpdateIngress struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the ingress content
	// in: body
	// required: true
	Body model.Ingress `json:"body"`
}

// swagger:parameters reqGetNamesIngress
type reqGetNamesIngress struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersIngress
type reqUpdateOrdersIngress struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.Ingress `json:"body"`
}

// swagger:parameters reqListIngressTpl
type reqListIngressTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// ingress id
	// in: query
	IngressId int `json:"ingressId"`

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

// swagger:parameters reqCreateIngressTpl
type reqCreateIngressTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the ingress template content
	// in: body
	// required: true
	Body model.IngressTemplate `json:"body"`
}

// swagger:parameters reqGetIngressTpl
type reqGetIngressTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateIngressTpl
type reqUpdateIngressTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the ingress template content
	// in: body
	// required: true
	Body model.IngressTemplate `json:"body"`
}

// swagger:parameters reqDeleteIngressTpl
type reqDeleteIngressTpl struct {
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
