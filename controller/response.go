package controller

import "github.com/gin-gonic/gin"

// Can be alternated by gin.H
type ResponseData struct {
	Code    RespCode    `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code RespCode) {
	c.JSON(200, &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code RespCode, msg interface{}) {
	c.JSON(200, &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	respData := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	c.JSON(200, respData)
}
