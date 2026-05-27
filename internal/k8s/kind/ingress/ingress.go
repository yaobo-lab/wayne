package ingress

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateOrUpdateIngress(c *kubernetes.Clientset, ingress *networkingv1.Ingress) (*Ingress, error) {

	ctx := context.Background()

	old, err := c.NetworkingV1().Ingresses(ingress.Namespace).Get(ctx, ingress.Name, metaV1.GetOptions{})

	if err != nil {
		if errors.IsNotFound(err) {
			kubeIngress, err := c.NetworkingV1().Ingresses(ingress.Namespace).Create(ctx, ingress, metaV1.CreateOptions{})
			if err != nil {
				return nil, err
			}
			return toIngress(kubeIngress), nil
		}
		return nil, err
	}

	// ingress.Spec.DeepCopyInto(&old.Spec)
	// also need update Labels、Annotations、Spec
	old.Labels = ingress.Labels
	old.Annotations = ingress.Annotations
	old.Spec = ingress.Spec

	kubeIngress, err := c.NetworkingV1().Ingresses(ingress.Namespace).Update(ctx, old, metaV1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return toIngress(kubeIngress), nil
}

func GetIngressDetail(c *kubernetes.Clientset, name, namespace string) (*Ingress, error) {

	ctx := context.Background()

	ingress, err := c.NetworkingV1().Ingresses(namespace).Get(ctx, name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return toIngress(ingress), nil
}

func GetIngress(c *kubernetes.Clientset, name, namespace string) (ingress *networkingv1.Ingress, err error) {
	ingress, err = c.NetworkingV1().Ingresses(namespace).Get(context.Background(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return
}

func DeleteIngress(c *kubernetes.Clientset, name, namespace string) error {
	return c.NetworkingV1().Ingresses(namespace).Delete(context.Background(), name, metaV1.DeleteOptions{})
}

func GetIngressList(cli *kubernetes.Clientset, namespace string, opts metaV1.ListOptions) (list []*Ingress, err error) {
	kubeIngressList, err := cli.NetworkingV1().Ingresses(namespace).List(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	for _, kubeIngress := range kubeIngressList.Items {
		list = append(list, toIngress(&kubeIngress))
	}
	return
}
