package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrypto_Encrypt(t *testing.T) {
	d := NewCrypto("xpto", []byte(raw))
	svc := NewService(d)

	t.Run("Test success encrypt", func(t *testing.T) {
		err := svc.Encrypt()
		assert.Nil(t, err)
	})

	t.Run("Test sucess decrypt", func(t *testing.T) {
		err := svc.Decrypt()
		r := string(svc.GetDecryptRaw())
		assert.Nil(t, err)
		assert.Equal(t, raw, r)
	})

	t.Run("Test failure to decrypt", func(t *testing.T) {
		d = NewCrypto("abc", []byte(raw))
		svc = NewService(d)
		err := svc.Decrypt()
		assert.Error(t, err)
	})
}
