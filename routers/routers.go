package routers

import (
	"lelForum/controller"
	"lelForum/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// Sign Up route
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ping successful"})
	})
	// Handle 404 for undefined routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "route not found"})
	})

	return r
}
