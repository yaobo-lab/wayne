package namespace

import (
	"context"
	"wayne/internal/k8s/client"
	api "wayne/internal/k8s/client"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	k8sDto "wayne/internal/k8s/dto"
	util "wayne/pkg"
)

func CreateNotExitNamespace(cli *kubernetes.Clientset, ns *v1.Namespace) (*v1.Namespace, error) {
	_, err := cli.CoreV1().Namespaces().Get(context.Background(), ns.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.CoreV1().Namespaces().Create(context.Background(), ns, metaV1.CreateOptions{})
		}
		return nil, err
	}
	return nil, nil
}

// ResourcesUsageByNamespace Count resource usage for a namespace
func ResourcesUsageByNamespace(cli client.ResourceHandler, namespace, selector string) (*k8sDto.ResourceList, error) {
	objs, err := cli.List(api.ResourceNamePod, namespace, selector)
	if err != nil {
		return nil, err
	}
	var cpuUsage, memoryUsage int64
	for _, obj := range objs {
		pod := obj.(*v1.Pod)
		if pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
			continue
		}
		resourceList := k8sDto.ContainersResourceList(pod.Spec.Containers)
		cpuUsage += resourceList.Cpu
		memoryUsage += resourceList.Memory
	}
	return &k8sDto.ResourceList{
		Cpu:    cpuUsage,
		Memory: memoryUsage,
	}, nil
}

// ResourcesOfAppByNamespace Count resource usage for a namespace
func ResourcesOfAppByNamespace(cli client.ResourceHandler, namespace, selector string) (map[string]*k8sDto.ResourceApp, error) {
	objs, err := cli.List(api.ResourceNamePod, namespace, selector)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*k8sDto.ResourceApp)
	for _, obj := range objs {
		pod := obj.(*v1.Pod)
		if pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
			continue
		}
		resourceList := k8sDto.ContainersResourceList(pod.Spec.Containers)
		if result[pod.Labels[util.AppLabelKey]] == nil {
			result[pod.Labels[util.AppLabelKey]] = &k8sDto.ResourceApp{
				Cpu:    resourceList.Cpu / 1000,
				Memory: resourceList.Memory / (1024 * 1024 * 1024),
				PodNum: 1,
			}
		} else {
			result[pod.Labels[util.AppLabelKey]].Cpu += resourceList.Cpu / 1000
			result[pod.Labels[util.AppLabelKey]].Memory += resourceList.Memory / (1024 * 1024 * 1024)
			result[pod.Labels[util.AppLabelKey]].PodNum += 1
		}
	}
	return result, nil
}
