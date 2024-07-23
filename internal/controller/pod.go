package controller

import (
	"Blueberry/internal/logic"
	"Blueberry/internal/model"
	"Blueberry/internal/validate"
	"github.com/gin-gonic/gin"
)

func CreatePodHandler(c *gin.Context) {
	pod := &model.Pod{}

	if err := c.ShouldBind(pod); err != nil {
		ResponseErrorWithMsg(c, CodeParameterError, err.Error())
		return
	}
	if err := validate.PodCreate(pod); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParameter, err.Error())
		return
	}
	if err := logic.CreatePod(pod); err != nil {
		ResponseErrorWithMsg(c, CodeCreatePodError, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

func UpdatePodHandler(c *gin.Context) {
	pod := &model.Pod{}

	if err := c.ShouldBind(pod); err != nil {
		ResponseErrorWithMsg(c, CodeParameterError, err.Error())
		return
	}
	if err := validate.PodCreate(pod); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParameter, err.Error())
		return
	}
	if err := logic.UpdatePod(pod); err != nil {
		ResponseErrorWithMsg(c, CodeCreatePodError, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

func DeletePodHandler(c *gin.Context) {
	pod := &model.Pod{}

	if err := c.ShouldBind(pod); err != nil {
		ResponseErrorWithMsg(c, CodeParameterError, err.Error())
		return
	}
	if err := validate.PodCreate(pod); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParameter, err.Error())
		return
	}
	if err := logic.DeletePod(pod); err != nil {
		ResponseErrorWithMsg(c, CodeCreatePodError, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

func GetPodHandler(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")

	// get a pod's detail in a namespace
	if name != "" {
		pod, err := logic.GetPodDetail(namespace, name)
		if err != nil {
			ResponseError(c, CodeGetPodDetailError)
			return
		}
		ResponseSuccess(c, pod)
		return
	}
	// get pod list in a namespace
	podList, err := logic.GetPodList(namespace)
	if err != nil {
		ResponseError(c, CodeGetPodListError)
		return
	}
	ResponseSuccess(c, podList)
	return
}
