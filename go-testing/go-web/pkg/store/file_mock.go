package store

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

type Mock struct {
	transactions []domain.Transaction
	ReadInvoked  bool
}

var MockTransactions = []domain.Transaction{
	{
		Id:              1,
		TransactionCode: "Before Update",
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

func NewMock() Mock {
	return Mock{
		transactions: MockTransactions,
	}
}

func (s *Mock) Read(data interface{}) (err error) {
	s.ReadInvoked = true
	castedData := data.(*[]domain.Transaction)
	*castedData = s.transactions
	return
}

func (s *Mock) Write(data interface{}) (err error) {
	castedData := data.(*[]domain.Transaction)
	s.transactions = append(s.transactions, *castedData...)
	return
}
