package teststorage_test

import (
	"new/insert/authorization/storage"
	"new/insert/authorization/storage/teststorage"
	"reflect"
	"testing"
	"time"
)

// тест подключения
func TestConnect(t *testing.T) {
	//создание заведомо верных данных
	expected := map[int]storage.User{
		1: {ID: 1,
			Login: "test1",
			Salt:  []byte("Test1Salt"),
			Hash:  []byte("test1Password"),
			Roly:  "admin",
			RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
		2: {
			ID:    2,
			Login: "test2",
			Salt:  []byte("Test2Salt"),
			Hash:  []byte("test2Password"),
			Roly:  "client",
			RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	t.Run("connect", func(t *testing.T) {
		result, _ := teststorage.Connect()

		// проверка резултата
		if !reflect.DeepEqual(result.DB, expected) {
			t.Errorf("\nОжидалось %v\nполучено %v", expected, result.DB)
		}
	})
}

// тест метода добавления
func TestAdd(t *testing.T) {
	//создаем тест кес
	Cases := []struct {
		name     string
		data     storage.NewUser
		expected []storage.User
	}{
		{"test1",
			storage.NewUser{Login: "test3", Salt: []byte("test3salt"), Hesh: []byte("test3Password"), Roly: "client"},
			[]storage.User{
				1: {ID: 1,
					Login: "test1",
					Salt:  []byte("Test1Salt"),
					Hash:  []byte("test1Password"),
					Roly:  "admin",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
				2: {
					ID:    2,
					Login: "test2",
					Salt:  []byte("Test2Salt"),
					Hash:  []byte("test2Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
				3: {
					ID:    3,
					Login: "test3",
					Salt:  []byte("test3salt"),
					Hash:  []byte("test3Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
		},
		{"test2",
			storage.NewUser{Login: "test4", Salt: []byte("test4salt"), Hesh: []byte("test4Password"), Roly: "client"},
			[]storage.User{
				1: {ID: 1,
					Login: "test1",
					Salt:  []byte("Test1Salt"),
					Hash:  []byte("test1Password"),
					Roly:  "admin",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
				2: {
					ID:    2,
					Login: "test2",
					Salt:  []byte("Test2Salt"),
					Hash:  []byte("test2Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
				3: {
					ID:    3,
					Login: "test3",
					Salt:  []byte("test3salt"),
					Hash:  []byte("test3Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
				4: {
					ID:    4,
					Login: "test4",
					Salt:  []byte("test4salt"),
					Hash:  []byte("test4Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
		},
	}

	db, _ := teststorage.Connect()

	for _, test := range Cases {

		t.Run(test.name, func(t *testing.T) {
			//вызов функции
			err := db.Add(&test.data)
			if err != nil {
				t.Errorf("err %e", err)
			}

			// проверка резултата
			if !reflect.DeepEqual(db.DB, test.expected) {
				t.Errorf("\nОжидалось %v\nполучено %v", test.expected, db.DB)
			}

		})
	}
}
