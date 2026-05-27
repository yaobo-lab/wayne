package hpa

import (
	"context"

	autoscaling "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateOrUpdateHPA(c *kubernetes.Clientset, hpa *autoscaling.HorizontalPodAutoscaler) (*HPA, error) {
	ctx := context.Background()

	old, err := c.AutoscalingV1().HorizontalPodAutoscalers(hpa.Namespace).Get(ctx, hpa.Name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			kubeHPA, err := c.AutoscalingV1().HorizontalPodAutoscalers(hpa.Namespace).Create(ctx, hpa, v1.CreateOptions{})
			if err != nil {
				return nil, err
			}
			return toHPA(kubeHPA), nil
		}
		return nil, err
	}
	hpa.Spec.DeepCopyInto(&old.Spec)
	kubeHPA, err := c.AutoscalingV1().HorizontalPodAutoscalers(hpa.Namespace).Update(ctx, old, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return toHPA(kubeHPA), nil
}
