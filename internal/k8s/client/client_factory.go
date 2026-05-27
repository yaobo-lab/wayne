package client

import (
	appsV1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

// 根据资源 类型，获取k8s client
func (h *resourceHandler) getClientByGroupVersion(groupVersion schema.GroupVersionResource) rest.Interface {

	switch groupVersion.Group {

	case corev1.GroupName:
		return h.client.CoreV1().RESTClient()

	case appsV1.GroupName:
		return h.client.AppsV1().RESTClient()

	case networkingv1.GroupName:
		return h.client.NetworkingV1().RESTClient()

	case autoscalingv1.GroupName:
		return h.client.AutoscalingV1().RESTClient()

	case batchv1.GroupName:
		return h.client.BatchV1().RESTClient()

	case storagev1.GroupName:
		return h.client.StorageV1().RESTClient()

	case rbacv1.GroupName:
		return h.client.RbacV1().RESTClient()

	default:
		return h.client.CoreV1().RESTClient()
	}
}
