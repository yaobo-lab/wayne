package ingress

import (
	k8sDto "wayne/internal/k8s/dto"

	networkingv1 "k8s.io/api/networking/v1"
)

type Ingress struct {
	k8sDto.ObjectMeta `json:"objectMeta"`
	k8sDto.TypeMeta   `json:"typeMeta"`
	Endpoints         []k8sDto.Endpoint `json:"endpoints"`
}

func getEndpoints(ingress *networkingv1.Ingress) []k8sDto.Endpoint {
	endpoints := make([]k8sDto.Endpoint, 0)
	if len(ingress.Status.LoadBalancer.Ingress) > 0 {
		for _, status := range ingress.Status.LoadBalancer.Ingress {
			endpoint := k8sDto.Endpoint{Host: status.IP}
			endpoints = append(endpoints, endpoint)
		}
	}
	return endpoints
}

func toIngress(ingress *networkingv1.Ingress) *Ingress {
	modelIngress := &Ingress{
		ObjectMeta: k8sDto.NewObjectMeta(ingress.ObjectMeta),
		TypeMeta:   k8sDto.NewTypeMeta("ingress"),
		Endpoints:  getEndpoints(ingress),
	}
	return modelIngress
}
