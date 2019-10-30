package crypto

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TestHash(t *testing.T)  {
	password := "123456"
	hash, _ := PasswordHash(password)

	fmt.Println("密码:", password)
	fmt.Println("hash:", hash)

	hash = "$2y$10$fUnmGkhKObjzU84nQvupaeEnPuqdxvJx6AZhW6u33tWVl9eiZAbSG"
	match := PasswordVerify(password, hash)
	fmt.Println("验证:", match)
}