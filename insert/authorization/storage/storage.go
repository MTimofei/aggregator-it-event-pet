package storage

import "time"

type Storage interface {

	//добавляем в бд данные юзера при регистрации
	Add(user *NewUser) (err error)

	//изменение данных юзера в бд
	Update(user *User) (err error)

	//удоление юзера из бд
	Removal(user *User) (err error)

	//получение данных юзера по логину
	Login(login string) (user *User, err error)

	//получение всех юзеров
	All() (user []User, err error)
}

type User struct {
	ID        int64
	Login     string
	Salt      []byte
	Hash      []byte
	Roly      string
	RegAt     time.Time
	UpdLastAt time.Time
}

type NewUser struct {
	Login string
	Salt  []byte
	Hesh  []byte
	Roly  string
}

func NweClient(login string, hesh []byte, salt []byte) *NewUser {
	var nUser = NewUser{Login: login, Hesh: hesh, Salt: salt, Roly: "client"}
	return &nUser
}
