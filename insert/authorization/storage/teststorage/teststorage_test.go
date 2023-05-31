package teststorage_test

import (
	"errors"
	"new/insert/authorization/storage"
	"new/insert/authorization/storage/teststorage"
	"reflect"
	"testing"
	"time"
)

// тест подключения
func TestConnect(t *testing.T) {
	//создание заведомо верных данных
	expected := map[int64]storage.User{
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
	testCases := []struct {
		name     string
		data     storage.NewUser
		expected map[int64]storage.User
	}{
		{
			"test1",
			storage.NewUser{Login: "test3", Salt: []byte("test3salt"), Hesh: []byte("test3Password"), Roly: "client"},
			map[int64]storage.User{
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
		{
			"test2",
			storage.NewUser{Login: "test4", Salt: []byte("test4salt"), Hesh: []byte("test4Password"), Roly: "client"},
			map[int64]storage.User{
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

	for _, test := range testCases {

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

// тест изменения данных бд
func TestUpdata(t *testing.T) {
	//создание заведомо верных данных
	testCases := []struct {
		name        string
		data        storage.User
		expected    map[int64]storage.User
		errExpected error
	}{
		{
			"test1",
			storage.User{ID: 1,
				Login: "NewTest1",
				Salt:  []byte("Test1Salt"),
				Hash:  []byte("NewTest1Password"),
				Roly:  "admin",
				RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			map[int64]storage.User{
				1: {
					ID:    1,
					Login: "NewTest1",
					Salt:  []byte("Test1Salt"),
					Hash:  []byte("NewTest1Password"),
					Roly:  "admin",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
				2: {
					ID:    2,
					Login: "test2",
					Salt:  []byte("Test2Salt"),
					Hash:  []byte("test2Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
			nil,
		},
		{
			"test2",
			storage.User{ID: 3,
				Login: "test3",
				Salt:  []byte(""),
				Hash:  []byte(""),
				Roly:  "client",
				RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			map[int64]storage.User{
				1: {
					ID:    1,
					Login: "NewTest1",
					Salt:  []byte("Test1Salt"),
					Hash:  []byte("NewTest1Password"),
					Roly:  "admin",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
				2: {
					ID:    2,
					Login: "test2",
					Salt:  []byte("Test2Salt"),
					Hash:  []byte("test2Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
			errors.New("there is no record this id"),
		},
	}

	db, _ := teststorage.Connect()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			err := db.Update(&test.data)
			// проверка резултата
			if !reflect.DeepEqual(db.DB, test.expected) && err == test.errExpected {
				t.Errorf("\nОжидалось %v\nполучено %v", test.expected, db.DB)
			}
		})
	}
}

// тест метода удаления
func TestRemoval(t *testing.T) {
	testCases := []struct {
		name        string
		id          int64
		expected    map[int64]storage.User
		errExpected error
	}{
		{
			"test1",
			1,
			map[int64]storage.User{
				2: {
					ID:    2,
					Login: "test2",
					Salt:  []byte("Test2Salt"),
					Hash:  []byte("test2Password"),
					Roly:  "client",
					RegAt: time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
			nil,
		},
		{
			"test2",
			2,
			map[int64]storage.User{},
			nil,
		},
	}
	db, err := teststorage.Connect()
	if err != nil {
		t.Errorf("err testRemoval: %e", err)
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			err := db.Removal(tC.id)
			if !reflect.DeepEqual(db.DB, tC.expected) && err != tC.errExpected {
				t.Errorf("\nОжидалось: %v\nполучено: %v\nerr: %e", tC.expected, db.DB, err)
			}
		})
	}
}

// тест поискак данных по логину
func TestLogin(t *testing.T) {
	testCases := []struct {
		name        string
		id          int64
		expected    map[int64]storage.User
		errExpected error
	}{}

	db, err := teststorage.Connect()

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {

			if !reflect.DeepEqual(db.DB, tC.expected) && err != tC.errExpected {
				t.Errorf("\nОжидалось: %v\nполучено: %v\nerr: %e", tC.expected, db.DB, err)
			}
		})
	}
}
