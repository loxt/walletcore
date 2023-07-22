package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	client, _ := NewClient("John doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, account.Client, client)
}

func TestNewAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestDepositAccount(t *testing.T) {
	client, _ := NewClient("John doe", "j@j.com")
	account := NewAccount(client)
	account.Deposit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John doe", "j@j.com")
	account := NewAccount(client)
	account.Deposit(100)
	account.Debit(25)
	assert.Equal(t, float64(75), account.Balance)
}
