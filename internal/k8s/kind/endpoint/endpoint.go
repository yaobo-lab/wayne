package endpoint

import (
	"wayne/internal/k8s/client"
	k8sDto "wayne/internal/k8s/dto"

	v1 "k8s.io/api/core/v1"
)

type Endpoint struct {
	ObjectMeta k8sDto.ObjectMeta `json:"objectMeta"`
	TypeMeta   k8sDto.TypeMeta   `json:"typeMeta"`
	Host       string            `json:"host"`
	NodeName   *string           `json:"nodeName"`
	Ready      bool              `json:"ready"`
	Ports      []v1.EndpointPort `json:"ports"`
}

func GetServiceEndpointsFromCache(cache *client.CacheFactory, namespace, name string) ([]Endpoint, error) {
	endpoint, err := cache.EndpointLister().Endpoints(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	return toEndpointList(endpoint), nil
}

func toEndpointList(endpoint *v1.Endpoints) (list []Endpoint) {
	for _, subSets := range endpoint.Subsets {
		for _, address := range subSets.Addresses {
			list = append(list, *toEndpoint(address, subSets.Ports, true))
		}
		for _, notReadyAddress := range subSets.NotReadyAddresses {
			list = append(list, *toEndpoint(notReadyAddress, subSets.Ports, false))
		}
	}

	return
}

func toEndpoint(address v1.EndpointAddress, ports []v1.EndpointPort, ready bool) *Endpoint {
	return &Endpoint{
		TypeMeta: k8sDto.NewTypeMeta(k8sDto.ResourceKind("endpoint")),
		Host:     address.IP,
		Ports:    ports,
		Ready:    ready,
		NodeName: address.NodeName,
	}
}
