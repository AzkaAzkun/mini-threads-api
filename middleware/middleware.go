package middleware

import (
	"net/http"
	"strings"

	"github.com/AzkaAzkun/mini-threads-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to verify token",
				"error":   "Token not found",
			})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to verify token",
				"error":   "Invalid token format",
			})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		idToken, err := utils.GetPayloadInsideToken(token)
		if err != nil {
			errorMessage := "Failed to verify token"
			if err.Error() == "token expired" {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": errorMessage,
					"error":   "Token expired",
				})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": errorMessage,
				"error":   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("token", token)
		ctx.Set("payload", idToken)
		ctx.Set("user_id", idToken["user_id"])
		ctx.Set("email", idToken["email"])

		ctx.Next()
	}
}
