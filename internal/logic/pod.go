package logic

import (
	"Blueberry/internal/model"
	"Blueberry/pkg/k8s"
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreatePod(pod *model.Pod) error {
	podK8s := pod.ConvertToK8s()

	ctx := context.TODO()
	createdPod, err := k8s.Client.CoreV1().Pods(podK8s.Namespace).Create(ctx, podK8s, v1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Println(createdPod.Name)
	return nil
}
