package secret

import (
	"context"

	kapi "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	k8sDto "wayne/internal/k8s/dto"
)

type SecretType string

type Secret struct {
	ObjectMeta k8sDto.ObjectMeta `json:"objectMeta"`
	Data       map[string][]byte `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`
	StringData map[string]string `json:"stringData,omitempty" protobuf:"bytes,4,rep,name=stringData"`
	Type       SecretType        `json:"type,omitempty" protobuf:"bytes,3,opt,name=type,casttype=SecretType"`
}

func CreateOrUpdateSecret(cli *kubernetes.Clientset, secret *kapi.Secret) (*kapi.Secret, error) {

	ctx := context.Background()

	old, err := cli.CoreV1().Secrets(secret.Namespace).Get(ctx, secret.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.CoreV1().Secrets(secret.Namespace).Create(ctx, secret, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = secret.Labels
	old.Data = secret.Data

	return cli.CoreV1().Secrets(secret.Namespace).Update(ctx, old, metaV1.UpdateOptions{})
}

func DeleteSecret(cli *kubernetes.Clientset, name, namespace string) error {
	return cli.CoreV1().Secrets(namespace).Delete(context.Background(), name, metaV1.DeleteOptions{})
}
