package transactions

import (
	"fmt"
	"net/url"
	"reflect"
)

type Service interface {
	GetAll(queries url.Values) ([]Transaction, error)
	GetOne(id int) (Transaction, error)
	Store(transaction Transaction) (Transaction, error)
	Update(id int, transaction Transaction) (Transaction, error)
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

func (s *service) GetAll(queries url.Values) (ts []Transaction, err error) {
	repositoryTs, _ := s.repository.FindAll()
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

func (s *service) GetOne(id int) (t Transaction, err error) {
	t, err = s.repository.FindOne(id)
	if err != nil {
		err = fmt.Errorf("error: problems in repository: %w", err)
	}
	return
}

func (s *service) Store(transaction Transaction) (t Transaction, err error) {
	tr, SaveErr := s.repository.Save(transaction)
	if SaveErr != nil {
		err = fmt.Errorf("error: problems in repository: %w", err)
	}
	t = tr
	return
}

func (s *service) Update(id int, transaction Transaction) (t Transaction, err error) {
	return
}

func (s *service) Remove(id int) (_ int, err error) {
	return id, err
}
