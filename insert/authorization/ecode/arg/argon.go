package arg

import (
	"errors"
	"new/insert/authorization/config"
	"new/insert/authorization/ecode"
	"new/pkg/gen"

	"golang.org/x/crypto/argon2"
)

const (
	ErrInValidSalt = "invalid salt"
)

// структура реализует интерфейс ECode
type Hash struct {
	StaticSalt  []byte
	DynamicSalt []byte
}

// создает структуру для новаого пользователя(регистрации)
func New() *Hash {
	return &Hash{StaticSalt: []byte(config.StaticSalt), DynamicSalt: gen.Rand8bytes()}
}

// создает структуру c salt, len(salt)==8 (используется для идентификации)
// встучи не правильной длины слайса возврощает ошибку
func Create(salt []byte) (*Hash, error) {

	switch {
	case len(salt) != 8:
		return nil, errors.New(ErrInValidSalt)
	default:
		return &Hash{StaticSalt: []byte(config.StaticSalt), DynamicSalt: salt}, nil
	}

}

// извекает хэшь пользователя с использованием argon2
func (object *Hash) Hesh(password string) (h *ecode.Hash) {
	saticPass := argon2.IDKey([]byte(password), object.StaticSalt, 4, 16*1024, 4, 128)
	pass := argon2.IDKey(saticPass, object.DynamicSalt, 4, 16*1024, 4, 128)

	return ecode.New(pass, object.DynamicSalt)
}
