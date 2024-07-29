package controller

import (
	"Blueberry/internal/logic"
	"Blueberry/internal/model"
	"github.com/gin-gonic/gin"
)

func GetConfigMapHandler(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	if namespace == "" {
		ResponseError(c, CodeInvalidParameter)
		return
	}
	if name == "" {
		configMapList, err := logic.GetConfigMapList(namespace)
		if err != nil {
			ResponseError(c, CodeGetConfigMapListError)
			return
		}
		ResponseSuccess(c, configMapList)
		return
	}
	configMapDetail, err := logic.GetConfigMapDetail(namespace, name)
	if err != nil {
		ResponseError(c, CodeGetConfigMapDetailError)
		return
	}
	ResponseSuccess(c, configMapDetail)
	return
}

func CreateConfigMapHandler(c *gin.Context) {
	postData := &model.ConfigMap{}
	err := c.ShouldBind(postData)
	if err != nil {
		ResponseError(c, CodeInvalidParameter)
		return
	}

	err = logic.CreateConfigMap(postData)
	if err != nil {
		ResponseError(c, CodeCreateConfigMapError)
		return
	}

	ResponseSuccess(c, nil)
}

func UpdateConfigMapHandler(c *gin.Context) {
	postData := &model.ConfigMap{}
	err := c.ShouldBind(postData)
	if err != nil {
		ResponseError(c, CodeInvalidParameter)
		return
	}

	err = logic.UpdateConfigMap(postData)
	if err != nil {
		ResponseError(c, CodeUpdateConfigMapError)
		return
	}

	ResponseSuccess(c, nil)
}

func DeleteConfigMapHandler(c *gin.Context) {
	name := c.Query("name")
	namespace := c.Query("namespace")
	err := logic.DeleteConfigMap(namespace, name)
	if err != nil {
		ResponseError(c, DeleteConfigMapError)
	}
	ResponseSuccess(c, nil)
}
