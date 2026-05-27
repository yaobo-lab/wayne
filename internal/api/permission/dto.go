package permission

import "wayne/internal/model"

// swagger:parameters reqListAppUser
type reqListAppUser struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`
}

// swagger:parameters reqCreateAppUser
type reqCreateAppUser struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the app user content
	// in: body
	// required: true
	Body model.AppUser `json:"body"`
}

// swagger:parameters reqGetAppUser
type reqGetAppUser struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateAppUser
type reqUpdateAppUser struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the app user content
	// in: body
	// required: true
	Body model.AppUser `json:"body"`
}

// swagger:parameters reqDeleteAppUser
type reqDeleteAppUser struct {
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

// swagger:parameters reqGetPermissionByAppAppUser
type reqGetPermissionByAppAppUser struct {
	// the appid
	// in: path
	AppId string `json:"appid"`

	// the id you want to delete
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqListGroup reqListPermission reqListUser
type reqListGroup struct {
	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`
}

// swagger:parameters reqCreateGroup
type reqCreateGroup struct {
	// the group content
	// in: body
	// required: true
	Body model.Group `json:"body"`
}

// swagger:parameters reqGetGroup reqGetPermission reqGetUser
type reqGetGroup struct {
	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateGroup
type reqUpdateGroup struct {
	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the group content
	// in: body
	// required: true
	Body model.Group `json:"body"`
}

// swagger:parameters reqDeleteGroup reqDeletePermission reqDeleteUser
type reqDeleteGroup struct {
	// the id you want to delete
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqListNamespaceUser
type reqListNamespaceUser struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the page current no
	// in: query
	PageNo int `json:"pageNo"`

	// the page size
	// in: query
	PageSize int `json:"pageSize"`
}

// swagger:parameters reqCreateNamespaceUser
type reqCreateNamespaceUser struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the namespace user content
	// in: body
	// required: true
	Body model.NamespaceUser `json:"body"`
}

// swagger:parameters reqGetNamespaceUser
type reqGetNamespaceUser struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the id you want to get
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqUpdateNamespaceUser
type reqUpdateNamespaceUser struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the namespace user content
	// in: body
	// required: true
	Body model.NamespaceUser `json:"body"`
}

// swagger:parameters reqDeleteNamespaceUser
type reqDeleteNamespaceUser struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the id you want to delete
	// in: path
	Id string `json:"id"`

	// is logical deletion,default true
	// in: query
	Logical bool `json:"logical"`
}

// swagger:parameters reqGetPermissionByNSNamespaceUser
type reqGetPermissionByNSNamespaceUser struct {
	// the namespaceid
	// in: path
	NamespaceId string `json:"namespaceid"`

	// the ns id
	// in: path
	Id string `json:"id"`
}

// swagger:parameters reqCreatePermission
type reqCreatePermission struct {
	// the permission content
	// in: body
	// required: true
	Body model.Permission `json:"body"`
}

// swagger:parameters reqUpdatePermission
type reqUpdatePermission struct {
	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the permission content
	// in: body
	// required: true
	Body model.Permission `json:"body"`
}

// swagger:parameters reqCreateUser
type reqCreateUser struct {
	// the user content
	// in: body
	// required: true
	Body model.User `json:"body"`
}

// swagger:parameters reqUpdateUser reqUpdateAdminUser
type reqUpdateUser struct {
	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the user content
	// in: body
	// required: true
	Body model.User `json:"body"`
}

// swagger:parameters reqResetPasswordUser
type reqResetPasswordUser struct {
	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the body
	// in: body
	// required: true
	Body struct {
		Id       int64  `json:"id"`
		Password string `json:"password"`
	} `json:"body"`
}

// swagger:parameters reqGetNamesUser reqUserStatisticsUser
type reqGetNamesUser struct {
}
