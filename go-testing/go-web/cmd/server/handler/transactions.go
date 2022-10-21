package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/transactions"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/web"
	"net/http"
	"strconv"
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

type PATCHRequest struct {
	TransactionCode string  `json:"transactionCode" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}

func NewHandler(service transactions.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func NewTransactionFromRequest(request Request) domain.Transaction {
	return domain.Transaction{
		TransactionCode: request.TransactionCode,
		Currency:        request.Currency,
		Amount:          request.Amount,
		Sender:          request.Sender,
		Receiver:        request.Receiver,
		Date:            request.Date,
	}
}

func (c *Controller) HelloMessageHandler(ctx *gin.Context) {
	name := ctx.Request.URL.Query().Get("name")
	ctx.JSON(
		http.StatusOK,
		web.NewResponse(
			200,
			gin.H{
				"message": fmt.Sprintf("Hola, %s!", name),
				"warning": "Always query your name with the 'name' variable",
				"tip":     "Spaces are represented by %20 (Ej: ?name=Lucas%20Soria)",
			},
			"",
		),
	)
}

// GetAll
// @Summary List transactions
// @Produce json
// @Success 200 {object} web.Response
// @Router  /transactions [get]
func (c *Controller) GetAll(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	ts, _ := c.service.GetAll(queries)
	ctx.JSON(
		http.StatusOK,
		web.NewResponse(
			200,
			ts,
			"",
		),
	)
}

// GetOne
// @Summary Retrieve one transaction
// @Param   id path int true "transaction ID"
// @Produce json
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 400 {object} web.Response
// @Router  /transactions/{id} [get]
func (c *Controller) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.NewResponse(
				400,
				nil,
				fmt.Sprintf("while converting id: %s", err),
			),
		)
		return
	}
	t, err := c.service.GetOne(id)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.NewResponse(
				404,
				nil,
				fmt.Sprintf("while retrieving transaction. %s", err),
			),
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		web.NewResponse(
			200,
			t,
			"",
		),
	)
}

func validateFields(err error) (errs string) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errs = "this fields are required: "
		for i, fe := range ve {
			if i == len(ve)-1 {
				errs += fmt.Sprintf("'%s'", fe.Field())
			} else {
				errs += fmt.Sprintf("'%s', ", fe.Field())
			}
		}
	}
	return
}

// Create
// @Summary Create a new transaction
// @Accept  json
// @Produce json
// @Param   token       header   string  true "Security Token"
// @Param   transaction body     Request true "Transaction to be created"
// @Success 201         {object} web.Response
// @Failure 401         {object} web.Response
// @Failure 400         {object} web.Response
// @Router  /transactions/ [post]
func (c *Controller) Create(ctx *gin.Context) {
	var r Request
	if err := ctx.ShouldBindJSON(&r); err != nil {
		errs := validateFields(err)
		ctx.JSON(
			http.StatusBadRequest,
			web.NewResponse(
				400,
				nil,
				errs,
			),
		)
		return
	}
	t, _ := c.service.Store(NewTransactionFromRequest(r))
	ctx.JSON(
		http.StatusCreated,
		web.NewResponse(
			201,
			t,
			"",
		),
	)
}

// Update
// @Summary Update existing transaction
// @Accept  json
// @Produce json
// @Param   id          path     int     true "transaction ID"
// @Param   token       header   string  true "Security Token"
// @Param   transaction body     Request true "Updated values of the transaction"
// @Success 200         {object} web.Response
// @Failure 400         {object} web.Response
// @Failure 404         {object} web.Response
// @Router  /transactions/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.NewResponse(
				400,
				nil,
				fmt.Sprintf("while converting id: %s", errId),
			),
		)
		return
	}
	var r Request
	if errBind := ctx.ShouldBindJSON(&r); errBind != nil {
		errs := validateFields(errBind)
		ctx.JSON(
			http.StatusBadRequest,
			web.NewResponse(
				400,
				nil,
				errs,
			),
		)
		return
	}
	t := NewTransactionFromRequest(r)
	tr, err := c.service.Update(id, t)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.NewResponse(
				404,
				nil,
				fmt.Sprintf("while updating transaction: %s", err),
			),
		)
		return
	}
	ctx.JSON(http.StatusOK,
		web.NewResponse(
			200,
			tr,
			"",
		),
	)
}

// PartialUpdate
// @Summary Update some values of an existing transaction
// @Accept  json
// @Produce json
// @Param   id          path     int     true "transaction ID"
// @Param   token       header   string  true "Security Token"
// @Param   transaction body     Request true "Updated values of the transaction"
// @Success 200         {object} web.Response
// @Failure 400         {object} web.Response
// @Failure 404         {object} web.Response
// @Router  /transactions/{id} [patch]
func (c *Controller) PartialUpdate(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.NewResponse(
				400,
				nil,
				fmt.Sprintf("while converting id: %s", errId),
			),
		)
		return
	}
	var pr PATCHRequest
	if err := ctx.ShouldBindJSON(&pr); err != nil {
		errs := validateFields(err)
		ctx.JSON(
			http.StatusBadRequest,
			web.NewResponse(
				400,
				nil,
				errs,
			),
		)
		return
	}
	tr, err := c.service.PartialUpdate(id, pr.TransactionCode, pr.Amount)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.NewResponse(
				404,
				nil,
				fmt.Sprintf("while updating transaction: %s", err),
			),
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		web.NewResponse(
			200,
			tr,
			"",
		),
	)
}

// Delete
// @Summary Delete transaction
// @Produce json
// @Param   id    path   int    true "transaction ID"
// @Param   token header string true "Security Token"
// @Success 204
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router  /transactions/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.NewResponse(
				400,
				nil,
				fmt.Sprintf("while converting id: %s", errId),
			),
		)
		return
	}
	_, errRemove := c.service.Remove(id)
	if errRemove != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.NewResponse(
				404,
				nil,
				fmt.Sprintf("while deleting transaction: %s", errRemove),
			),
		)
		return
	}
	ctx.Status(http.StatusNoContent)
}
