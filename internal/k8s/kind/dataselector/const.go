package dataselector

type PropertyName string

const (
	NameProperty              PropertyName = "name"
	CreationTimestampProperty PropertyName = "creationTimestamp"
	NamespaceProperty         PropertyName = "namespace"
	StatusProperty            PropertyName = "status"
	PodIPProperty             PropertyName = "podIP"
	NodeNameProperty          PropertyName = "nodeName"
)
