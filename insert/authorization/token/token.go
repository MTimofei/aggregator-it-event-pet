package token

type Token interface {

	//создает токен
	Create(u *User) (token string, err error)

	// проверяет токен и возврощает структуру c данными юзра
	Verifation(token string) (u *User, err error)
}

type User struct {
	Id    int64
	Login string
	Role  string
}

func NewUser(id int64, login string, role string) (u *User) {
	return &User{Id: id, Login: login, Role: role}
}
