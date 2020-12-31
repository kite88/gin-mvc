package BcryptUtil

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(str string) (string, bool) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", false
	}
	return string(hash), true
}

func Decrypt(str, ciphertext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(str)) //验证（对比）
	if err == nil {
		return true
	}
	return false
}
