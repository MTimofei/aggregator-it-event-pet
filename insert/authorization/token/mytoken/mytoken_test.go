package mytoken_test

import (
	"new/insert/authorization/token"
	"new/insert/authorization/token/mytoken"
	"reflect"
	"testing"
)

var _ token.Token = (*mytoken.Key)(nil)

func TestNew(t *testing.T) {
	var err error
	testCases := []struct {
		name        string
		errExpected error
	}{
		{
			"test1",
			nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err = mytoken.New()
			if err != tC.errExpected {
				t.Errorf("ivalid func mytoken.New: %v", err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	var err error
	testCases := []struct {
		name        string
		data        *token.User
		errExpected error
	}{
		{
			"test1",
			token.NewUser(1, "Test1Login", "admin"),
			nil,
		},
		{
			"test2",
			token.NewUser(1, "Test2Login", "client"),
			nil,
		},
	}

	key, err := mytoken.New()
	if err != nil {
		t.Error(err)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := key.Create(tC.data)
			if err != tC.errExpected {
				t.Errorf("invalid method mytoken.Create: %v", err)
			}
		})
	}
}

func TestVerifation(t *testing.T) {
	var err error
	var tkn string
	var key *mytoken.Key
	var user *token.User

	testCases := []struct {
		name        string
		data        *token.User
		expected    *token.User
		errExpected error
	}{
		{
			"test1",
			token.NewUser(1, "Test1Login", "client"),
			token.NewUser(1, "Test1Login", "client"),
			nil,
		},
		{
			"test2",
			token.NewUser(2, "Test2Login", "client"),
			token.NewUser(2, "Test2Login", "client"),
			nil,
		},
	}

	key, err = mytoken.New()
	if err != nil {
		t.Error(err)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tkn, _ = key.Create(tC.data)
			user, err = key.Verifation(tkn)
			if !reflect.DeepEqual(user, tC.expected) && err != tC.errExpected {
				t.Errorf("invalid method mytoken.Verifation:\n\texpected: %v\n\tresult: %v\n\terr: %v\n\terr expected:%v", tC.expected, user, err, tC.errExpected)
			}
		})
	}
}
