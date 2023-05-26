package ecode

type ECode interface {

	//генерирует соль
	GenerationSalt() //(salt []byte)

	//возврощает структуру с хеш
	Hesh(password string) (h *Hash)
}

type Hash struct {

	//хеш пороля
	Password []byte

	//хэш динамической соли
	Salt []byte
}
