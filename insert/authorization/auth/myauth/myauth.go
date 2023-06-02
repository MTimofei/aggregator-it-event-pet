package myauth

import (
	"fmt"
	"new/insert/authorization/auth"
	"new/insert/authorization/ecode"
	"new/insert/authorization/storage"
	"new/pkg/e"
)

const (

	//юзер с таким именем уже существует
	errUsernameUsed = "username used"
)

type MyAuth struct {
	login    string
	password string
	stor     storage.Storage
	eCod     ecode.ECode
}

func New(login string, password string) (a *MyAuth) {
	return &MyAuth{login: login, password: password}
}

// регистрация пользователя
func (a *MyAuth) Reg() (err error) {
	defer func() { err = e.IfErr("cen't reg", err) }()

	err = a.validName()
	if err != nil {
		return err
	}

	//получение хэша пароля и динамической соли
	h := a.eCod.Hesh(a.password)

	//отпровляем в стор
	err = a.stor.Add(storage.NweClient(a.login, h.Password, h.Salt))
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
	result, err := a.stor.Login(a.login)
	if err != nil {
		return err
	}
	if result != nil {
		return fmt.Errorf(errUsernameUsed)
	}

	return nil
}
