package util

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func MustGeneratePassword(password string) string {
	hash, err := GeneratePassword(password)
	if err != nil {
		panic(err)
	}

	return hash
}

func ConfirmPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err != nil
}
