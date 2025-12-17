package controller

import (
	"lelForum/logic"
	"lelForum/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	//Bind to the struct first (fast fail for invalid params)
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("CreatePost with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//Get the user ID from the context
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//Save to the database
	if err = logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//Return response
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	postID, err := strconv.ParseUint(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid post id", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostDetail(postID)
	if err != nil {
		zap.L().Error("logic.GetPostDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
