package middleware

import (
	"go-authapi-adv/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := utils.TokenValid(ctx); err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
