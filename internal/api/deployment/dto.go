package deployment

import "wayne/internal/model"

// swagger:parameters reqListDeployment
type reqListDeployment struct {
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

// swagger:parameters reqCreateDeployment
type reqCreateDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the deployment content
	// in: body
	// required: true
	Body model.Deployment `json:"body"`
}

// swagger:parameters reqDeleteDeployment
type reqDeleteDeployment struct {
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

// swagger:parameters reqGetDeployment
type reqGetDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateDeployment
type reqUpdateDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the deployment content
	// in: body
	// required: true
	Body model.Deployment `json:"body"`
}

// swagger:parameters reqGetNamesDeployment
type reqGetNamesDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersDeployment
type reqUpdateOrdersDeployment struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.Deployment `json:"body"`
}

// swagger:parameters reqListDeploymentTpl
type reqListDeploymentTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// deployment id
	// in: query
	DeploymentId int `json:"deploymentId"`

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

// swagger:parameters reqCreateDeploymentTpl
type reqCreateDeploymentTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the deployment template content
	// in: body
	// required: true
	Body model.DeploymentTemplate `json:"body"`
}

// swagger:parameters reqUpdateDeploymentTpl
type reqUpdateDeploymentTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the deployment template content
	// in: body
	// required: true
	Body model.DeploymentTemplate `json:"body"`
}

// swagger:parameters reqDeleteDeploymentTpl
type reqDeleteDeploymentTpl struct {
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

// swagger:parameters reqGetDeploymentTpl
type reqGetDeploymentTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
