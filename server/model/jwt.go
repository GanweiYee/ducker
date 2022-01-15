package model

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       string
	Username string
	Account  string
}
