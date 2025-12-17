package routers

import (
	"lelForum/controller"
	"lelForum/logger"
	"lelForum/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// Sign Up route
	r.POST("/signup", controller.SignUpHandler)
	//Login route
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/api/v1")
	//Only valid for routes defined after this line
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		// Protected routes go here
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
	}
	/*r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ping successful"})
	})*/
	// Handle 404 for undefined routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "route not found"})
	})

	return r
}
