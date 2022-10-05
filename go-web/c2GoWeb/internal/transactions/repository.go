package transactions

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	TransactionNotFound = "transaction not found"
	CantUpdate          = "problems updating transaction"
	CantDelete          = "problems deleting transaction"
)

type Repository interface {
	FindAll() ([]Transaction, error)
	FindOne(id int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(id int, transaction Transaction) (Transaction, error)
	Remove(id int) (int, error)
}

type repository struct {
	ts     []Transaction
	LastId int
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

func NewRepository() Repository {
	var r = &repository{}
	r.ReadFile()
	return r
}

func (r *repository) findIndex(id int) (index int, err error) {
	for i, t := range r.ts {
		if t.Id == id {
			index = i
			return
		}
	}
	err = fmt.Errorf("%s: transaction with id %d", TransactionNotFound, id)
	return
}

func (r *repository) getLastId() (lastId int) {
	if len(r.ts) != 0 {
		r.LastId = r.ts[len(r.ts)-1].Id
	}
	for _, t := range r.ts {
		if t.Id > r.LastId {
			r.LastId = t.Id
		}
	}
	return
}

func (r *repository) ReadFile() {
	if textBytes, err := os.ReadFile("./transactions.json"); err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	} else if err1 := json.Unmarshal(textBytes, &r.ts); err1 != nil {
		panic(fmt.Sprintf("error: %v", err1.Error()))
	}
	r.getLastId()
	fmt.Println(r.LastId)
}

func (r *repository) FindAll() (ts []Transaction, err error) {
	return r.ts, nil
}

func (r *repository) FindOne(id int) (t Transaction, err error) {
	for _, tr := range r.ts {
		if tr.Id == id {
			t = tr
			return
		}
	}
	err = fmt.Errorf("%s: transaction with id %d", TransactionNotFound, id)
	return
}

func (r *repository) Save(transaction Transaction) (t Transaction, err error) {
	r.LastId++
	transaction.Id = r.LastId
	r.ts = append(r.ts, transaction)
	t = transaction
	return
}

func (r *repository) Update(id int, transaction Transaction) (t Transaction, err error) {
	update, err := r.FindOne(id)
	if err != nil {
		err = fmt.Errorf("%s:\n\t%w", CantUpdate, err)
		return
	}
	index, _ := r.findIndex(update.Id)
	r.ts[index].TransactionCode = transaction.TransactionCode
	r.ts[index].Currency = transaction.Currency
	r.ts[index].Amount = transaction.Amount
	r.ts[index].Sender = transaction.Sender
	r.ts[index].Receiver = transaction.Receiver
	r.ts[index].Date = transaction.Date
	t = r.ts[index]
	return
}

func (r *repository) Remove(id int) (deletedId int, err error) {
	index, err := r.findIndex(id)
	if err != nil {
		err = fmt.Errorf("%s:\n\t%w", CantDelete, err)
		return
	}
	copy(r.ts[index:], r.ts[index+1:])
	r.ts[len(r.ts)-1] = Transaction{}
	r.ts = r.ts[:len(r.ts)-1]
	deletedId = id
	return
}
