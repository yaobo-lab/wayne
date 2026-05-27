package cronjob

import (
	"context"

	batchv1beta1 "k8s.io/api/batch/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"wayne/internal/k8s/dto"
	"wayne/pkg/maps"
)

type Cronjob struct {
	ObjectMeta dto.ObjectMeta `json:"objectMeta"`
	Spec       CronJobSpec    `json:"spec"`
}

type CronJobSpec struct {
	Suspend *bool `json:"suspend,omitempty" protobuf:"varint,4,opt,name=suspend"`
}

func CreateOrUpdateCronjob(cli *kubernetes.Clientset, cronjob *batchv1beta1.CronJob) (*batchv1beta1.CronJob, error) {
	ctx := context.Background()

	cronjobClient := cli.BatchV1beta1().CronJobs(cronjob.Namespace)
	old, err := cronjobClient.Get(ctx, cronjob.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cronjobClient.Create(ctx, cronjob, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = maps.MergeLabels(old.Labels, cronjob.Labels)
	oldTemplateLabels := old.Spec.JobTemplate.Spec.Template.Labels
	old.Spec = cronjob.Spec
	old.Spec.JobTemplate.Spec.Template.Labels = maps.MergeLabels(oldTemplateLabels, cronjob.Spec.JobTemplate.Spec.Template.Labels)
	return cronjobClient.Update(ctx, old, metaV1.UpdateOptions{})
}

func GetCronjobDetail(cli *kubernetes.Clientset, name, namespace string) (*Cronjob, error) {
	cronjob, err := cli.BatchV1beta1().CronJobs(namespace).Get(context.Background(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	result := &Cronjob{
		ObjectMeta: dto.NewObjectMeta(cronjob.ObjectMeta),
		Spec: CronJobSpec{
			Suspend: cronjob.Spec.Suspend,
		},
	}
	return result, nil
}

func SuspendCronjob(cli *kubernetes.Clientset, name, namespace string) error {

	ctx := context.Background()

	cronjob, err := cli.BatchV1beta1().CronJobs(namespace).Get(ctx, name, metaV1.GetOptions{})
	if err != nil {
		return err
	}

	suspend := true

	cronjob.Spec.Suspend = &suspend
	_, err = cli.BatchV1beta1().CronJobs(namespace).Update(ctx, cronjob, metaV1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeleteCronjob(cli *kubernetes.Clientset, name, namespace string) error {
	deletionPropagation := metaV1.DeletePropagationBackground
	return cli.BatchV1beta1().
		CronJobs(namespace).
		Delete(context.Background(), name, metaV1.DeleteOptions{PropagationPolicy: &deletionPropagation})
}
