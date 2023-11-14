package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meirgenuine/fintech-payment-gateway/internal/model"
	"github.com/meirgenuine/fintech-payment-gateway/internal/service"
)

// InitiateTransaction handles the request to initiate a transaction.
func InitiateTransaction(c *gin.Context) {
	var request model.TransactionInitiationRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionID, err := service.InitiateTransaction(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.TransactionInitiationResponse{TransactionID: transactionID})
}
