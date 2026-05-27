package auth

// swagger:parameters reqCurrentUserAuth reqLogoutAuth
type reqCurrentUserAuth struct{}

// swagger:parameters reqLoginAuth
type reqLoginAuth struct {
	// the type
	// in: path
	Type string `json:"type"`

	// the name
	// in: path
	Name string `json:"name"`

	// the body
	// in: body
	// required: true
	Body string `json:"body"`
}
