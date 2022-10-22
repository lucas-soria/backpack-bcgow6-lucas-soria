package store

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

type Mock struct {
	Transactions []domain.Transaction
	ErrRead      error
	ErrWrite     error
	ReadInvoked  bool
	WriteInvoked bool
}

func (s *Mock) Read(data interface{}) (err error) {
	s.ReadInvoked = true
	if s.ErrRead != nil {
		return s.ErrRead
	}
	castedData := data.(*[]domain.Transaction)
	*castedData = s.Transactions
	return
}

func (s *Mock) Write(data interface{}) (err error) {
	s.WriteInvoked = true
	if s.ErrWrite != nil {
		return s.ErrWrite
	}
	castedData := data.(*[]domain.Transaction)
	s.Transactions = append(s.Transactions, *castedData...)
	return
}
