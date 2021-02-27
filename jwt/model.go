package jwt

import "github.com/dgrijalva/jwt-go"

//UserChaim token中加密的信息
type UserChaim struct {
	UID int64
	Uname string
	jwt.StandardClaims
}
