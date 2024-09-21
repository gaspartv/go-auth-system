package util

import (
	"crypto/sha256"
	"encoding/hex"
)

type Crypto struct{}

func (c Crypto) Hash(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
