package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/handler"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/cmd/server/middleware"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/transactions"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	BaseURL = "/transactions"
)

func createServer() (engine *gin.Engine) {
	if err := godotenv.Load("../.env.local"); err != nil {
		log.Fatalf("Error reading .env file:\n%+v", err.Error())
	}
	// Initialization of services
	transactionsStore := store.NewStore(store.FileType, "../transactions.json")
	transactionsRepository := transactions.NewRepository(transactionsStore)
	transactionsService := transactions.NewService(transactionsRepository)
	transactionsHandler := handler.NewHandler(transactionsService)
	engine = gin.New()
	// Endpoints
	// Welcome message
	engine.GET("/message", transactionsHandler.HelloMessageHandler) // Example of usage http://localhost:8080/message?name=Lucas%20Soria
	// Transactions
	transactionsPrefix := engine.Group("/transactions")
	{ // -> This does nothing, but helps with clean code
		// GETs
		transactionsPrefix.GET("", transactionsHandler.GetAll)
		transactionsPrefix.GET("/:id", transactionsHandler.GetOne)
		// POST
		transactionsPrefix.POST("", middleware.TokenValidationMiddleware(), transactionsHandler.Create)
		// PUT
		transactionsPrefix.PUT("/:id", middleware.TokenValidationMiddleware(), transactionsHandler.Update)
		// PATCH
		transactionsPrefix.PATCH("/:id", middleware.TokenValidationMiddleware(), transactionsHandler.PartialUpdate)
		// DELETE
		transactionsPrefix.DELETE("/:id", middleware.TokenValidationMiddleware(), transactionsHandler.Delete)
	}
	return
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", os.Getenv("SUPER_SECRET_TOKEN"))

	return request, httptest.NewRecorder()
}

func testGetTransactionsOk(t *testing.T) {
	server := createServer()
	request, responseRecorder := createRequestTest(http.MethodGet, BaseURL, "")
	server.ServeHTTP(responseRecorder, request)
	var responseObjects domain.Transaction
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseObjects)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func testSaveTransactionOk(t *testing.T) {
	server := createServer()
	request, responseRecorder := createRequestTest(
		http.MethodPost,
		BaseURL,
		`{
					"transactionCode": "209c09sd09s",
					"currency": "ARS",
					"amount": 11.3,
                    "sender": "lkm2lkm2lkm2",
					"receiver": "lkm2lkm23",
					"date": "2022-10-07T00:00:00-03:00"
               }`,
	)
	server.ServeHTTP(responseRecorder, request)
	var responseObject domain.Transaction
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseObject)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func testDeleteTransactionOk(t *testing.T) {
	server := createServer()
	deleteID := "/12"
	request, responseRecorder := createRequestTest(http.MethodDelete, BaseURL+deleteID, "")
	server.ServeHTTP(responseRecorder, request)
	assert.Equal(t, http.StatusNoContent, responseRecorder.Code)
}

func testDeleteTransactionNotFound(t *testing.T) {
	server := createServer()
	deleteID := "/123"
	request, responseRecorder := createRequestTest(http.MethodDelete, BaseURL+deleteID, "")
	server.ServeHTTP(responseRecorder, request)
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func testDeleteTransaction(t *testing.T) {
	t.Run("DELETE OK", testDeleteTransactionOk)
	t.Run("DELETE Not Found", testDeleteTransactionNotFound)
}

func Test_Suit(t *testing.T) {
	t.Run("GET", testGetTransactionsOk)
	t.Run("Test Suit Delete", testDeleteTransaction)
	t.Run("POST", testSaveTransactionOk)
}
