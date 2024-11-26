package middleware

import (
	"strings"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/utils"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: "Missing token"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateToken(token)

		if err != nil {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.ID)
		c.Next()
	}
}
