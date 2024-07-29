package router

import (
	_ "Blueberry/docs"
	"Blueberry/internal/controller"
	"Blueberry/internal/middleware"
	"Blueberry/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORSMiddleware(), logger.GinLogger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		k8sGroup := v1.Group("/k8s")
		{
			// namespace
			k8sGroup.GET("/namespace", controller.GetNamespace)
			// pod
			k8sGroup.GET("/pod", controller.GetPodHandler)
			k8sGroup.POST("/pod", controller.CreatePodHandler)
			k8sGroup.DELETE("/pod", controller.DeletePodHandler)
			k8sGroup.PUT("/pod", controller.UpdatePodHandler)

			// node
			k8sGroup.GET("/node", controller.GetNodeHandler)
			k8sGroup.PUT("/node/labels", controller.UpdateNodeLabelsHandler)
			k8sGroup.PUT("/node/taints", controller.UpdateNodeTaintsHandler)

			// configMap
			k8sGroup.GET("/configmap", controller.GetConfigMapHandler)
			k8sGroup.POST("/configmap", controller.CreateConfigMapHandler)
			k8sGroup.PUT("/configmap", controller.UpdateConfigMapHandler)
			k8sGroup.DELETE("/configmap", controller.DeleteConfigMapHandler)

			// secret
			k8sGroup.GET("/secret", controller.GetSecretHandler)
			k8sGroup.POST("/secret", controller.CreateSecretHandler)
			k8sGroup.PUT("/secret", controller.UpdateSecretHandler)
			k8sGroup.DELETE("/secret", controller.DeleteSecretHandler)

		}
	}
	return r
}
