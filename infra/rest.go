package infra

import (
	"abanku/controller"
	"abanku/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ()

func NewRest(
	services service.Services,
) {
	router := gin.Default()

	controller.SetRoute(router, services)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	server.ListenAndServe()
}
