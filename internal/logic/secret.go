package logic

import (
	"Blueberry/internal/model"
	"Blueberry/pkg/k8s"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateSecret(secretData *model.Secret) error {
	secretK8s := secretData.ConvertToK8s()
	_, err := k8s.Client.CoreV1().Secrets(secretData.Namespace).Create(context.TODO(), secretK8s, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return err
}

func UpdateSecret(secretData *model.Secret) error {
	secretK8s := secretData.ConvertToK8s()
	if _, err := k8s.Client.CoreV1().
		Secrets(secretData.Namespace).
		Get(context.TODO(), secretData.Name, metav1.GetOptions{}); err != nil {
		return err
	}

	if _, err := k8s.Client.CoreV1().
		Secrets(secretData.Namespace).
		Update(context.TODO(), secretK8s, metav1.UpdateOptions{}); err != nil {
		return err
	}

	return nil
}

func DeleteSecret(namespace string, name string) error {
	if err := k8s.Client.CoreV1().
		Secrets(namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func GetSecretDetail(namespace string, name string) (*model.Secret, error) {
	k8sSecret, err := k8s.Client.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	result := &model.Secret{}
	result.ConvertToJSONDetail(k8sSecret)
	return result, nil
}

func GetSecretList(namespace string) ([]*model.Secret, error) {
	k8sSecretList, err := k8s.Client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	result := make([]*model.Secret, 0)
	for _, item := range k8sSecretList.Items {
		singleSecret := &model.Secret{}
		singleSecret.ConvertToJSONList(&item)
		result = append(result, singleSecret)
	}
	return result, nil
}
