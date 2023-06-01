package teststorage

import (
	"new/insert/authorization/storage"
	"new/pkg/e"
	"reflect"
	"time"
)

const (
	errUpdata = "there is no record this id"
	errLogin  = "record not faund"
)

// структура реализующая бд для тестов
type DataBase struct {
	DB map[int64]storage.User
}

// подключение к тестовой бд
func Connect() (tdb *DataBase, err error) {
	tdb = &DataBase{
		DB: make(map[int64]storage.User),
	}

	//данные по умолчания для дб
	tdb.DB[1] = storage.User{
		ID:    1,
		Login: "test1",
		Salt:  []byte("Test1Salt"),
		Hash:  []byte("test1Password"),
		Roly:  "admin",
		RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	tdb.DB[2] = storage.User{
		ID:    2,
		Login: "test2",
		Salt:  []byte("Test2Salt"),
		Hash:  []byte("test2Password"),
		Roly:  "client",
		RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	return tdb, nil
}

// добавление в db
func (tdb *DataBase) Add(user *storage.NewUser) (err error) {
	var maxId int64 = 0
	for i := range tdb.DB {
		if maxId < i {
			maxId = i
		}
	}

	maxId++
	tdb.DB[maxId] = storage.User{
		ID:    int64(maxId),
		Login: user.Login,
		Salt:  user.Salt,
		Hash:  user.Hesh,
		Roly:  user.Roly,
		RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	return nil
}

// изменени даных записи по ее id
func (tdb *DataBase) Update(user *storage.User) (err error) {
	var c storage.User

	u := tdb.DB[user.ID]

	if reflect.DeepEqual(u, c) {
		return e.Err(errUpdata, nil)
	}
	u.Login = user.Login
	u.Salt = user.Salt
	u.Hash = user.Hash
	u.Roly = user.Roly
	tdb.DB[user.ID] = u

	return nil
}

// удаление записи из дб по id
func (tdb *DataBase) Removal(id int64) (err error) {
	delete(tdb.DB, id)
	return nil
}

// поиск записи по логину
func (tdb *DataBase) Login(login string) (user *storage.User, err error) {
	for i := range tdb.DB {
		if tdb.DB[i].Login == login {
			u := tdb.DB[i]
			return &u, nil
		}
	}
	return &storage.User{}, e.Err(errLogin, nil)
}

// полученние всех записей из бд
func (tdb *DataBase) All() (users []storage.User, err error) {
	for _, u := range tdb.DB {
		users = append(users, u)
	}
	return users, nil
}
