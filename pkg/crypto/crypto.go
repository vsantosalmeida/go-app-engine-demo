package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"strings"
)

type Crypto struct {
	Key  string `json:"key"`
	Data []byte `json:"raw"`
}

func NewCrypto(key string, data []byte) dataCrypt {
	return &Crypto{
		Key:  key,
		Data: data,
	}
}

func (c *Crypto) createHash() []byte {
	hasher := md5.New()
	hasher.Write([]byte(c.Key))
	return hasher.Sum(nil)
}

func (c *Crypto) Encrypt() error {
	block, err := aes.NewCipher(c.createHash())
	if err != nil {
		log.Printf("Couldn't create an aes block %q", err)
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Couldn't create a cipher gcm %q", err)
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Printf("Failed do copy rand to nonce %q", err)
		return err
	}
	c.Data = gcm.Seal(nonce, nonce, c.Data, nil)
	return nil
}

func (c *Crypto) Decrypt() error {
	block, err := aes.NewCipher(c.createHash())
	if err != nil {
		log.Printf("Couldn't create an aes block %q", err)
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Couldn't create a cipher gcm %q", err)
		return err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := c.Data[:nonceSize], c.Data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Printf("Couldn't decrypt the data %q", err)
		return err
	}
	c.Data = plaintext
	return nil
}

func (c *Crypto) GetEncryptRaw() string {
	var b strings.Builder
	for _, n := range c.Data {
		fmt.Fprintf(&b, "%d,", n)
	}
	s := b.String()
	s = s[:b.Len()-1]

	return "[" + s + "]"
}

func (c *Crypto) GetDecryptRaw() []byte {
	return c.Data
}
