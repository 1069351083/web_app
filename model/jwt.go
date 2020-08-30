package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Id        int
	LoginName string
	Name      string
	jwt.StandardClaims
}
