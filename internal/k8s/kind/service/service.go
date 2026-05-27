package service

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	kapi "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"

	"wayne/internal/k8s/client"
	k8sDto "wayne/internal/k8s/dto"
	"wayne/internal/k8s/kind/endpoint"
	"wayne/internal/k8s/kind/event"
	"wayne/internal/k8s/kind/pod"
)

type ServiceDetail struct {
	ObjectMeta k8sDto.ObjectMeta `json:"objectMeta"`
	TypeMeta   k8sDto.TypeMeta   `json:"typeMeta"`

	// InternalEndpoint of all Kubernetes services that have the same label selector as connected Replication
	// Controller. Endpoints is DNS name merged with ports.
	InternalEndpoint k8sDto.Endpoint `json:"internalEndpoint"`

	// ExternalEndpoints of all Kubernetes services that have the same label selector as connected Replication
	// Controller. Endpoints is external IP address name merged with ports.
	ExternalEndpoints []k8sDto.Endpoint `json:"externalEndpoints"`

	// List of Endpoint obj. that are endpoints of this Service.
	EndpointList []endpoint.Endpoint `json:"endpointList"`

	// Label selector of the service.
	Selector map[string]string `json:"selector"`

	// Type determines how the service will be exposed.  Valid options: ClusterIP, NodePort, LoadBalancer
	Type v1.ServiceType `json:"type"`

	// ClusterIP is usually assigned by the master. Valid values are None, empty string (""), or
	// a valid IP address. None can be specified for headless services when proxying is not required
	ClusterIP string `json:"clusterIP"`

	// List of events related to this Service
	EventList []k8sDto.Event `json:"eventList"`

	// PodInfos represents list of pods status targeted by same label selector as this service.
	PodList []*v1.Pod `json:"podList"`

	// Show the value of the SessionAffinity of the Service.
	SessionAffinity v1.ServiceAffinity `json:"sessionAffinity"`
}

func GetServiceDetail(cli *kubernetes.Clientset, indexer *client.CacheFactory, namespace, name string) (*ServiceDetail, error) {

	serviceDate, err := cli.CoreV1().Services(namespace).Get(context.Background(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	endpoint, err := endpoint.GetServiceEndpointsFromCache(indexer, namespace, name)
	if err != nil {
		return nil, err
	}

	podList, err := pod.ListKubePod(indexer, namespace, serviceDate.Spec.Selector)
	if err != nil {
		return nil, err
	}

	eventList, err := event.GetPodsWarningEvents(indexer, podList)
	if err != nil {
		return nil, err
	}

	detail := toServiceDetail(serviceDate, eventList, podList, endpoint)

	return &detail, nil
}

func toServiceDetail(service *v1.Service, events []k8sDto.Event, pods []*v1.Pod, endpoints []endpoint.Endpoint) ServiceDetail {
	return ServiceDetail{
		ObjectMeta:        k8sDto.NewObjectMeta(service.ObjectMeta),
		TypeMeta:          k8sDto.NewTypeMeta(k8sDto.ResourceKind("service")),
		InternalEndpoint:  k8sDto.GetInternalEndpoint(service.Name, service.Namespace, service.Spec.Ports),
		ExternalEndpoints: k8sDto.GetExternalEndpoints(service),
		EndpointList:      endpoints,
		Selector:          service.Spec.Selector,
		ClusterIP:         service.Spec.ClusterIP,
		Type:              service.Spec.Type,
		EventList:         events,
		PodList:           pods,
		SessionAffinity:   service.Spec.SessionAffinity,
	}
}

func CreateOrUpdateService(cli *kubernetes.Clientset, service *kapi.Service) (*kapi.Service, error) {
	old, err := cli.CoreV1().Services(service.Namespace).Get(context.Background(), service.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.CoreV1().Services(service.Namespace).Create(context.Background(), service, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = service.Labels
	old.Spec.ExternalIPs = service.Spec.ExternalIPs
	old.Spec.Selector = service.Spec.Selector
	old.Spec.Ports = service.Spec.Ports

	return cli.CoreV1().Services(service.Namespace).Update(context.Background(), old, metaV1.UpdateOptions{})
}
