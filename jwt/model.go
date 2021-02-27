package jwt

import "github.com/dgrijalva/jwt-go"

type UserChaim struct {
	UID int64
	Uname string
	jwt.StandardClaims
}
