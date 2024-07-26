package main

import (
	"Blueberry/internal/router"
	"Blueberry/pkg/k8s"
	"Blueberry/pkg/logger"
	"Blueberry/pkg/setting"
	"fmt"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
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
