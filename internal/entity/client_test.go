package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Ben Bao", "ben@bao.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Ben Bao", client.Name)
	assert.Equal(t, "ben@bao.com", client.Email)
}

func TestCreateNewClientArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("ben bao", "ben@ba.com")
	err := client.Update("Ben Bao", "ben@bao.com")
	assert.Nil(t, err)
	assert.Equal(t, "Ben Bao", client.Name)
	assert.Equal(t, "ben@bao.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Ben Bao", "ben@bao.com")
	err := client.Update("", "b@b.com")
	assert.Error(t, err, "name is required")

}

func TestAddAcclountToClient(t *testing.T) {
	client, _ := NewClient("Ben Bao", "ben@bao.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
