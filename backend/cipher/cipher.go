package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var secretKey = "iVLIEHoynVmZj5WVHkqR2B6qhxmrO3nv"

type Cipher struct {
	block cipher.Block
}

func InitCipher() (*Cipher, error) {
	newCipher, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	var cip = &Cipher{
		block: newCipher,
	}
	return cip, nil
}

func (c *Cipher) Encrypt(plainText []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(c.block)
	if err != nil {
		return nil, err
	}

	// number used once
	nonce := make([]byte, gcm.NonceSize())
	
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	encoded := base64.StdEncoding.EncodeToString(cipherText)
	return []byte(encoded), nil
}

func (c *Cipher) Decrypt(cipherText []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(c.block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("cipherText is short")
	}

	nonce, encrypted := cipherText[:nonceSize], cipherText[nonceSize:]
	
	plainText, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
