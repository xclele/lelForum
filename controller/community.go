package controller

import (
	"lelForum/logic"

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
