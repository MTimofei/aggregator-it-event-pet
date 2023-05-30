package teststorage

import (
	"new/insert/authorization/storage"
	"time"
)

// структура реализующая бд для тестов
type DataBase struct {
	DB map[int]storage.User
}

// подключение к тестовой бд
func Connect() (tdb *DataBase, err error) {
	tdb = &DataBase{
		DB: make(map[int]storage.User),
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
	var maxId int = 0
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

func (tdb *DataBase) Update(user *storage.User) (err error) {
	return nil
}

func (tdb *DataBase) Removal(user *storage.User) (err error) {
	return nil
}

func (tdb *DataBase) Login(login string) (user *storage.User, err error) {
	return user, nil
}

func (tdb *DataBase) All() (users []storage.User, err error) {
	return users, nil
}
