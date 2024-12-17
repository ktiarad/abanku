package infra

import (
	"abanku/controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ()

func NewRest() {
	router := gin.Default()

	controller.SetRoute(router)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	server.ListenAndServe()
}
