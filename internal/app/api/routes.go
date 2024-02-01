package api

import (
	"net/http"

	"github.com/Amilcarpio/go-backend-banco/internal/database"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rota para transferência de dinheiro
	router.POST("/transfer", TransferMoneyHandler)
}

func TransferMoneyHandler(c *gin.Context) {
	// Obter dados da requisição
	var transferRequest TransferRequest

	if err := c.ShouldBindJSON(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Realizar transferência de dinheiro
	err := database.TransferMoney(c.Request.Context(), transferRequest.Payer, transferRequest.Payee, transferRequest.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder com sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Transferência realizada com sucesso"})
}
