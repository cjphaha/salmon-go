package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

type RSAVerify interface {
	GenToken(jwtBody UserChaim) (token string, err error)
	ParseToken(token string) (data *UserChaim, err error)
}

type rsaVerify struct {
	pubPath, priPath string
}

func NewRSA(pubPath, priPath string) RSAVerify {
	return &rsaVerify{
		pubPath: pubPath,
		priPath: priPath,
	}
}

func (this *rsaVerify) getPubKey() (pubKey *rsa.PublicKey, err error) {
	// 公钥
	pubKeyBytes, err := ioutil.ReadFile(this.pubPath)
	if err != nil {
		log.Fatal("公钥文件读取失败")
	}
	pubKey, err = jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		log.Fatal("公钥文件不正确")
	}
	return
}

func (this *rsaVerify) getPriKey() (priKey *rsa.PrivateKey, err error) {
	// 私钥
	priKeyBytes, err := ioutil.ReadFile(this.priPath)
	if err != nil {
		log.Fatal("私钥文件读取失败")
	}
	priKey, err = jwt.ParseRSAPrivateKeyFromPEM(priKeyBytes)
	if err != nil {
		log.Fatal("私钥文件不正确")
	}
	return
}

func (this *rsaVerify) GenToken(jwtBody UserChaim) (token string, err error){
	priKey, err := this.getPriKey()
	if err != nil {
		return
	}
	token_obj := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtBody)
	token, err = token_obj.SignedString(priKey)
	if err != nil {
		return
	}
	return
}

func (this *rsaVerify) ParseToken(token string) (data *UserChaim, err error) {
	pubKey, err := this.getPubKey()
	if err != nil {
		return
	}
	data = &UserChaim{}
	getToken, err := jwt.ParseWithClaims(token, data, func (token *jwt.Token) (i interface{}, err error) {
		return pubKey, nil
	})
	if getToken != nil && getToken.Valid {
		fmt.Println(getToken.Claims.(*UserChaim).Uname)
		fmt.Println(getToken.Claims.(*UserChaim).ExpiresAt)
	} else if v, ok := err.(*jwt.ValidationError); ok {
		if v.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New(JWT_ERROR_TOKEN)
		} else if v.Errors&(jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New(JWT_TOKEN_EXPIRED)
		} else {
			return nil, errors.New(JWT_CAN_NOT_HANDLE_THIS_TOKEN + err.Error())
		}
	} else {
		return nil, errors.New(JWT_CAN_NOT_DECODE_TOEKN)
	}
	return
}
