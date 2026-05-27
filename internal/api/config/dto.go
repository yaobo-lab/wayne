package config

import "wayne/internal/model"

// swagger:parameters reqListBaseBaseConfig reqListConfig reqListSystemConfig
type reqListBaseBaseConfig struct {
}

// swagger:parameters reqCreateConfig
type reqCreateConfig struct {
	// the config content
	// in: body
	Body model.Config `json:"body"`
}

// swagger:parameters reqUpdateConfig
type reqUpdateConfig struct {
	// the id you want to update
	// in: path
	Id string `json:"id"`

	// the config content
	// in: body
	Body model.Config `json:"body"`
}

// swagger:parameters reqGetConfig reqDeleteConfig
type reqGetConfig struct {
	// the id
	// in: path
	Id string `json:"id"`
}
