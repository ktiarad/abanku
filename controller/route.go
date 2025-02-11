package controller

import (
	"abanku/service"

	"github.com/gin-gonic/gin"
)

type ServiceImplementation struct {
	service.Services
}

func SetRoute(
	r *gin.Engine,
	s service.Services,
) {
	handler := ServiceImplementation{Services: s}

	r.GET("/", checkHealth())
	r.GET("/accounts", handler.handleGetAccounts())
}

func checkHealth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	}
}
