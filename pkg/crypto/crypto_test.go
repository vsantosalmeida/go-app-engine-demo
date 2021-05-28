package crypto

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const raw = `{"name": "test","last": "xpto"}`

func TestEncryptData(t *testing.T) {
	//given
	data, _ := json.Marshal(raw)
	key := "person"
	c := NewCrypto(key, data)

	//when
	err := c.Encrypt()

	//then
	assert.Nil(t, err)
	assert.NotEmptyf(t, c.GetEncryptRaw(), "encrypt must return a encrypt with success")
}

func TestEncrypt_DecryptData(t *testing.T) {
	// given
	key := "person"
	c := NewCrypto(key, []byte(raw))
	_ = c.Encrypt()

	t.Run("Success encrypt and decrypt person", func(t *testing.T) {
		// when
		err := c.Decrypt()

		// then
		assert.Nil(t, err)
		assert.Equal(t, raw, string(c.GetDecryptRaw()))
	})
}
