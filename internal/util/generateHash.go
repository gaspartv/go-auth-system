package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomHash() (string, error) {
	var result string
	for i := 0; i < 6; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result += string(charset[index.Int64()])
	}
	return strings.ToUpper(result), nil
}
