package utils

import "golang.org/x/crypto/bcrypt"

//生成密码
func PasswordNew(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword = string(hash)
	return
}

//比对密码
func PasswordValidate(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
