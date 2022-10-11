package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/handler"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/docs"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/transactions"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title       MELI Bootcamp API Lucas
// @version     1.0
// @description This API Handle MELI Products.
func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("Error reading .env file:\n%+v", err.Error())
	}
	// Initialization of services
	transactionsStore := store.NewStore(store.FileType, "./transactions.json")
	transactionsRepository := transactions.NewRepository(transactionsStore)
	transactionsService := transactions.NewService(transactionsRepository)
	transactionsHandler := handler.NewHandler(transactionsService)
	engine := gin.Default()
	// Swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
		log.Fatalf("Error running engine:\n%+v", err.Error())
	}
}
