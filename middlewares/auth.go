package middlewares

import (
	"lelForum/controller"
	"lelForum/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Clients can carry the token via header, body, or URI (URI is insecure).
		// Here we assume the token is in the Authorization header with the **Bearer** scheme.
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// Split the header by space.
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "invalid authorization header format",
			})
			c.Abort()
			return
		}
		// parts[1] is the token string; parse it with the previously defined function.
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "invalid token",
			})
			c.Abort()
			return
		}
		// Save the username to the request context for downstream handlers.
		c.Set(controller.CtxUserIDKey, mc.UserId)
		c.Next() // Later handlers can use c.Get("username") to retrieve it.
	}
}
