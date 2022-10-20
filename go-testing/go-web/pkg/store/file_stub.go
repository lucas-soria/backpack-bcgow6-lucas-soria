package store

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

type Stub struct {
	transactions []domain.Transaction
}

var StubTransactions = []domain.Transaction{
	{
		Id:              1,
		TransactionCode: "3345fsd",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	},
	{
		Id:              2,
		TransactionCode: "3345fse",
		Currency:        "USD",
		Amount:          30.67,
		Sender:          "987ssd9asd",
		Receiver:        "80as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	},
}

func NewStub() Store {
	return &Stub{
		transactions: StubTransactions,
	}
}

func (s *Stub) Read(data interface{}) (err error) {
	castedData := data.(*[]domain.Transaction)
	*castedData = s.transactions
	return
}

func (s *Stub) Write(data interface{}) (err error) {
	return
}
