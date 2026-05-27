package dto

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Event struct {
	ObjectMeta      ObjectMeta `json:"objectMeta"`
	TypeMeta        TypeMeta   `json:"typeMeta"`
	Message         string     `json:"message"`
	SourceComponent string     `json:"sourceComponent"`
	Name            string     `json:"name"`
	SubObject       string     `json:"object"`
	Count           int32      `json:"count"`
	FirstSeen       v1.Time    `json:"firstSeen"`
	LastSeen        v1.Time    `json:"lastSeen"`
	Reason          string     `json:"reason"`
	Type            string     `json:"type"`
}
