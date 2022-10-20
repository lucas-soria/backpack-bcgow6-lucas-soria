package transactions

import (
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

func TestRepository_UpdateName(t *testing.T) {
	mockStore := store.NewMock()
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
	transactionRepository.Update(1, expected)
	data, _ := transactionRepository.FindOne(1)
	assert.Equal(t, expected, data)
	assert.Equal(t, true, mockStore.ReadInvoked)
}