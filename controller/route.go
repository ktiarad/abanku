package controller

import "github.com/gin-gonic/gin"

func SetRoute(
	r *gin.Engine,
) {
	r.GET("/", checkHealth())
	r.GET("/accounts")
}

func checkHealth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	}
}
