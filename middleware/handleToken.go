package middleware

import (
	"net/http"
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
	if tokenString == "" {
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
	c.Set("role", claims.Role)
	c.Next()
}

// Function for Admin Authen and Authorize
func AuthenAdminMiddelWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Missing Token"})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Invalid Token Format"})
		c.Abort()
		return
	}
	// Parse Token to get Claim
	claims, err := service.ParseToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	if claims.Role != 1 {
		c.JSON(http.StatusForbidden, gin.H{"Error": "You don't have permission to access this site! "})
		c.Abort()
		return
	}
	c.Next()
}
