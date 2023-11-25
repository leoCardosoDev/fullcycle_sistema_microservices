package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "joe@doe.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "joe@doe.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@doe.com")
	err := client.Update("John Doe Update", "joe_update@doe.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "joe_update@doe.com", client.Email)
}

func TestUpdateClientWhenInvalidName(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@doe.com")
	err := client.Update("", "joe@doe.com")
	assert.Error(t, err, "name is required")
}

func TestUpdateClientWhenInvalidEmail(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@doe.com")
	err := client.Update("John Doe", "")
	assert.Error(t, err, "email is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@doe.com")
	account := NewAccount(client)
	err := client.AddAccount((account))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccountWhenInvalidClient(t *testing.T) {
	client1, _ := NewClient("John Doe", "joe@doe.com")
	client2, _ := NewClient("Ane Truck", "ane@truck.com")
	account := NewAccount(client2)
	err := client1.AddAccount(account)
	assert.Error(t, err, "account does not belong to client")
}
