package store

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

type Stub struct {
	Transactions []domain.Transaction
}

func (s *Stub) Read(data interface{}) (err error) {
	castedData := data.(*[]domain.Transaction)
	*castedData = s.Transactions
	return
}

func (s *Stub) Write(data interface{}) (err error) {
	return
}
