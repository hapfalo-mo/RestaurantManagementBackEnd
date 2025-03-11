package middleware

import (
	"strings"

	service "RestuarantBackend/service"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Missing Token"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == " " {
		c.JSON(401, gin.H{"error": "Invalid Token format"})
		c.Abort()
		return
	}

	claims, err := service.ParseToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("userId", claims.UserID)
	c.Next()
}
