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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	query := url.Values{}
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
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
	data, err := transactionService.GetAll(query)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}

func testGetAllErrRepository(t *testing.T) {
	mockStore := store.NewMock([]domain.Transaction{})
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	query := url.Values{}
	mockStore.ErrRead = fmt.Errorf("forced error reading storage")
	data, err := transactionService.GetAll(query)
	assert.EqualError(t, err, mockStore.ErrRead.Error())
	assert.Equal(t, []domain.Transaction{}, data)
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	expected := domain.Transaction{
		Id:              3,
		TransactionCode: "ovd98dfg8dfs",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "After Update",
		Currency:        "ARS",
		Amount:          0,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	data, err := transactionService.PartialUpdate(1, "After Update", 0)
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
	mockStore := store.NewMock(mockTransactions)
	transactionRepository := NewRepository(&mockStore)
	transactionService := NewService(transactionRepository)
	expected := domain.Transaction{
		Id:              1,
		TransactionCode: "kjnas76ask",
		Currency:        "ARS",
		Amount:          215.53,
		Sender:          "987asd9asd",
		Receiver:        "89as99a9",
		Date:            "2022-10-20T00:00:00-03:00",
	}
	data, err := transactionService.GetOne(1)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}
