package controller

import (
	"abanku/model"

	"github.com/gin-gonic/gin"
)

func handleGetAccounts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response []model.Account

		response = Account.GetAllAccount()

	}
}
