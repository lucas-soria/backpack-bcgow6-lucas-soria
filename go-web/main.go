package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"reflect"
)

type transaction struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transactionCode"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Sender          string  `json:"sender"`
	Receiver        string  `json:"receiver"`
	Date            string  `json:"date"` // con db cambiar a time.Time
}

func helloMessageHandler(ctx *gin.Context) {
	name := ctx.Request.URL.Query().Get("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hola %s", name),
		"warning": "Always query your name with the 'name' variable",
	})
}

func GetAll(ctx *gin.Context) {
	if textBytes, err := os.ReadFile("./transactions.json"); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		var transactions []transaction
		if err1 := json.Unmarshal(textBytes, &transactions); err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			var filteredTransactions []transaction
			queries := ctx.Request.URL.Query()
			var c bool
			for _, t := range transactions {
				reflection := reflect.ValueOf(t)
				c = true
				for i := 0; i < reflection.NumField(); i++ {
					if name := reflection.Type().Field(i).Tag.Get("json"); queries.Has(name) {
						if fmt.Sprintf("%v", reflection.Field(i)) == queries.Get(name) {
							c = c && true
						} else {
							c = c && false
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
	}
}

func GetOne(ctx *gin.Context) {
	if textBytes, err := os.ReadFile("./transactions.json"); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		var transactions []transaction
		if err1 := json.Unmarshal(textBytes, &transactions); err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			for _, t := range transactions {
				reflection := reflect.ValueOf(t)
				if id := reflection.FieldByName("Id"); fmt.Sprintf("%v", id) == ctx.Param("id") {
					ctx.JSON(http.StatusOK, t)
				}
			}
			ctx.JSON(http.StatusNotFound, []transaction{})
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/message", helloMessageHandler) // Example of usage http://localhost:8080/message?name=Lucas%20Soria
	transactions := router.Group("/transactions")
	{
		transactions.GET("/", GetAll)
		transactions.GET("/:id", GetOne)
	}
	if err := router.Run(); err != nil {
		log.Fatalf("Error: %+v", err)
	}
}
