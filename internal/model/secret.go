package model

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Secret struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Type      corev1.SecretType `json:"type"`
	DataNum   int               `json:"dataNum"`
	Age       int64             `json:"age"`
	Labels    []ListMapItem     `json:"labels"`
	Data      []ListMapItem     `json:"data"`
}

func (s *Secret) ConvertToK8s() *corev1.Secret {
	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      s.Name,
			Namespace: s.Namespace,
			Labels:    ToMap(s.Labels),
		},
		StringData: ToMap(s.Data),
		Type:       s.Type,
	}
}

func (s *Secret) ConvertToJSONList(k8sSecret *corev1.Secret) {
	s.Name = k8sSecret.Name
	s.Namespace = k8sSecret.Namespace
	s.Age = k8sSecret.CreationTimestamp.Unix()
	s.DataNum = len(k8sSecret.Data)
	s.Type = k8sSecret.Type

}

func (s *Secret) ConvertToJSONDetail(k8sSecret *corev1.Secret) {
	s.ConvertToJSONList(k8sSecret)
	s.Labels = ToList(k8sSecret.Labels)
	s.Data = ToListWithByte(k8sSecret.Data)
}
