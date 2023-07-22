package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	client, _ := NewClient("John doe", "john@doe.com")
	account := NewAccount(client)

	client2, _ := NewClient("John doe 2", "john@doe2.com")
	account2 := NewAccount(client2)

	account.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, float64(1100), account2.Balance)
	assert.Equal(t, float64(900), account.Balance)
}

func TestNewTransactionWithInsufficientFunds(t *testing.T) {
	client, _ := NewClient("John doe", "john@doe.com")
	account := NewAccount(client)

	client2, _ := NewClient("John doe 2", "john@doe2.com")
	account2 := NewAccount(client2)

	account.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account, account2, 2000)

	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, float64(1000), account2.Balance)
	assert.Equal(t, float64(1000), account.Balance)
}
