package ecod

type ECod interface {

	//генерирует соль
	GenerationSalt() //(salt []byte)
	Heah(password string) (h *Hash)
}

type Hash struct {
	Password []byte
	Salt     []byte
}
