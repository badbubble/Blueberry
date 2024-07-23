package main

import (
	"Blueberry/internal/router"
	"Blueberry/pkg/k8s"
	"Blueberry/pkg/logger"
	"Blueberry/pkg/setting"
	"fmt"
)

func main() {
	// initialize settings
	if err := setting.Init("config/dev.yaml"); err != nil {
		fmt.Printf("Fail to load configuration: %v", err)
		return
	}
	// create k8s client
	if err := k8s.Init(); err != nil {
		fmt.Printf("Fail to create k8s client: %v", err)
		return
	}
	// initialize logger
	if err := logger.Init(setting.Conf); err != nil {
		fmt.Printf("Fail to initialize logger: %v", err)
		return
	}

	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}


