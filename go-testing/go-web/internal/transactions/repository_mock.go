package transactions

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
)

type MockRepository struct {
	db        store.Store
	repo      Repository
	RemoveErr error
}

func (r *MockRepository) FindAll() (ts []domain.Transaction, err error) {
	return
}

func (r *MockRepository) FindOne(id int) (t domain.Transaction, err error) {
	return
}

func (r *MockRepository) Save(transaction domain.Transaction) (t domain.Transaction, err error) {
	return
}

func (r *MockRepository) Update(id int, transaction domain.Transaction) (t domain.Transaction, err error) {
	return
}

func (r *MockRepository) PartialUpdate(id int, transactionCode string, amount float64) (t domain.Transaction, err error) {
	return
}

func (r *MockRepository) Remove(id int) (deletedID int, err error) {
	if r.RemoveErr != nil {
		return 0, r.RemoveErr
	}
	return r.repo.Remove(id)
}
