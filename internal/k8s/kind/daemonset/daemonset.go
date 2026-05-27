package daemonset

import (
	"context"

	"wayne/internal/k8s/client"
	"wayne/internal/k8s/dto"
	"wayne/internal/k8s/kind/event"
	"wayne/internal/k8s/kind/pod"
	"wayne/pkg/maps"

	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DaemonSet struct {
	ObjectMeta dto.ObjectMeta `json:"objectMeta"`
	Pods       dto.PodInfo    `json:"pods"`
}

func CreateOrUpdateDaemonSet(cli *kubernetes.Clientset, daemonSet *appsV1.DaemonSet) (*appsV1.DaemonSet, error) {
	ctx := context.Background()

	old, err := cli.AppsV1().DaemonSets(daemonSet.Namespace).Get(ctx, daemonSet.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.AppsV1().DaemonSets(daemonSet.Namespace).Create(ctx, daemonSet, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = maps.MergeLabels(old.Labels, daemonSet.Labels)
	oldTemplateLabels := old.Spec.Template.Labels
	old.Spec = daemonSet.Spec
	old.Spec.Template.Labels = maps.MergeLabels(oldTemplateLabels, daemonSet.Spec.Template.Labels)

	return cli.AppsV1().DaemonSets(daemonSet.Namespace).Update(ctx, old, metaV1.UpdateOptions{})
}

func GetDaemonSetDetail(cli *kubernetes.Clientset, indexer *client.CacheFactory, name, namespace string) (*DaemonSet, error) {

	ctx := context.Background()

	daemonSet, err := cli.AppsV1().DaemonSets(namespace).Get(ctx, name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	result := &DaemonSet{
		ObjectMeta: dto.NewObjectMeta(daemonSet.ObjectMeta),
	}

	podInfo := dto.PodInfo{}
	podInfo.Current = daemonSet.Status.NumberAvailable
	podInfo.Desired = daemonSet.Status.DesiredNumberScheduled

	pods, err := pod.ListKubePod(indexer, namespace, daemonSet.Spec.Template.Labels)
	if err != nil {
		return nil, err
	}

	podInfo.Warnings, err = event.GetPodsWarningEvents(indexer, pods)
	if err != nil {
		return nil, err
	}

	result.Pods = podInfo

	return result, nil
}

func DeleteDaemonSet(cli *kubernetes.Clientset, name, namespace string) error {
	deletionPropagation := metaV1.DeletePropagationBackground
	return cli.AppsV1().
		DaemonSets(namespace).
		Delete(context.Background(), name, metaV1.DeleteOptions{PropagationPolicy: &deletionPropagation})
}
