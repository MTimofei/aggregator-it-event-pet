package token

type Token interface {
	Create(*User) (string, error)
}

type User struct {
	Id    int64
	Login string
	Role  string
}

func NewUser(id int64, login string, role string) (u *User) {
	return &User{Id: id, Login: login, Role: role}
}
