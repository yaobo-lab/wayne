package ingress

import (
	"wayne/internal/k8s/kind/dataselector"
)

type IngressCell Ingress

func (self IngressCell) GetProperty(name dataselector.PropertyName) dataselector.ComparableValue {
	switch name {
	case dataselector.NameProperty:
		return dataselector.StdComparableString(self.ObjectMeta.Name)
	case dataselector.CreationTimestampProperty:
		return dataselector.StdComparableTime(self.ObjectMeta.CreationTimestamp.Time)
	case dataselector.NamespaceProperty:
		return dataselector.StdComparableString(self.ObjectMeta.Namespace)
	default:
		return nil
	}
}

func toCells(std []Ingress) []dataselector.DataCell {
	cells := make([]dataselector.DataCell, len(std))
	for i := range std {
		cells[i] = IngressCell(std[i])
	}
	return cells
}
