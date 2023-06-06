package mytoken

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"new/insert/authorization/config"
	"new/insert/authorization/token"
	"new/pkg/e"
	"new/pkg/help"
	"strings"

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
	defer func() { err = e.IfErr("can't create signed", err) }()
	help.Up(&config.KeyId, 1)

	jwttoken := jwt.New(jwt.SigningMethodES256)

	key, err := x509.MarshalECPrivateKey(k.key)
	if err != nil {
		return "", err
	}

	jwttoken.Header["name"] = "acc"
	jwttoken.Header["jti"] = config.KeyId
	jwttoken.Header["key"] = key
	jwttoken.Claims = jwt.MapClaims{
		"id":   u.Id,
		"name": u.Login,
		"role": u.Role,
	}

	token, err = jwttoken.SignedString(k.key)
	if err != nil {
		return "", err
	}
	//fmt.Println(jwttoken.Method.Alg())
	return token, nil
}

func (k *Key) Verifation(token string) (u *token.User, err error) {
	defer func() { err = e.IfErr("cen't verifation token", err) }()

	var key *ecdsa.PrivateKey

	head := struct {
		Name string `json:"name"`
		Id   int    `json:"jti"`
		Key  []byte `json:"key"`
	}{}

	header, err := base64.RawURLEncoding.DecodeString(strings.Split(token, ".")[0])
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(header, &head)
	if err != nil {
		return nil, err
	}

	key, err = x509.ParseECPrivateKey(head.Key)
	if err != nil {
		return nil, err
	}

	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return &key.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}

	return u, nil
}
