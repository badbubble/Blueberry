package logic

import (
	"Blueberry/internal/model"
	"Blueberry/pkg/k8s"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetConfigMapDetail(namespace string, name string) (*model.ConfigMap, error) {
	result := &model.ConfigMap{}
	k8sConfigMap, err := k8s.Client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	result.ConvertToJSONDetail(k8sConfigMap)
	return result, nil
}

func GetConfigMapList(namespace string) ([]*model.ConfigMap, error) {
	configMapList := make([]*model.ConfigMap, 0)
	k8sList, err := k8s.Client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range k8sList.Items {
		configMapItem := &model.ConfigMap{}
		configMapItem.ConvertToJSONList(&item)
		configMapList = append(configMapList, configMapItem)
	}
	return configMapList, nil
}

func CreateConfigMap(configmapData *model.ConfigMap) error {
	_, err := k8s.Client.CoreV1().ConfigMaps(configmapData.Namespace).Create(
		context.TODO(), configmapData.ConvertToK8s(), metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func UpdateConfigMap(configmapData *model.ConfigMap) error {
	_, err := k8s.Client.CoreV1().ConfigMaps(configmapData.Namespace).Get(
		context.TODO(), configmapData.Name, metav1.GetOptions{},
	)
	if err != nil {
		return err
	}

	if _, err := k8s.Client.CoreV1().ConfigMaps(configmapData.Namespace).Update(
		context.TODO(), configmapData.ConvertToK8s(), metav1.UpdateOptions{},
	); err != nil {
		return err
	}

	return nil
}

func DeleteConfigMap(namespace string, name string) error {
	err := k8s.Client.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
