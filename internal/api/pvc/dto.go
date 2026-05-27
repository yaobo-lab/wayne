package pvc

import "wayne/internal/model"

// swagger:parameters reqListPersistentVolumeClaim
type reqListPersistentVolumeClaim struct {
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

// swagger:parameters reqCreatePersistentVolumeClaim
type reqCreatePersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the deployment content
	// in: body
	// required: true
	Body model.PersistentVolumeClaim `json:"body"`
}

// swagger:parameters reqDeletePersistentVolumeClaim
type reqDeletePersistentVolumeClaim struct {
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

// swagger:parameters reqUpdatePersistentVolumeClaim
type reqUpdatePersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the daemonset content
	// in: body
	// required: true
	Body model.PersistentVolumeClaim `json:"body"`
}

// swagger:parameters reqGetPersistentVolumeClaim
type reqGetPersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqGetNamesPersistentVolumeClaim
type reqGetNamesPersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// is deleted,default false.
	// in: query
	Deleted string `json:"deleted"`
}

// swagger:parameters reqUpdateOrdersPersistentVolumeClaim
type reqUpdateOrdersPersistentVolumeClaim struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the body
	// in: body
	// required: true
	Body []model.PersistentVolumeClaim `json:"body"`
}

// swagger:parameters reqListPersistentVolumeClaimTpl
type reqListPersistentVolumeClaimTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`

	// pvc id
	// in: query
	PersistentVolumeClaimId int `json:"pvcId"`

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

// swagger:parameters reqCreatePersistentVolumeClaimTpl
type reqCreatePersistentVolumeClaimTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the pvc template content
	// in: body
	// required: true
	Body model.PersistentVolumeClaimTemplate `json:"body"`
}

// swagger:parameters reqUpdatePersistentVolumeClaimTpl
type reqUpdatePersistentVolumeClaimTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the pvc template content
	// in: body
	// required: true
	Body model.PersistentVolumeClaimTemplate `json:"body"`
}

// swagger:parameters reqDeletePersistentVolumeClaimTpl
type reqDeletePersistentVolumeClaimTpl struct {
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

// swagger:parameters reqGetPersistentVolumeClaimTpl
type reqGetPersistentVolumeClaimTpl struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}
