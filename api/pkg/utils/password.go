package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func GenerateSalt(size int) (string, error) {
	if size <= 0 {
		size = 16
	}
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func HashPassword(password, salt string) (string, error) {
	if password == "" || salt == "" {
		return "", errors.New("password or salt is empty")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePassword 校验密码是否正确
func ComparePassword(hashedPwd, password, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password+salt))
	return err == nil
}

func IsStrongPassword(password string) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",\.<>\/\?\\|]`).MatchString(password)
	return hasNumber && hasLower && hasUpper && hasSpecial
}
