package entity_test

import (
	"github.com/loxt/walletcore/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := entity.NewClient("John doe", "john@doe.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John doe", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
}

func TestNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := entity.NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}
