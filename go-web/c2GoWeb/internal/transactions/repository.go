package transactions

import (
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
)

const (
	TransactionNotFound = "transaction not found"
	CantUpdate          = "could not update transaction"
	CantDelete          = "could not delete transaction"
)

type Repository interface {
	FindAll() ([]Transaction, error)
	FindOne(id int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(id int, transaction Transaction) (Transaction, error)
	PartialUpdate(id int, transactionCode string, amount float64) (Transaction, error)
	Remove(id int) (int, error)
}

type repository struct {
	db store.Store
}

type Transaction struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transactionCode" binding:"required"`
	Currency        string  `json:"currency" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	Sender          string  `json:"sender" binding:"required"`
	Receiver        string  `json:"receiver" binding:"required"`
	Date            string  `json:"date" binding:"required"` // TODO: con db cambiar a time.Time
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) getLastId() (lastId int, err error) {
	var ts []Transaction
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
	var ts []Transaction
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

func (r *repository) FindAll() (ts []Transaction, err error) {
	err = r.db.Read(&ts)
	return
}

func (r *repository) FindOne(id int) (t Transaction, err error) {
	var ts []Transaction
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

func (r *repository) Save(transaction Transaction) (t Transaction, err error) {
	lastId, err := r.getLastId()
	transaction.Id = lastId + 1
	var ts []Transaction
	err = r.db.Read(&ts)
	ts = append(ts, transaction)
	t = transaction
	err = r.db.Write(&ts)
	return
}

func (r *repository) Update(id int, transaction Transaction) (t Transaction, err error) {
	update, err := r.FindOne(id)
	if err != nil {
		err = fmt.Errorf("%s. %w", CantUpdate, err)
		return
	}
	index, _ := r.findIndex(update.Id)
	var ts []Transaction
	err = r.db.Read(&ts)
	oid := ts[index].Id
	ts[index] = transaction
	ts[index].Id = oid
	err = r.db.Write(&ts)
	t = ts[index]
	return
}

func (r *repository) PartialUpdate(id int, transactionCode string, amount float64) (t Transaction, err error) {
	update, err := r.FindOne(id)
	if err != nil {
		err = fmt.Errorf("%s. %w", CantUpdate, err)
		return
	}
	index, _ := r.findIndex(update.Id)
	var ts []Transaction
	err = r.db.Read(&ts)
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
	var ts []Transaction
	err = r.db.Read(&ts)
	ts = append(ts[:index], ts[index+1:]...)
	err = r.db.Write(&ts)
	deletedId = id
	return
}
