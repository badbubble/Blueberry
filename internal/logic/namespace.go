package logic

import (
	"Blueberry/pkg/k8s"
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespace() error {
	ctx := context.TODO()
	list, err := k8s.Client.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Println(list)
	return nil
}
