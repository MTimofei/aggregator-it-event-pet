package ecode

// интерфэйс для хэширования пароля
type ECode interface {

	//возврощает структуру с хеш
	Hesh(password string) (h *Hash)
}

type Hash struct {

	//хеш пороля
	Password []byte

	//хэш динамической соли
	Salt []byte
}

func New(password, salt []byte) *Hash {
	return &Hash{Password: password, Salt: salt}
}
