package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/clim-bot/todo-list-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
            c.Abort()
            return
        }

        _, err := utils.ValidateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
