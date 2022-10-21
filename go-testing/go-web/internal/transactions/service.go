package transactions

import (
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"net/url"
	"reflect"
)

type Service interface {
	GetAll(queries url.Values) ([]domain.Transaction, error)
	GetOne(id int) (domain.Transaction, error)
	Store(transaction domain.Transaction) (domain.Transaction, error)
	Update(id int, transaction domain.Transaction) (domain.Transaction, error)
	PartialUpdate(id int, transactionCode string, amount float64) (domain.Transaction, error)
	Remove(id int) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(queries url.Values) (ts []domain.Transaction, err error) {
	var repositoryTs []domain.Transaction
	repositoryTs, err = s.repository.FindAll()
	if err != nil {
		ts = []domain.Transaction{}
		return
	}
	for _, t := range repositoryTs {
		reflection := reflect.ValueOf(t)
		b := true
		for i := 0; i < reflection.NumField(); i++ {
			if name := reflection.Type().Field(i).Tag.Get("json"); queries.Has(name) {
				if fmt.Sprintf("%v", reflection.Field(i)) != queries.Get(name) {
					b = false
					break
				}
			}
		}
		if b {
			ts = append(ts, t)
		}
	}
	return
}

func (s *service) GetOne(id int) (t domain.Transaction, err error) {
	t, err = s.repository.FindOne(id)
	return
}

func (s *service) Store(transaction domain.Transaction) (t domain.Transaction, err error) {
	t, _ = s.repository.Save(transaction)
	return
}

func (s *service) Update(id int, transaction domain.Transaction) (t domain.Transaction, err error) {
	t, err = s.repository.Update(id, transaction)
	return
}

func (s *service) PartialUpdate(id int, transactionCode string, amount float64) (t domain.Transaction, err error) {
	t, err = s.repository.PartialUpdate(id, transactionCode, amount)
	return
}

func (s *service) Remove(id int) (idDeleted int, err error) {
	idDeleted, err = s.repository.Remove(id)
	return
}
