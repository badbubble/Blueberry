package logic

import (
	"Blueberry/internal/model"
	"Blueberry/pkg/k8s"
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespaceList() (error, []*model.Namespace) {
	ctx := context.TODO()
	list, err := k8s.Client.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return err, nil
	}
	namespaceList := make([]*model.Namespace, 0)
	for _, item := range list.Items {
		namespaceList = append(namespaceList, &model.Namespace{
			Name:              item.Name,
			CreationTimestamp: item.CreationTimestamp.Unix(),
			Status:            string(item.Status.Phase),
		})
	}
	return nil, namespaceList
}
