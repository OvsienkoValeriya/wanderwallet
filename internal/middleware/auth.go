package middleware

import (
	"net/http"
	"os"
	"wanderwallet/initializers"
	"wanderwallet/internal/models"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *gin.Context) {
	// Пытаемся получить токен из заголовка Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		// Если нет заголовка, пытаемся взять из cookie
		tokenStr, err := c.Cookie("Authorization")
		if err != nil || tokenStr == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		authHeader = "Bearer " + tokenStr
	}

	// Ожидается формат: "Bearer <token>"
	const prefix = "Bearer "
	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenStr := authHeader[len(prefix):]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := initializers.DB.First(&user, sub).Error; err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
