package mytoken

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"new/insert/authorization/config"
	"new/insert/authorization/token"
	"new/pkg/e"
	"new/pkg/help"

	"github.com/golang-jwt/jwt/v5"
)

type Key struct {
	key *ecdsa.PrivateKey
}

func New() (k *Key, err error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, e.Err("can't generatekey", err)
	}

	k = &Key{key: key}
	return k, nil
}

func (k *Key) Create(u *token.User) (token string, err error) {
	help.Up(&config.KeyId, 1)

	jwttoken := jwt.New(jwt.SigningMethodES256)

	jwttoken.Header["name"] = "acc"
	jwttoken.Header["jti"] = config.KeyId
	jwttoken.Header["key"] = k.key.PublicKey

	jwttoken.Claims = jwt.MapClaims{
		"id":   u.Id,
		"name": u.Login,
		"role": u.Role,
	}

	token, err = jwttoken.SignedString(k.key)
	if err != nil {
		return "", e.Err("can't create signed", err)
	}

	return token, nil
}
