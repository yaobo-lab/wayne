package job

import (
	"context"
	"fmt"
	"sort"

	batchv1 "k8s.io/api/batch/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	"wayne/internal/k8s/client"
	api "wayne/internal/k8s/client"
	k8sDto "wayne/internal/k8s/dto"
	"wayne/internal/k8s/kind/dataselector"
	"wayne/internal/k8s/kind/event"
	"wayne/internal/k8s/kind/pod"
	"wayne/pkg/dto"
)

func GetRelatedJobByCronJob(kubeClient client.ResourceHandler, namespace, cronJob string, q *dto.QueryParam) (*dto.Page, error) {
	var objs = make([]runtime.Object, 0)
	var err error
	objs, err = kubeClient.List(api.ResourceNameJob, namespace, labels.Everything().String())
	if err != nil {
		return nil, err
	}

	relateJob := make([]*batchv1.Job, 0)
	for _, obj := range objs {
		job, ok := obj.(*batchv1.Job)
		if !ok {
			return nil, fmt.Errorf("Convert pod obj (%v) error. ", obj)
		}
		for _, ref := range job.OwnerReferences {
			if ref.Kind == api.KindNameCronJob && cronJob == ref.Name {
				relateJob = append(relateJob, job)
			}
		}

	}
	return pageResult(relateJob, q), nil
}

func pageResult(relateJob []*batchv1.Job, q *dto.QueryParam) *dto.Page {
	commonObjs := make([]dataselector.DataCell, 0)
	for _, job := range relateJob {
		commonObjs = append(commonObjs, ObjectCell(*job))
	}

	sort.Slice(commonObjs, func(i, j int) bool {
		return commonObjs[j].GetProperty(dataselector.CreationTimestampProperty).
			Compare(commonObjs[i].GetProperty(dataselector.CreationTimestampProperty)) == -1
	})

	return dataselector.DataSelectPage(commonObjs, q)
}

func GetJobsByCronjobName(cli *kubernetes.Clientset, namespace, cronjobName string) ([]batchv1.Job, error) {
	ctx := context.Background()

	cronjob, err := cli.BatchV2alpha1().CronJobs(namespace).Get(ctx, cronjobName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	jobSelector := labels.SelectorFromSet(cronjob.ObjectMeta.Labels).String()
	jobList, err := cli.BatchV1().Jobs(namespace).List(ctx, metaV1.ListOptions{LabelSelector: jobSelector})
	if err != nil {
		return nil, err
	}
	return jobList.Items, nil
}

func GetPodsEvent(cli *kubernetes.Clientset, indexer *client.CacheFactory, namespace, jobName, cronjobName string) (k8sDto.PodInfo, error) {
	podInfo := k8sDto.PodInfo{}
	cronjob, err := cli.BatchV2alpha1().CronJobs(namespace).Get(context.Background(), cronjobName, metaV1.GetOptions{})
	if err != nil {
		return podInfo, err
	}

	labelDetail := cronjob.ObjectMeta.Labels
	labelDetail["job-name"] = jobName

	pods, err := pod.ListKubePod(indexer, namespace, labelDetail)
	if err != nil {
		return podInfo, err
	}

	podInfo.Current = int32(len(pods))
	// 目前写死为1
	podInfo.Desired = 1

	podInfo.Warnings, err = event.GetPodsWarningEvents(indexer, pods)
	if err != nil {
		return k8sDto.PodInfo{}, err
	}

	return podInfo, nil
}
