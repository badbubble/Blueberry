package main

import (
	"Blueberry/internal/router"
	"Blueberry/pkg/k8s"
	"Blueberry/pkg/setting"
	"fmt"
)

func main() {
	if err := setting.Init("config/dev.yaml"); err != nil {
		return
	}
	k8s.Init()

	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
