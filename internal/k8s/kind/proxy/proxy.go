package proxy

import (
	"sort"

	"k8s.io/apimachinery/pkg/util/json"

	"wayne/internal/k8s/client"
	"wayne/internal/k8s/kind/dataselector"
	k8sDto "wayne/internal/model/dto"
	"wayne/pkg/dto"
)

func GetPage(kubeClient client.ResourceHandler, kind string, namespace string, q *dto.QueryParam) (*dto.Page, error) {
	objs, err := kubeClient.List(kind, namespace, q.LabelSelector)
	if err != nil {
		return nil, err
	}
	commonObjs := make([]dataselector.DataCell, 0)
	for _, obj := range objs {
		objCell, err := getRealObjCellByKind(kind, obj)
		if err != nil {
			return nil, err
		}
		commonObjs = append(commonObjs, objCell)
	}

	sort.Slice(commonObjs, func(i, j int) bool {
		return commonObjs[j].GetProperty(dataselector.NameProperty).Compare(commonObjs[i].GetProperty(dataselector.NameProperty)) == 1
	})

	return dataselector.DataSelectPage(commonObjs, q), nil
}

func GetNames(kubeClient client.ResourceHandler, kind string, namespace string) ([]k8sDto.NamesObject, error) {
	objs, err := kubeClient.List(kind, namespace, "")
	if err != nil {
		return nil, err
	}

	commonObjs := make([]k8sDto.NamesObject, 0)
	for _, obj := range objs {
		objByte, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var commonObj ObjectCell
		err = json.Unmarshal(objByte, &commonObj)
		if err != nil {
			return nil, err
		}
		commonObjs = append(commonObjs, k8sDto.NamesObject{
			Name: commonObj.Name,
		})
	}

	sort.Slice(commonObjs, func(i, j int) bool {
		return commonObjs[i].Name < commonObjs[j].Name
	})

	return commonObjs, nil
}
