package controller

import (
	"Blueberry/internal/logic"
	"Blueberry/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateSecretHandler(c *gin.Context) {
	secretPostData := &model.Secret{}
	if err := c.ShouldBind(secretPostData); err != nil {
		ResponseError(c, CodeInvalidParameter)
		return
	}
	if err := logic.CreateSecret(secretPostData); err != nil {
		ResponseError(c, CodeCreateSecretError)
		return
	}
	ResponseSuccess(c, nil)
}

func UpdateSecretHandler(c *gin.Context) {
	secretPostData := &model.Secret{}
	if err := c.ShouldBind(secretPostData); err != nil {
		ResponseError(c, CodeInvalidParameter)
		return
	}
	if err := logic.UpdateSecret(secretPostData); err != nil {
		ResponseError(c, CodeCreateSecretError)
		return
	}
	ResponseSuccess(c, nil)
}

func DeleteSecretHandler(c *gin.Context) {
	if err := logic.DeleteSecret(c.Query("namespace"), c.Query("name")); err != nil {
		ResponseError(c, CodeDeleteSecretError)
		return
	}

	ResponseSuccess(c, nil)
}

func GetSecretHandler(c *gin.Context) {
	name := c.Query("name")
	namespace := c.Query("namespace")
	if name != "" {
		secretDetail, err := logic.GetSecretDetail(namespace, name)
		if err != nil {
			ResponseError(c, CodeGetSecretDetailError)
		}
		ResponseSuccess(c, secretDetail)
		return
	}
	k8sList, err := logic.GetSecretList(namespace)
	if err != nil {
		ResponseError(c, CodeGetSecretListError)
	}
	ResponseSuccess(c, k8sList)
	return
}
