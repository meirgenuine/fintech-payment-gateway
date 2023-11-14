package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meirgenuine/fintech-payment-gateway/internal/handler"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Fintech Payment Gateway!",
		})
	})

	// Register endpoint for transaction initiation
	router.POST("/transactions/initiate", handler.InitiateTransaction)

	router.Run(":8080")
}
