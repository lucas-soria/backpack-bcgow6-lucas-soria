package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/handler"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/transactions"
	"log"
)

func main() {
	transactionsRepository := transactions.NewRepository()
	transactionsService := transactions.NewService(transactionsRepository)
	transactionsHandler := handler.NewHandler(transactionsService)
	engine := gin.Default()
	// Endpoints
	// Welcome message
	engine.GET("/message", transactionsHandler.HelloMessageHandler) // Example of usage http://localhost:8080/message?name=Lucas%20Soria
	// Transactions
	transactionsPrefix := engine.Group("/transactions")
	{ // -> This does nothing, but helps with clean code
		// GETs
		transactionsPrefix.GET("/", transactionsHandler.GetAll)
		transactionsPrefix.GET("/:id", transactionsHandler.GetOne)
		// POST
		transactionsPrefix.POST("/", transactionsHandler.Create)
		// PUTs
		transactionsPrefix.PUT("/:id", transactionsHandler.Update)
		// PATCH
		transactionsPrefix.PATCH("/:id", transactionsHandler.PartialUpdate)
		// DELETE
		transactionsPrefix.DELETE("/:id", transactionsHandler.Delete)
	}

	if err := engine.Run(); err != nil {
		log.Fatalf("Error: %+v", err)
	}
}
