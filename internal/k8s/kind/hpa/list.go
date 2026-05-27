package hpa

import (
	autoscaling "k8s.io/api/autoscaling/v1"

	k8sDto "wayne/internal/k8s/dto"
)

type HPA struct {
	k8sDto.ObjectMeta `json:"objectMeta"`
	k8sDto.TypeMeta   `json:"typeMeta"`
	//ScaleTargetRef                  ScaleTargetRef `json:"scaleTargetRef"`
	MinReplicas                     *int32 `json:"minReplicas"`
	MaxReplicas                     int32  `json:"maxReplicas"`
	CurrentCPUUtilizationPercentage *int32 `json:"currentCPUUtilizationPercentage"`
	TargetCPUUtilizationPercentage  *int32 `json:"targetCPUUtilizationPercentage"`
}

func toHPA(hpa *autoscaling.HorizontalPodAutoscaler) *HPA {
	modelHPA := HPA{
		ObjectMeta: k8sDto.NewObjectMeta(hpa.ObjectMeta),
		TypeMeta:   k8sDto.NewTypeMeta("HorizontalPodAutoscaler"),

		MinReplicas:                     hpa.Spec.MinReplicas,
		MaxReplicas:                     hpa.Spec.MaxReplicas,
		CurrentCPUUtilizationPercentage: hpa.Status.CurrentCPUUtilizationPercentage,
		TargetCPUUtilizationPercentage:  hpa.Spec.TargetCPUUtilizationPercentage,
	}
	return &modelHPA
}
