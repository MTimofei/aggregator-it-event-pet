package myauth

import (
	"fmt"
	"new/insert/authorization/auth"
	"new/insert/authorization/ecod"
	"new/insert/authorization/storage"
	"new/pkg/e"
)

const (

	//юзер с таким именем уже существует
	errUsernameUsed = "username used"
)

type MyAuth struct {
	Login    string
	Password string
	Stor     storage.Storage
	ECod     ecod.ECod
}

func New(login string, password string) (a *MyAuth) {
	return &MyAuth{Login: login, Password: password}
}

func (a *MyAuth) Reg() (err error) {
	// проверка наличия юзера
	result, err := a.Stor.Login(a.Login)
	if err != nil {
		return e.Err("cen't reg", err)
	}
	if result != nil {
		return e.Err("cen't reg", fmt.Errorf(errUsernameUsed))
	}

	//получение хэша пароля и динамической соли
	a.ECod.GenerationSalt()
	h := a.ECod.Heah(a.Password)

	//отпровляем в стор
	a.Stor.Add(&storage.NewUser{Login: a.Login,Hash: h.Password,Salt: h.Salt,Roly: ""})

	return e.Err("cen't reg", nil)
}

func (a *MyAuth) Auth() (user *auth.User, err error) {
	return
}
