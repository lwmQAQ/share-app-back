package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
密码加密工具类
*/
func SecretPassword(password string) (string, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func Verify(hashedPassword string, inputPassword string) error {
	// Compare the input password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}
