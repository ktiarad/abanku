package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (x ServiceImplementation) handleGetAccounts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := x.GetAllAccounts()

		jsonRes, err := json.Marshal(response)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)

			return
		}

		ctx.JSON(http.StatusOK, jsonRes)
	}
}
