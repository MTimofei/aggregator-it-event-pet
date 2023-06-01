package arg

import (
	"new/insert/authorization/config"
	"new/insert/authorization/ecode"
	"new/pkg/gen"

	"golang.org/x/crypto/argon2"
)

// структура реализует интерфейс ECode
type Hash struct {
	staticSalt  []byte
	dynamicSalt []byte
}

// создает структуру для новаого пользователя(регистрации)
func New() *Hash {
	return &Hash{staticSalt: []byte(config.StaticSalt), dynamicSalt: gen.Rand8bytes()}
}

// создает структуру c salt (используется для идентификации)
func Create(salt []byte) *Hash {
	return &Hash{staticSalt: []byte(config.StaticSalt), dynamicSalt: salt}
}

// извекает хэшь пользователя с использованием argon2
func (object *Hash) Hesh(password string) (h *ecode.Hash) {
	saticPass := argon2.IDKey([]byte(password), object.staticSalt, 4, 16*1024, 4, 128)
	pass := argon2.IDKey(saticPass, object.dynamicSalt, 4, 16*1024, 4, 128)

	return ecode.New(pass, object.dynamicSalt)
}
