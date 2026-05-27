package crd

import (
	"context"
	"encoding/json"
	"fmt"

	"wayne/pkg/logger"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	"wayne/internal/k8s/kind/dataselector"
	"wayne/pkg/dto"
)

func GetCustomCRD(cli *kubernetes.Clientset, group, version, kind, namespace, name string) (runtime.Object, error) {
	req := cli.RESTClient().Verb("GET").RequestURI(
		fmt.Sprintf("/apis/%s/%s/namespaces/%s/%s/%s",
			group,
			version,
			namespace,
			kind,
			name))
	raw, err := req.Do(context.Background()).Raw()
	if err != nil {
		return nil, err
	}
	result := &runtime.Unknown{}
	err = json.Unmarshal(raw, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreatCustomCRD(cli *kubernetes.Clientset, group, version, kind, namespace string, body interface{}) (runtime.Object, error) {
	req := cli.RESTClient().Verb("POST").RequestURI(
		fmt.Sprintf("/apis/%s/%s/namespaces/%s/%s",
			group,
			version,
			namespace,
			kind)).Body(body)
	raw, err := req.Do(context.Background()).Raw()
	if err != nil {
		return nil, err
	}
	result := &runtime.Unknown{}
	err = json.Unmarshal(raw, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateCustomCRD(cli *kubernetes.Clientset, group, version, kind, namespace, name string, object *runtime.Unknown) (runtime.Object, error) {
	req := cli.RESTClient().Verb("PUT").RequestURI(
		fmt.Sprintf("/apis/%s/%s/namespaces/%s/%s/%s",
			group,
			version,
			namespace,
			kind,
			name)).
		Body([]byte(object.Raw)).
		SetHeader("Content-Type", "application/json")
	raw, err := req.Do(context.Background()).Raw()
	if err != nil {
		logger.Errorf(req.URL().String(), err)
		return nil, err
	}
	result := &runtime.Unknown{}
	err = json.Unmarshal(raw, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteCustomCRD(cli *kubernetes.Clientset, group, version, kind, namespace, name string) error {
	req := cli.RESTClient().Verb("DELETE").RequestURI(
		fmt.Sprintf("/apis/%s/%s/namespaces/%s/%s/%s",
			group,
			version,
			namespace,
			kind,
			name))
	return req.Do(context.Background()).Error()
}

func GetCustomCRDPage(cli *kubernetes.Clientset, group, version, kind, namespace string, q *dto.QueryParam) (*dto.Page, error) {
	req := cli.RESTClient().Verb("GET").RequestURI(
		fmt.Sprintf("/apis/%s/%s/namespaces/%s/%s",
			group,
			version,
			namespace,
			kind))
	result, err := req.Do(context.Background()).Raw()
	if err != nil {
		return nil, err
	}

	crdList := &CustomCRDList{}
	err = json.Unmarshal(result, crdList)
	if err != nil {
		return nil, err
	}
	return dataselector.DataSelectPage(toCustomCRDCells(crdList.Items), q), nil
}

func toCustomCRDCells(deploy []CustomCRD) []dataselector.DataCell {
	cells := make([]dataselector.DataCell, len(deploy))
	for i := range deploy {
		cells[i] = CustomCRDCell(deploy[i])
	}
	return cells
}
