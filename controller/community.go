package controller

import (
	"lelForum/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommunityHandler handles requests to the /community endpoint
func CommunityHandler(c *gin.Context) {
	//query the list of communities from the database
	data, err := logic.GetCommunity()
	if err != nil {
		zap.L().Error("logic.GetCommunity() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	//query the list of communities from the database
	idStr := c.Param("id")
	// convert idStr to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunity() failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, data)
}
