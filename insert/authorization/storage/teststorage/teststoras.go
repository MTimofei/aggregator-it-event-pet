package teststorage

import (
	"encoding/json"
	"new/insert/authorization/storage"
	"new/pkg/e"
	"os"
	"time"
)

// структура реализующая бд для тестов
type DataBase struct {
	DB *os.File
}

// подключение к тестовой бд
// после вызова бд незабудте закрыть отложеной функцией defer tdb.DB.Close()
func Connect() (tdb *DataBase, err error) {
	file, err := os.Create("db.txt")
	if err != nil {
		return nil, e.Err("cen't connect test db", err)
	}

	return &DataBase{DB: file}, nil
}

func (tdb *DataBase) Add(user *storage.NewUser) (err error) {
	defer func() { err = e.IfErr("cen't add test db", err) }()

	bs, err := json.Marshal(storage.User{
		ID:        2,
		Login:     user.Login,
		Salt:      user.Salt,
		Hash:      user.Hesh,
		Roly:      user.Roly,
		RegAt:     time.Now(),
		UpdLastAt: time.Now(),
	})
	if err != nil {
		return err
	}

	_, err = tdb.DB.Write(bs)
	if err != nil {
		return err
	}

	return nil
}

func (tdb *DataBase) Update(user *storage.User) (err error) {
	return
}

func (tdb *DataBase) Removal(user *storage.User) (err error) {
	return
}

func (tdb *DataBase) Login(login string) (user *storage.User, err error) {
	return
}

func (tdb *DataBase) All() (user []storage.User, err error) {
	return
}
