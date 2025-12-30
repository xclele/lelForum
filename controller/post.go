package controller

import (
	"lelForum/logic"
	"lelForum/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func GetPostListHandler(c *gin.Context) {
	//Pagination
	pageNumStr := c.DefaultQuery("page", "0")
	pageSizeStr := c.DefaultQuery("size", "10")
	var (
		page     int64
		pageSize int64
		err      error
	)
	if page, err = strconv.ParseInt(pageNumStr, 10, 64); err != nil {
		page = 1
	}
	if pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64); err != nil {
		pageSize = 10
	}
	//Retrieve data
	data, err := logic.GetPostList(page, pageSize)
	//Return response
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// Voting Part
func PostVoteHandler(c *gin.Context) {
	// Param Validation
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err = logic.VoteForPost(userID, p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
