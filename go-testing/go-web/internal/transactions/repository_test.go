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
	stubStore := store.NewStub()
	transactionRepository := NewRepository(stubStore)
	expected := store.StubTransactions
	data, _ := transactionRepository.FindAll()
	assert.Equal(t, expected, data)
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	expected := domain.Transaction{}
	expectedErr := fmt.Errorf("%s: id %d", TransactionNotFound, 10)
	data, err := transactionRepository.FindOne(10)
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	data, err := transactionRepository.Update(1, expected)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
	assert.Equal(t, true, mockStore.ReadInvoked)
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	expectedErr := fmt.Sprintf("%s. %s: id %d", CantUpdate, TransactionNotFound, 10)
	data, err := transactionRepository.Update(10, expected)
	assert.EqualError(t, err, expectedErr)
	assert.Equal(t, domain.Transaction{}, data)
	assert.Equal(t, true, mockStore.ReadInvoked)
}

func TestRepository_Update(t *testing.T) {
	t.Run("Happy Path", testUpdate)
	t.Run("Not found transaction to update", testUpdateNotFound)
}

func testSaveErrRead(t *testing.T) {
	mockStore := store.NewMock([]domain.Transaction{})
	transactionRepository := NewRepository(&mockStore)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "kjas87sjk",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore.ErrRead = errors.New("error reading storage")
	data, err := transactionRepository.Save(expected)
	assert.EqualError(t, err, mockStore.ErrRead.Error())
	assert.Equal(t, domain.Transaction{}, data)
}

func testSaveErrWrite(t *testing.T) {
	mockStore := store.NewMock([]domain.Transaction{})
	transactionRepository := NewRepository(&mockStore)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "kjas87sjk",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore.ErrWrite = errors.New("error saving in storage")
	data, err := transactionRepository.Save(expected)
	assert.EqualError(t, err, mockStore.ErrWrite.Error())
	assert.Equal(t, domain.Transaction{}, data)
}

func TestRepository_Save(t *testing.T) {
	// Save tested on service
	t.Run("Forced error reading storage", testSaveErrRead)
	t.Run("Forced error writing to storage", testSaveErrWrite)
}

func testPartialUpdateNotFound(t *testing.T) {
	mockStore := store.NewMock([]domain.Transaction{})
	transactionRepository := NewRepository(&mockStore)
	expected := domain.Transaction{}
	expectedErr := fmt.Errorf("could not update transaction. transaction not found: id %d", 1)
	data, err := transactionRepository.PartialUpdate(1, "After Update", 0)
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	expected := domain.Transaction{}
	mockStore.ErrWrite = fmt.Errorf("forced error writing to storage")
	data, err := transactionRepository.PartialUpdate(1, "After Update", 0)
	assert.EqualError(t, err, mockStore.ErrWrite.Error())
	assert.Equal(t, expected, data)
}

func TestPartialUpdate(t *testing.T) {
	t.Run("Not found transaction to partially update", testPartialUpdateNotFound)
	t.Run("Forced error writing to storage", testPartialUpdateErrWrite)
}
