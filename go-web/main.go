package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"os"
	"reflect"
)

const (
	SuperSecretToken = "1357924680a"
)

var (
	transactions []transaction
)

type transaction struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transactionCode" binding:"required"`
	Currency        string  `json:"currency" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	Sender          string  `json:"sender" binding:"required"`
	Receiver        string  `json:"receiver" binding:"required"`
	Date            string  `json:"date" binding:"required"` // con db cambiar a time.Time
}

func helloMessageHandler(ctx *gin.Context) {
	name := ctx.Request.URL.Query().Get("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hola %s", name),
		"warning": "Always query your name with the 'name' variable",
	})
}

func GetAll(ctx *gin.Context) {
	var filteredTransactions []transaction
	queries := ctx.Request.URL.Query()
	var c bool
	for _, t := range transactions {
		reflection := reflect.ValueOf(t)
		c = true
		for i := 0; i < reflection.NumField(); i++ {
			if name := reflection.Type().Field(i).Tag.Get("json"); queries.Has(name) {
				if fmt.Sprintf("%v", reflection.Field(i)) != queries.Get(name) {
					c = false
					break
				}
			}
		}
		if c {
			filteredTransactions = append(filteredTransactions, t)
		}
	}
	if len(filteredTransactions) != 0 {
		ctx.JSON(http.StatusOK, filteredTransactions)
	} else {
		ctx.JSON(http.StatusNotFound, &[]string{})
	}
}

func GetOne(ctx *gin.Context) {
	for _, t := range transactions {
		reflection := reflect.ValueOf(t)
		if id := reflection.FieldByName("Id"); fmt.Sprintf("%v", id) == ctx.Param("id") {
			ctx.JSON(http.StatusOK, t)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, []transaction{})
}

func validateToken(ctx *gin.Context) (err error) {
	if token := ctx.Request.Header.Get("token"); token != SuperSecretToken {
		return fmt.Errorf("invalid token: %s", token)
	}
	return
}

func validateFields(err error) (errs []string) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errs = make([]string, len(ve))
		for i, fe := range ve {
			errs[i] = fmt.Sprintf("field '%s' is required", fe.Field())
		}
	}
	return
}

func Create(ctx *gin.Context) {
	if err := validateToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you do not have permission to do this request",
		})
		return
	}
	var t transaction
	if err := ctx.ShouldBindJSON(&t); err != nil {
		errs := validateFields(err)
		if len(errs) == 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errs[0]})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		}
		return
	}
	t.Id = getLastId() + 1
	transactions = append(transactions, t)
	ctx.JSON(http.StatusOK, t)
}

func getLastId() (lastId int) {
	if len(transactions) != 0 {
		lastId = transactions[0].Id
	}
	for _, t := range transactions {
		if t.Id > lastId {
			lastId = t.Id
		}
	}
	return
}

func readFile() {
	if textBytes, err := os.ReadFile("./transactions.json"); err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	} else if err1 := json.Unmarshal(textBytes, &transactions); err1 != nil {
		panic(fmt.Sprintf("error: %v", err1.Error()))
	}
	getLastId()
}

func main() {
	readFile()
	engine := gin.Default()

	// Paths
	// Welcome message
	engine.GET("/message", helloMessageHandler) // Example of usage http://localhost:8080/message?name=Lucas%20Soria
	// Transaction
	transactionsPrefix := engine.Group("/transactions")
	{
		// POST
		transactionsPrefix.POST("/", Create)
		// GETs
		transactionsPrefix.GET("/", GetAll)
		transactionsPrefix.GET("/:id", GetOne)
	}

	if err := engine.Run(); err != nil {
		log.Fatalf("Error: %+v", err)
	}
}
