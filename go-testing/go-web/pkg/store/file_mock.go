package store

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

type Mock struct {
	transactions []domain.Transaction
	ErrNotfound  error
	ReadInvoked  bool
}

func NewMock(transactions []domain.Transaction) Mock {
	return Mock{
		transactions: transactions,
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
