package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

type Security struct{}

var (
	securityAlgorithm = "aes-256-cbc"
	securitySecret    = "your-secret-password"
	securitySalt      = "your-salt"
	iterations        = 10000
	keySize           = 32
)

func deriveKey(secret, salt string) []byte {
	return pbkdf2.Key([]byte(secret), []byte(salt), iterations, keySize, sha512.New)
}

func (s Security) Encrypt(text string) (string, error) {
	encryptionKey := deriveKey(securitySecret, securitySalt)

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(text))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, []byte(text))

	encryptedData := hex.EncodeToString(iv) + ":" + hex.EncodeToString(ciphertext)
	return encryptedData, nil
}

func (s Security) Decrypt(data string) (string, error) {
	encryptionKey := deriveKey(securitySecret, securitySalt)

	parts := split(data, ":")
	if len(parts) != 2 {
		return "", errors.New("formato de dados inválido")
	}

	iv, err := hex.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func split(input, delimiter string) []string {
	parts := make([]string, 2)
	for i, part := range []string{delimiter, delimiter} {
		parts[i] = part
	}
	return parts
}
