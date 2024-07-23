package logic

import (
	"Blueberry/internal/model"
	"Blueberry/pkg/k8s"
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"strings"
)

func CreatePod(pod *model.Pod) error {
	podK8s := pod.ConvertToK8s()

	ctx := context.TODO()
	_, err := k8s.Client.CoreV1().Pods(podK8s.Namespace).Create(ctx, podK8s, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeletePod(pod *model.Pod) error {
	ctx := context.TODO()
	err := k8s.Client.CoreV1().Pods(pod.Base.Namespace).Delete(ctx, pod.Base.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func UpdatePod(pod *model.Pod) error {
	if err := DeletePod(pod); err != nil {
		return err
	}
	ctx := context.TODO()
	labels := pod.GetLabels()
	var labelsList []string
	for k, v := range labels {
		labelsList = append(labelsList, fmt.Sprintf("%s=%s", k, v))
	}
	watcher, err := k8s.Client.CoreV1().Pods(pod.Base.Namespace).Watch(ctx, metav1.ListOptions{
		LabelSelector: strings.Join(labelsList, ","),
	})

	if err != nil {
		return err
	}
	for event := range watcher.ResultChan() {
		// Pod may be deleted before watching
		_, err := k8s.Client.CoreV1().Pods(pod.Base.Namespace).Get(ctx, pod.Base.Name, metav1.GetOptions{})
		if k8sErrors.IsNotFound(err) {
			break
		}
		// it must be the same pod
		k8sPod := event.Object.(*corev1.Pod)
		if k8sPod.Name != pod.Base.Name {
			continue
		}
		if event.Type == watch.Deleted {
			break
		}
	}

	if err := CreatePod(pod); err != nil {
		return err
	}
	return nil
}

func GetPodDetail(namespace string, name string) (*model.Pod, error) {
	pod := &model.Pod{}
	ctx := context.TODO()
	k8sPod, err := k8s.Client.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	pod.ConvertToPod(k8sPod)
	return pod, nil
}

func GetPodList(namespace string) ([]*model.PodItem, error) {
	podItemList := make([]*model.PodItem, 0)

	ctx := context.TODO()
	podList, err := k8s.Client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range podList.Items {
		podItem := &model.PodItem{}
		podItem.Convert(&item)
		podItemList = append(podItemList, podItem)
	}
	return podItemList, nil
}
