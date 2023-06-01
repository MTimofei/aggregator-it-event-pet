package gen

import (
	"crypto/rand"
)

// генерирует слайс на 8 рандомных байт
func Rand8bytes() []byte {
	// cоздаем байтовый срез заданной длины
	length := 8
	randomBytes := make([]byte, length)

	// читаем случайные данные в байтовый срез
	_, err := rand.Read(randomBytes)
	if err != nil {
		return []byte{}
	}

	return randomBytes
}
