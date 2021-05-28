package crypto

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-app-engine-demo/pkg/entity"
	"testing"
	"time"
)

func TestEncryptData(t *testing.T) {
	//given
	data, _ := json.Marshal(generatePerson())
	key := "person"
	c := NewCrypto(key, data)

	//when
	err := c.Encrypt()

	//then
	assert.Nil(t, err)
	assert.NotEmptyf(t, c.GetEncryptRaw(), "Encrypt must return a encrypt with success")
}

func TestEncrypt_DecryptData(t *testing.T) {
	// given
	p := generatePerson()
	data, _ := json.Marshal(p)
	key := "person"
	c := NewCrypto(key, data)
	_ = c.Encrypt()

	t.Run("Success encrypt and decrypt person", func(t *testing.T) {
		// when
		err := c.Decrypt()
		var decrypt *entity.Person
		err = json.Unmarshal(c.GetEncryptRaw(), &decrypt)

		// then
		assert.Equal(t, p, decrypt)
		assert.Nil(t, err)
	})

	t.Run("Failure to decrypt person", func(t *testing.T) {
		// given
		c.Key = "invalid"

		//when
		err := c.Decrypt()

		//then
		assert.Error(t, err)
	})
}

func generatePerson() *entity.Person {
	return &entity.Person{
		Key:       uuid.New().String(),
		FirstName: "Joaquim",
		LastName:  "Barbosa",
		BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 1, time.UTC),
		Address: entity.Address{
			City:  "São Paulo",
			State: "SP",
		},
	}
}
