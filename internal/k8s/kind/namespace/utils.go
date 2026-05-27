package namespace

import (
	v1 "k8s.io/api/core/v1"

	k8sDto "wayne/internal/k8s/dto"
	"wayne/internal/k8s/kind/dataselector"
)

type Namespace struct {
	ObjectMeta k8sDto.ObjectMeta `json:"objectMeta"`
	Status     v1.NamespacePhase `json:"status"`
}

func toNamespace(namespace *v1.Namespace) *Namespace {
	result := &Namespace{
		ObjectMeta: k8sDto.NewObjectMeta(namespace.ObjectMeta),
	}
	result.Status = namespace.Status.Phase

	return result
}

type NamespaceCell Namespace

func (cell NamespaceCell) GetProperty(name dataselector.PropertyName) dataselector.ComparableValue {
	switch name {
	case dataselector.NameProperty:
		return dataselector.StdComparableString(cell.ObjectMeta.Name)
	case dataselector.CreationTimestampProperty:
		return dataselector.StdComparableTime(cell.ObjectMeta.CreationTimestamp.Time)
	case dataselector.NamespaceProperty:
		return dataselector.StdComparableString(cell.ObjectMeta.Namespace)
	default:
		// if name is not supported then just return a constant dummy value, sort will have no effect.
		return nil
	}
}
