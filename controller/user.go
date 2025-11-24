package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	//"lelForum/logic"
	"lelForum/models"
	"net/http"
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
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": errs.Translate(trans),
		})
		return
	}
	// Can Also use manual validation

	// Business Logic
	//logic.SignUp()
	// Response
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
