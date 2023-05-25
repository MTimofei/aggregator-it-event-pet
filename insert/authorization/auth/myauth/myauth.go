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

// регистрация пользователя
func (a *MyAuth) Reg() (err error) {
	defer func() { err = e.IfErr("cen't reg", err) }()

	err = a.validName()
	if err != nil {
		return err
	}

	//получение хэша пароля и динамической соли
	a.ECod.GenerationSalt()
	h := a.ECod.Hesh(a.Password)

	//отпровляем в стор
	err = a.Stor.Add(storage.NweClient(a.Login, h.Password, h.Salt))
	if err != nil {
		return err
	}

	return nil
}

func (a *MyAuth) Auth() (user *auth.User, err error) {
	defer func() { err = e.IfErr("cen't auth", err) }()
	//получение данных пользователя

	//проверяем пороль

	return nil, nil
}

// проверка уникальности имени, е
// сли функия возврощает нулевую ошибку
// то имя доступно
func (a *MyAuth) validName() (err error) {
	defer func() { err = e.IfErr("validName", err) }()
	result, err := a.Stor.Login(a.Login)
	if err != nil {
		return err
	}
	if result != nil {
		return fmt.Errorf(errUsernameUsed)
	}

	return nil
}
