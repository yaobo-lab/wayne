package appstarred

import "wayne/internal/model"

// swagger:parameters reqCreateAppStarred
type reqCreateAppStarred struct {
	// the app content
	// in: body
	// required: true
	Body model.AppStarred `json:"body"`
}

// swagger:parameters reqDeleteAppStarred
type reqDeleteAppStarred struct {
	// the appId you want to delete
	// in: path
	AppId string `json:"appId"`
}
