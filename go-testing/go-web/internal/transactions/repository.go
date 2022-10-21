package transactions

import (
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
)

const (
	TransactionNotFound = "transaction not found"
	CantUpdate          = "could not update transaction"
	CantDelete          = "could not delete transaction"
)

type Repository interface {
	FindAll() ([]domain.Transaction, error)
	FindOne(id int) (domain.Transaction, error)
	Save(transaction domain.Transaction) (domain.Transaction, error)
	Update(id int, transaction domain.Transaction) (domain.Transaction, error)
	PartialUpdate(id int, transactionCode string, amount float64) (domain.Transaction, error)
	Remove(id int) (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) getLastId() (lastId int, err error) {
	var ts []domain.Transaction
	if err = r.db.Read(&ts); err != nil {
		return
	}
	if len(ts) != 0 {
		lastId = ts[len(ts)-1].Id
	}
	for _, t := range ts {
		if t.Id > lastId {
			lastId = t.Id
		}
	}
	return
}

func (r *repository) findIndex(id int) (index int, err error) {
	var ts []domain.Transaction
	err = r.db.Read(&ts)
	for i, t := range ts {
		if t.Id == id {
			index = i
			return
		}
	}
	err = fmt.Errorf("%s: id %d", TransactionNotFound, id)
	return
}

func (r *repository) FindAll() (ts []domain.Transaction, err error) {
	err = r.db.Read(&ts)
	return
}

func (r *repository) FindOne(id int) (t domain.Transaction, err error) {
	var ts []domain.Transaction
	err = r.db.Read(&ts)
	for _, tr := range ts {
		if tr.Id == id {
			t = tr
			return
		}
	}
	err = fmt.Errorf("%s: id %d", TransactionNotFound, id)
	return
}

func (r *repository) Save(transaction domain.Transaction) (t domain.Transaction, err error) {
	var lastId int
	lastId, err = r.getLastId()
	if err != nil {
		return
	}
	transaction.Id = lastId + 1
	var ts []domain.Transaction
	err = r.db.Read(&ts)
	if err != nil {
		return
	}
	ts = append(ts, transaction)
	t = transaction
	err = r.db.Write(&ts)
	return
}

func (r *repository) Update(id int, transaction domain.Transaction) (t domain.Transaction, err error) {
	update, err := r.FindOne(id)
	if err != nil {
		err = fmt.Errorf("%s. %w", CantUpdate, err)
		return
	}
	index, _ := r.findIndex(update.Id)
	var ts []domain.Transaction
	err = r.db.Read(&ts)
	if err != nil {
		return
	}
	oid := ts[index].Id
	ts[index] = transaction
	ts[index].Id = oid
	err = r.db.Write(&ts)
	t = ts[index]
	return
}

func (r *repository) PartialUpdate(id int, transactionCode string, amount float64) (t domain.Transaction, err error) {
	update, err := r.FindOne(id)
	if err != nil {
		err = fmt.Errorf("%s. %w", CantUpdate, err)
		return
	}
	index, _ := r.findIndex(update.Id)
	var ts []domain.Transaction
	err = r.db.Read(&ts)
	if err != nil {
		return
	}
	ts[index].TransactionCode = transactionCode
	ts[index].Amount = amount
	err = r.db.Write(&ts)
	t = ts[index]
	return
}

func (r *repository) Remove(id int) (deletedId int, err error) {
	index, err := r.findIndex(id)
	if err != nil {
		err = fmt.Errorf("%s. %w", CantDelete, err)
		return
	}
	var ts []domain.Transaction
	err = r.db.Read(&ts)
	if err != nil {
		return
	}
	ts = append(ts[:index], ts[index+1:]...)
	err = r.db.Write(&ts)
	deletedId = id
	return
}
