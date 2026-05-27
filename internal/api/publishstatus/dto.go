package publishstatus

// swagger:parameters reqListPublishStatus
type reqListPublishStatus struct {
	// the ResourceId id,e.g. deployment id
	// in: query
	// required: true
	ResourceId int `json:"resourceId"`

	// the ResourceId type, DEPLOYMENT 0 SERVICE 1 CONFIGMAP 2 SECRET 3
	// in: query
	// required: true
	Type int `json:"type"`
}

// swagger:parameters reqDeletePublishStatus
type reqDeletePublishStatus struct {
	// the id you want to delete
	// in: path
	Id int `json:"id"`
}
