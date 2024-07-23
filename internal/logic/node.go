package logic

import (
	"Blueberry/internal/model"
	"Blueberry/pkg/k8s"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeList() ([]*model.Node, error) {
	nodeList := make([]*model.Node, 0)
	ctx := context.TODO()
	k8sNodeList, err := k8s.Client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range k8sNodeList.Items {
		n := &model.Node{}
		n.FillWithK8sNode(&item)
		nodeList = append(nodeList, n)
	}
	return nodeList, nil
}
