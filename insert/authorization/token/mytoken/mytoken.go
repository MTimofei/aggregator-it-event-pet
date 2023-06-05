package mytoken

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
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
	defer func() { err = e.IfErr("can't create signed", err) }()
	help.Up(&config.KeyId, 1)

	b := bytes.NewBuffer([]byte{})
	_, err = fmt.Fscan(b, k.key.PublicKey)
	if err != nil {
		return "", err
	}

	jwttoken := jwt.New(jwt.SigningMethodES256)

	jwttoken.Header["name"] = "acc"
	jwttoken.Header["jti"] = config.KeyId
	jwttoken.Header["key"] = b.Bytes()

	jwttoken.Claims = jwt.MapClaims{
		"id":   u.Id,
		"name": u.Login,
		"role": u.Role,
	}

	token, err = jwttoken.SignedString(k.key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (k *Key) Verifation(token string) (u *token.User, err error) {
	defer func() { err = e.IfErr("cen't verifation token", err) }()

	// //head := make(map[string]interface{}, 3)

	// head := struct {
	// 	Name string `json:"name"`
	// 	Id   int    `json:"jti"`
	// 	Key  []byte `json:"key"`
	// }{}
	// header, err := base64.RawURLEncoding.DecodeString(strings.Split(token, ".")[0])
	// if err != nil {
	// 	return nil, err
	// }

	// err = json.Unmarshal(header, &head)
	// if err != nil {
	// 	return nil, err
	// }

	// //k.key.PublicKey
	// //fmt.Println(k.key.PublicKey)
	// fmt.Println(head.Key)
	// t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
	// 	return &head.Key, nil
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(t)
	return u, nil
}
