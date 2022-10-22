package transactions

import (
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/pkg/store"
	"github.com/stretchr/testify/assert"
	"net/url"
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
	searchID := 1
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "USD",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	data, err := transactionService.Update(searchID, expected)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
	assert.True(t, mockStore.ReadInvoked)
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
	searchID := 1
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "USD",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	id, err := transactionService.Remove(searchID)
	assert.Nil(t, err)
	assert.Equal(t, expected.Id, id)
	assert.True(t, mockStore.ReadInvoked)
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
	searchID := 1900
	expectedErr := fmt.Sprintf("%s. %s: id %d", CantDelete, TransactionNotFound, searchID)
	expected := 0
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	id, err := transactionService.Remove(searchID)
	assert.EqualError(t, err, expectedErr)
	assert.Equal(t, expected, id)
	assert.True(t, mockStore.ReadInvoked)
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
	searchID := 1900
	expectedErr := fmt.Errorf("%s. %s: id %d", CantDelete, TransactionNotFound, searchID)
	expected := 0
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	mockRepository := MockRepository{db: &mockStore, repo: transactionRepository, RemoveErr: expectedErr}
	transactionService := NewService(&mockRepository)
	id, err := transactionService.Remove(searchID)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, expected, id)
}

func TestService_Remove(t *testing.T) {
	t.Run("Happy Path", removeOK)
	t.Run("Sad Path", removeNotFound)
	t.Run("Forced Sad Path", removeNotFoundForced)
}

func testGetAll(t *testing.T) {
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
	query := url.Values{}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	data, err := transactionService.GetAll(query)
	assert.Nil(t, err)
	assert.Equal(t, mockTransactions, data)
}

func testGetAllQueried(t *testing.T) {
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
	query := url.Values{}
	query.Set("currency", "USD")
	expected := []domain.Transaction{
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
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	data, err := transactionService.GetAll(query)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}

func testGetAllErrRepository(t *testing.T) {
	expectedErr := fmt.Errorf("forced error reading storage")
	expected := make([]domain.Transaction, 0)
	query := url.Values{}
	mockStore := store.Mock{Transactions: []domain.Transaction{}, ErrRead: expectedErr}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	data, err := transactionService.GetAll(query)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, expected, data)
}

func TestService_GetAll(t *testing.T) {
	t.Run("Happy simple Path", testGetAll)
	t.Run("Happy not simple Path", testGetAllQueried)
	t.Run("Forced error in repository", testGetAllErrRepository)
}

func TestService_Store(t *testing.T) {
	mockTransactions := []domain.Transaction{
		{
			Id:              2,
			TransactionCode: "ovd98dfg8dfs",
			Currency:        "ARS",
			Amount:          215.53,
			Sender:          "987asd9asd",
			Receiver:        "89as99a9",
			Date:            "2022-10-20T00:00:00-03:00",
		},
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
	expected := domain.Transaction{
		Id:              3,
		TransactionCode: "ovd98dfg8dfs",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	data, err := transactionService.Store(expected)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}

func TestService_PartialUpdate(t *testing.T) {
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
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "ARS",
		Amount:          0,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	data, err := transactionService.PartialUpdate(searchID, "After Update", 0)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}

func TestService_GetOne(t *testing.T) {
	mockTransactions := []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "kjnas76ask",
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
		TransactionCode: "kjnas76ask",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	mockStore := store.Mock{Transactions: mockTransactions}
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	data, err := transactionService.GetOne(searchID)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}
