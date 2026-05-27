package deployment

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"

	"wayne/internal/k8s/client"
	api "wayne/internal/k8s/client"
	k8sDto "wayne/internal/k8s/dto"
	"wayne/internal/k8s/kind/event"
	"wayne/internal/k8s/kind/pod"
	"wayne/pkg/dto"
	"wayne/pkg/maps"
)

type Deployment struct {
	ObjectMeta k8sDto.ObjectMeta `json:"objectMeta"`
	Pods       k8sDto.PodInfo    `json:"pods"`
	Containers []string          `json:"containers"`
}

func GetDeploymentList(indexer *client.CacheFactory, namespace string) ([]*appsV1.Deployment, error) {
	deployments, err := indexer.DeploymentLister().Deployments(namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	sort.Slice(deployments, func(i, j int) bool {
		return deployments[i].Name > deployments[j].Name
	})
	return deployments, nil
}

func GetDeploymentResource(cli client.ResourceHandler, deployment *appsV1.Deployment) (*k8sDto.ResourceList, error) {
	obj, err := cli.Get(api.ResourceNameDeployment, deployment.Namespace, deployment.Name)

	if err != nil {
		if errors.IsNotFound(err) {
			return k8sDto.DeploymentResourceList(deployment), nil
		}
		return nil, err
	}
	old := obj.(*appsV1.Deployment)
	oldResourceList := k8sDto.DeploymentResourceList(old)
	newResourceList := k8sDto.DeploymentResourceList(deployment)

	return &k8sDto.ResourceList{
		Cpu:    newResourceList.Cpu - oldResourceList.Cpu,
		Memory: newResourceList.Memory - oldResourceList.Memory,
	}, nil
}

func CreateOrUpdateDeployment(cli *kubernetes.Clientset, deployment *appsV1.Deployment) (*appsV1.Deployment, error) {
	ctx := context.Background()

	old, err := cli.AppsV1().Deployments(deployment.Namespace).Get(ctx, deployment.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.AppsV1().Deployments(deployment.Namespace).Create(ctx, deployment, metaV1.CreateOptions{})
		}
		return nil, err
	}
	err = checkDeploymentLabelSelector(deployment, old)
	if err != nil {
		return nil, err
	}

	old.Labels = deployment.Labels
	old.Annotations = deployment.Annotations
	old.Spec = deployment.Spec
	return cli.AppsV1().Deployments(deployment.Namespace).Update(ctx, old, metaV1.UpdateOptions{})
}

func UpdateDeployment(cli *kubernetes.Clientset, deployment *appsV1.Deployment) (*appsV1.Deployment, error) {
	ctx := context.Background()

	old, err := cli.AppsV1().Deployments(deployment.Namespace).Get(ctx, deployment.Name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	err = checkDeploymentLabelSelector(deployment, old)
	if err != nil {
		return nil, err
	}

	return cli.AppsV1().Deployments(deployment.Namespace).Update(ctx, deployment, metaV1.UpdateOptions{})
}

func checkDeploymentLabelSelector(new *appsV1.Deployment, old *appsV1.Deployment) error {
	for key, value := range new.Spec.Selector.MatchLabels {
		oldValue, ok := old.Spec.Selector.MatchLabels[key]
		if !ok || oldValue != value {
			return &dto.ErrorResult{
				Code: http.StatusBadRequest,
				Msg: fmt.Sprintf("New's Deployment MatchLabels(%s) not match old MatchLabels(%s), do not allow deploy to prevent the orphan ReplicaSet. ",
					maps.LabelsToString(new.Spec.Selector.MatchLabels), maps.LabelsToString(old.Spec.Selector.MatchLabels)),
			}
		}
	}

	return nil
}

func GetDeployment(cli *kubernetes.Clientset, name, namespace string) (*appsV1.Deployment, error) {
	return cli.AppsV1().Deployments(namespace).Get(context.Background(), name, metaV1.GetOptions{})
}

func GetDeploymentDetail(cli *kubernetes.Clientset, indexer *client.CacheFactory, name, namespace string) (*Deployment, error) {
	deployment, err := cli.AppsV1().Deployments(namespace).Get(context.Background(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return toDeployment(deployment, indexer)
}

func toDeployment(deployment *appsV1.Deployment, indexer *client.CacheFactory) (*Deployment, error) {
	result := &Deployment{
		ObjectMeta: k8sDto.NewObjectMeta(deployment.ObjectMeta),
	}

	podInfo := k8sDto.PodInfo{}
	podInfo.Current = deployment.Status.AvailableReplicas
	podInfo.Desired = *deployment.Spec.Replicas
	var err error

	pods, err := pod.ListKubePod(indexer, deployment.Namespace, deployment.Spec.Template.Labels)
	if err != nil {
		return nil, err
	}

	podInfo.Warnings, err = event.GetPodsWarningEvents(indexer, pods)
	if err != nil {
		return nil, err
	}

	result.Pods = podInfo

	containers := make([]string, 0)
	for _, container := range deployment.Spec.Template.Spec.Containers {
		containers = append(containers, container.Image)
	}

	result.Containers = containers

	return result, nil
}

func DeleteDeployment(cli *kubernetes.Clientset, name, namespace string) error {
	deletionPropagation := metaV1.DeletePropagationBackground
	return cli.AppsV1().
		Deployments(namespace).
		Delete(context.Background(), name, metaV1.DeleteOptions{PropagationPolicy: &deletionPropagation})
}

func UpdateScale(cli *kubernetes.Clientset, deploymentname string, namespace string, newreplica int32) error {

	ctx := context.Background()

	deployments := cli.AppsV1().Deployments(namespace)

	deployment, err := deployments.Get(ctx, deploymentname, metaV1.GetOptions{})
	if err != nil {
		return err
	}
	deployment.Spec.Replicas = &newreplica
	_, err = deployments.Update(ctx, deployment, metaV1.UpdateOptions{})
	return err
}
