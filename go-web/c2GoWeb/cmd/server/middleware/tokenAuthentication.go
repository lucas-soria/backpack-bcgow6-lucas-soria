package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/web"
	"net/http"
	"os"
)

func TokenValidationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.Request.Header.Get("token"); token != os.Getenv("SUPER_SECRET_TOKEN") {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.NewResponse(
					401,
					nil,
					"you do not have permission to do this request",
				),
			)
			return
		}
		ctx.Next()
	}
}
