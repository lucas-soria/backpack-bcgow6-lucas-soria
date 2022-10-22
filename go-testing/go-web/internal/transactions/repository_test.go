package transactions

import (
	"errors"
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_FindAll(t *testing.T) {
	stubTransactions := []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "3345fsd",
			Currency:        "ARS",
			Amount:          215.53,
			Sender:          "987asd9asd",
			Receiver:        "89as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
		{
			Id:              2,
			TransactionCode: "3345fse",
			Currency:        "USD",
			Amount:          30.67,
			Sender:          "987ssd9asd",
			Receiver:        "80as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
	}
	stubStore := store.Stub{Transactions: stubTransactions}
	transactionRepository := NewRepository(&stubStore)
	data, err := transactionRepository.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, stubTransactions, data)
}

func TestRepository_FindOneNotFound(t *testing.T) {
	mockTransactions := []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "Before Update",
			Currency:        "ARS",
			Amount:          215.53,
			Sender:          "987asd9asd",
			Receiver:        "89as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
	}
	searchID := 10
	expectedErr := fmt.Errorf("%s: id %d", TransactionNotFound, searchID)
	expected := domain.Transaction{}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	data, err := transactionRepository.FindOne(searchID)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expected, data)
}

func testUpdate(t *testing.T) {
	mockTransactions := []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "Before Update",
			Currency:        "ARS",
			Amount:          215.53,
			Sender:          "987asd9asd",
			Receiver:        "89as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
		{
			Id:              2,
			TransactionCode: "3345fse",
			Currency:        "USD",
			Amount:          30.67,
			Sender:          "987ssd9asd",
			Receiver:        "80as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
	}
	searchID := 1
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	data, err := transactionRepository.Update(searchID, expected)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
	assert.True(t, mockStore.ReadInvoked)
}

func testUpdateNotFound(t *testing.T) {
	mockTransactions := []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "Before Update",
			Currency:        "ARS",
			Amount:          215.53,
			Sender:          "987asd9asd",
			Receiver:        "89as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
	}
	searchID := 10
	expectedErr := fmt.Sprintf("%s. %s: id %d", CantUpdate, TransactionNotFound, searchID)
	expected := domain.Transaction{}
	update := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	data, err := transactionRepository.Update(searchID, update)
	assert.EqualError(t, err, expectedErr)
	assert.Equal(t, expected, data)
	assert.True(t, mockStore.ReadInvoked)
}

func TestRepository_Update(t *testing.T) {
	t.Run("Happy Path", testUpdate)
	t.Run("Not found transaction by ID to update", testUpdateNotFound)
}

func testSaveErrRead(t *testing.T) {
	expectedErr := errors.New("forced error reading storage")
	expected := domain.Transaction{}
	save := domain.Transaction{
		Id:              1,
		TransactionCode: "kjas87sjk",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: []domain.Transaction{}, ErrRead: expectedErr}
	transactionRepository := NewRepository(&mockStore)
	data, err := transactionRepository.Save(save)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, expected, data)
}

func testSaveErrWrite(t *testing.T) {
	expectedErr := errors.New("forced error writing to storage")
	expected := domain.Transaction{}
	save := domain.Transaction{
		Id:              1,
		TransactionCode: "kjas87sjk",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: []domain.Transaction{}, ErrWrite: expectedErr}
	transactionRepository := NewRepository(&mockStore)
	data, err := transactionRepository.Save(save)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, expected, data)
}

func TestRepository_Save(t *testing.T) {
	// Save tested on service
	t.Run("Forced error reading storage", testSaveErrRead)
	t.Run("Forced error writing to storage", testSaveErrWrite)
}

func testPartialUpdateNotFound(t *testing.T) {
	searchID := 1
	expectedErr := fmt.Errorf("could not update transaction. transaction not found: id %d", searchID)
	expected := domain.Transaction{}
	mockStore := store.Mock{Transactions: []domain.Transaction{}}
	transactionRepository := NewRepository(&mockStore)
	data, err := transactionRepository.PartialUpdate(searchID, "After Update", 0)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, expected, data)
}

func testPartialUpdateErrWrite(t *testing.T) {
	mockTransactions := []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "Before Update",
			Currency:        "ARS",
			Amount:          215.53,
			Sender:          "987asd9asd",
			Receiver:        "89as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
	}
	searchID := 1
	expectedErr := fmt.Errorf("forced error writing to storage")
	expected := domain.Transaction{}
	mockStore := store.Mock{Transactions: mockTransactions, ErrWrite: expectedErr}
	transactionRepository := NewRepository(&mockStore)
	data, err := transactionRepository.PartialUpdate(searchID, "After Update", 0)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, expected, data)
}

func TestPartialUpdate(t *testing.T) {
	t.Run("Not found transaction to partially update", testPartialUpdateNotFound)
	t.Run("Forced error writing to storage", testPartialUpdateErrWrite)
}
