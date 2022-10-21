package transactions

import (
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
)

type mockRepository struct {
	db   store.Store
	repo Repository
	err  error
}

func NewMockRepository(db store.Store, repo Repository, err error) mockRepository {
	return mockRepository{
		db:   db,
		repo: repo,
		err:  err,
	}
}

func (r *mockRepository) FindAll() (ts []domain.Transaction, err error) {
	return
}

func (r *mockRepository) FindOne(id int) (t domain.Transaction, err error) {
	return
}

func (r *mockRepository) Save(transaction domain.Transaction) (t domain.Transaction, err error) {
	return
}

func (r *mockRepository) Update(id int, transaction domain.Transaction) (t domain.Transaction, err error) {
	return
}

func (r *mockRepository) PartialUpdate(id int, transactionCode string, amount float64) (t domain.Transaction, err error) {
	return
}

func (r *mockRepository) Remove(id int) (deletedID int, err error) {
	if r.err != nil {
		return 0, r.err
	}
	return r.repo.Remove(id)
}
