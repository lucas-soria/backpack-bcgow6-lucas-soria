package clase

import "github.com/gin-gonic/gin"

func GinGonic() {
	router := gin.Default()
	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "bienvenidos a go web",
		})
	})
	router.Run()
}
