package util

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct{}

func (b Bcrypt) Hash(password string, saltRounds int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (b Bcrypt) Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
