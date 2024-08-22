package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/krishna/go-tasker/internal/utils"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
		c.Abort()
		return
	}

	tokenString = tokenString[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
		c.Abort()
		return
	}

	c.Set("user_id", uint(claims["user_id"].(float64)))
	c.Next()
}
