package dto

import (
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Secrets           []v1.ObjectReference              `json:"secrets,omitempty"`
	Provisioner       string                            `json:"provisioner,omitempty"`
	ReclaimPolicy     *v1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy,omitempty"`
	Subsets           interface{}                       `json:"subsets,omitempty"`
	Type              interface{}                       `json:"type,omitempty"`
	Data              interface{}                       `json:"data,omitempty"`
	Spec              interface{}                       `json:"spec,omitempty"`
	Status            interface{}                       `json:"status,omitempty"`
}

type BaseObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

func ToBaseObject(obj runtime.Object) (*BaseObject, error) {
	objByte, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var commonObj BaseObject
	err = json.Unmarshal(objByte, &commonObj)
	if err != nil {
		return nil, err
	}
	return &commonObj, nil
}
