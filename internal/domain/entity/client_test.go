package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John doe", "john@doe.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John doe", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
}

func TestNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John doe", "john@doe.com")

	err := client.Update("John doe Updated", "john@gmail.com")

	assert.Nil(t, err)
	assert.Equal(t, "John doe Updated", client.Name)
	assert.Equal(t, "john@gmail.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John doe", "john@doe.com")

	err := client.Update("", "john@gmail.com")

	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("John doe", "john@doe.com")
	account := NewAccount(client)
	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
