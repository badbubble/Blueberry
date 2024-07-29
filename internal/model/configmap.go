package model

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMap struct {
	Name      string        `json:"name"`
	Namespace string        `json:"namespace"`
	DataNum   int           `json:"dataNum"`
	Age       int64         `json:"age"`
	Labels    []ListMapItem `json:"labels"`
	Data      []ListMapItem `json:"data"`
}

func (c *ConfigMap) ConvertToK8s() *corev1.ConfigMap {
	k8sConfigMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.Name,
			Namespace: c.Namespace,
			Labels:    ToMap(c.Labels),
		},
		Data: ToMap(c.Data),
	}

	return k8sConfigMap
}

func (c *ConfigMap) ConvertToJSONList(k8sConfigmapData *corev1.ConfigMap) {
	c.Name = k8sConfigmapData.Name
	c.Namespace = k8sConfigmapData.Namespace
	c.DataNum = len(k8sConfigmapData.Data)
	c.Age = k8sConfigmapData.CreationTimestamp.Unix()
}

func (c *ConfigMap) ConvertToJSONDetail(k8sConfigmapData *corev1.ConfigMap) {
	c.ConvertToJSONList(k8sConfigmapData)
	c.Data = ToList(k8sConfigmapData.Data)
	c.Labels = ToList(k8sConfigmapData.Labels)
}
