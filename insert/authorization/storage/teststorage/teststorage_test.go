package teststorage_test

import (
	"new/insert/authorization/storage"
	"new/insert/authorization/storage/teststorage"
	"testing"
)

func TestAdd(t *testing.T) {
	//создаем тестовые данные
	type Case struct {
		name       string
		payload    *storage.NewUser
		experience storage.NewUser
	}
	var cases = []Case{
		{
			name:       "invalid",
			payload:    &storage.NewUser{Login: "test1", Salt: []byte("sqlt1"), Hesh: []byte("password2"), Roly: "admin"},
			experience: storage.NewUser{Login: "test1", Salt: []byte("sqlt1"), Hesh: []byte("password2"), Roly: "admin"},
		},
		{
			name:       "invalid",
			payload:    &storage.NewUser{Login: "test2", Salt: []byte("sqlt2"), Hesh: []byte("password1"), Roly: "client"},
			experience: storage.NewUser{Login: "test2", Salt: []byte("sqlt2"), Hesh: []byte("password1"), Roly: "client"},
		},
	}
	// создаем конект
	tdb, err := teststorage.Connect()
	if err != nil {
		t.Errorf("Unexpected error occurred: %v", err)
	}

	for i := range cases {
		// вызов функции Add
		err = tdb.Add(cases[i].payload)
		if err != nil {
			t.Errorf("Unexpected error occurred: %v", err)
		}
	}

}
