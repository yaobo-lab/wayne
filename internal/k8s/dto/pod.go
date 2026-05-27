package dto

type PodInfo struct {
	Current   int32   `json:"current"`
	Desired   int32   `json:"desired"`
	Running   int32   `json:"running"`
	Pending   int32   `json:"pending"`
	Failed    int32   `json:"failed"`
	Succeeded int32   `json:"succeeded"`
	Warnings  []Event `json:"warnings,omitempty"`
}
