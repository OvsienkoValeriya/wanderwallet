package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DebugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Логируем все куки
		cookies := c.Request.Cookies()
		for _, cookie := range cookies {
			fmt.Printf("[DEBUG] Cookie: %s = %s\n", cookie.Name, cookie.Value)
		}

		// Логируем хедеры авторизации
		if auth := c.GetHeader("Authorization"); auth != "" {
			fmt.Printf("[DEBUG] Authorization header: %s\n", auth)
		}

		// Если в контексте уже есть user — печатаем
		if user, exists := c.Get("user"); exists {
			fmt.Printf("[DEBUG] User in context: %#v\n", user)
		} else {
			fmt.Println("[DEBUG] No user in context")
		}

		c.Next()
	}
}
