package statefulset

import (
	"context"

	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"wayne/internal/k8s/client"
	api "wayne/internal/k8s/client"
	k8sDto "wayne/internal/k8s/dto"
	"wayne/internal/k8s/kind/event"
	"wayne/internal/k8s/kind/pod"
	"wayne/pkg/maps"
)

type Statefulset struct {
	ObjectMeta k8sDto.ObjectMeta `json:"objectMeta"`
	Pods       k8sDto.PodInfo    `json:"pods"`
}

// GetStatefulsetResource get StatefulSet resource statistics
func GetStatefulsetResource(cli client.ResourceHandler, statefulSet *appsV1.StatefulSet) (*k8sDto.ResourceList, error) {
	obj, err := cli.Get(api.ResourceNameStatefulSet, statefulSet.Namespace, statefulSet.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			return k8sDto.StatefulsetResourceList(statefulSet), nil
		}
		return nil, err
	}
	old := obj.(*appsV1.StatefulSet)
	oldResourceList := k8sDto.StatefulsetResourceList(old)
	newResourceList := k8sDto.StatefulsetResourceList(statefulSet)

	return &k8sDto.ResourceList{
		Cpu:    newResourceList.Cpu - oldResourceList.Cpu,
		Memory: newResourceList.Memory - oldResourceList.Memory,
	}, nil
}

func CreateOrUpdateStatefulset(cli *kubernetes.Clientset, statefulSet *appsV1.StatefulSet) (*appsV1.StatefulSet, error) {
	old, err := cli.AppsV1().StatefulSets(statefulSet.Namespace).Get(context.Background(), statefulSet.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.AppsV1().StatefulSets(statefulSet.Namespace).Create(context.Background(), statefulSet, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = maps.MergeLabels(old.Labels, statefulSet.Labels)
	oldTemplateLabels := old.Spec.Template.Labels
	old.Spec = statefulSet.Spec
	old.Spec.Template.Labels = maps.MergeLabels(oldTemplateLabels, statefulSet.Spec.Template.Labels)

	return cli.AppsV1().StatefulSets(statefulSet.Namespace).Update(context.Background(), old, metaV1.UpdateOptions{})
}

func GetStatefulsetDetail(cli *kubernetes.Clientset, indexer *client.CacheFactory, name, namespace string) (*Statefulset, error) {
	statefulSet, err := cli.AppsV1().StatefulSets(namespace).Get(context.Background(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	result := &Statefulset{
		ObjectMeta: k8sDto.NewObjectMeta(statefulSet.ObjectMeta),
	}

	podInfo := k8sDto.PodInfo{}
	podInfo.Current = statefulSet.Status.ReadyReplicas
	podInfo.Desired = *statefulSet.Spec.Replicas

	pods, err := pod.ListKubePod(indexer, namespace, statefulSet.Spec.Template.Labels)
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
