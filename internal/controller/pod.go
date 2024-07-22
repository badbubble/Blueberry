package controller

import (
	"Blueberry/pkg/k8s"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodListHandler(c *gin.Context) {
	ctx := context.TODO()
	list, err := k8s.Client.CoreV1().Pods("").List(ctx, v1.ListOptions{})
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "require list pod error")
		return
	}
	for _, item := range list.Items {
		fmt.Printf("%s, %s\n", item.Namespace, item.Name)
	}
	ResponseSuccess(c, nil)
}
