package configmap

import (
	"context"

	"wayne/internal/k8s/dto"

	kapi "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMap struct {
	ObjectMeta dto.ObjectMeta    `json:"objectMeta"`
	Data       map[string]string `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`
}

func CreateOrUpdateConfigMap(cli *kubernetes.Clientset, configMap *kapi.ConfigMap) (*kapi.ConfigMap, error) {

	ctx := context.Background()

	old, err := cli.CoreV1().ConfigMaps(configMap.Namespace).Get(ctx, configMap.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.CoreV1().ConfigMaps(configMap.Namespace).Create(ctx, configMap, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = configMap.Labels
	old.Data = configMap.Data

	return cli.CoreV1().ConfigMaps(configMap.Namespace).Update(ctx, old, metaV1.UpdateOptions{})
}
