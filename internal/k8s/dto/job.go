package dto

type JobInfo struct {
	Active    int32 `json:"active,omitempty"`
	Succeeded int32 `json:"succeeded,omitempty"`
	Failed    int32 `json:"failed,omitempty"`
}
