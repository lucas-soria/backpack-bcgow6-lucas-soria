package transactions

import (
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Update(t *testing.T) {
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
	transactionService := NewService(transactionRepository)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "USD",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	data, err := transactionService.Update(1, expected)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
	assert.Equal(t, true, mockStore.ReadInvoked)
}

func removeOK(t *testing.T) {
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
	transactionService := NewService(transactionRepository)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "USD",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	id, err := transactionService.Remove(1)
	assert.Nil(t, err)
	assert.Equal(t, expected.Id, id)
	assert.Equal(t, true, mockStore.ReadInvoked)
}

func removeNotFound(t *testing.T) {
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
	transactionService := NewService(transactionRepository)
	expected := 0
	id, err := transactionService.Remove(1900)
	assert.EqualError(t, err, fmt.Sprintf("%s. %s: id %d", CantDelete, TransactionNotFound, 1900))
	assert.Equal(t, expected, id)
	assert.Equal(t, true, mockStore.ReadInvoked)
}

func removeNotFoundForced(t *testing.T) {
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
	expected := 0
	expectedErr := fmt.Errorf("%s. %s: id %d", CantDelete, TransactionNotFound, 1900)
	repo := NewRepository(&mockStore)
	transactionRepository := NewMockRepository(&mockStore, repo, expectedErr)
	transactionService := NewService(&transactionRepository)
	id, err := transactionService.Remove(1900)
	assert.Equal(t, err, expectedErr)
	assert.Equal(t, expected, id)
}

func TestService_Remove(t *testing.T) {
	t.Run("Happy Path", removeOK)
	t.Run("Sad Path", removeNotFound)
	t.Run("Forced Sad Path", removeNotFoundForced)
}
