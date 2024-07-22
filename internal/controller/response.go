package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    RespCode    `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code RespCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code RespCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}
