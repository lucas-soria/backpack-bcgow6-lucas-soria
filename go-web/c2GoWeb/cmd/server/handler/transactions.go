package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/transactions"
	"net/http"
	"strconv"
)

const (
	SuperSecretToken = "1357924680a"
)

type Controller struct {
	service transactions.Service
}

type Request struct {
	TransactionCode string  `json:"transactionCode" binding:"required"`
	Currency        string  `json:"currency" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	Sender          string  `json:"sender" binding:"required"`
	Receiver        string  `json:"receiver" binding:"required"`
	Date            string  `json:"date" binding:"required"` // TODO: con db cambiar a time.Time
}

func NewHandler(service transactions.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func NewTransactionFromRequest(request Request) (transaction transactions.Transaction) {
	transaction.TransactionCode = request.TransactionCode
	transaction.Currency = request.Currency
	transaction.Amount = request.Amount
	transaction.Sender = request.Sender
	transaction.Receiver = request.Receiver
	transaction.Date = request.Date
	return
}

func (c *Controller) HelloMessageHandler(ctx *gin.Context) {
	name := ctx.Request.URL.Query().Get("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hola %s", name),
		"warning": "Always query your name with the 'name' variable",
	})
	return
}

func (c *Controller) GetAll(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	ts, err := c.service.GetAll(queries)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error while retrieving transactions: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, ts)
	return
}

func (c *Controller) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error while converting id: %s", err),
		})
		return
	}
	t, err := c.service.GetOne(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error while retrieving transaction: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, t)
	return
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

func (c *Controller) Create(ctx *gin.Context) {
	if err := validateToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you do not have permission to do this request",
		})
		return
	}
	var r Request
	if err := ctx.ShouldBindJSON(&r); err != nil {
		errs := validateFields(err)
		if len(errs) == 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errs[0]})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		}
		return
	}
	t, err := c.service.Store(NewTransactionFromRequest(r))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error while storing transaction: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusCreated, t)
	return
}
