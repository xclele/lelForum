package controller

import (
	"errors"
	"lelForum/database/postgres"
	"lelForum/logic"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	//"lelForum/logic"
	"lelForum/models"
)

// SignUpHandler handles user sign-up requests
func SignUpHandler(c *gin.Context) {
	// Param Validation
	var p models.ParamSignUp
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("SignUpHandler with invalid param", zap.Error(err))
		// Check if it's a validation error
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	// Can Also use manual validation

	// Business Logic
	if err = logic.SignUp(&p); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		// Check if user already exists
		if errors.Is(err, postgres.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		// Other errors
		ResponseError(c, CodeServerBusy)
		return
	}
	// Response
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// Get the login parameters
	p := models.ParamLogin{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))
		// Check if it's a validation error
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	// Logic
	token, err := logic.Login(&p)
	if err != nil {
		zap.L().Error("Login failed", zap.Error(err))
		if errors.Is(err, postgres.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// Response
	ResponseSuccess(c, token)
}
