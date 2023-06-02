package arg_test

import (
	"errors"
	"new/insert/authorization/config"
	"new/insert/authorization/ecode"
	"new/insert/authorization/ecode/arg"
	"reflect"
	"testing"
)

var _ ecode.ECode = (*arg.Hash)(nil)

// тест
func TestNew(t *testing.T) {
	result := arg.New()

	if reflect.DeepEqual(*result, arg.Hash{}) {
		t.Error("mathed New not valid")
	}
}

func TestCreate(t *testing.T) {
	var result *arg.Hash
	var err error

	testCases := []struct {
		nameTest    string
		salt        []byte
		expected    *arg.Hash
		errExpected error
	}{
		{
			"test1",
			[]byte("tryu1758"),
			&arg.Hash{[]byte(config.StaticSalt), []byte("tryu1758")},
			nil,
		},
		{
			"test2",
			[]byte("tryu175"),
			nil,
			errors.New(arg.ErrInValidSalt),
		},
		{
			"test3",
			[]byte("tryu17589"),
			nil,
			errors.New(arg.ErrInValidSalt),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.nameTest, func(t *testing.T) {
			result, err = arg.Create(tC.salt)

			if !reflect.DeepEqual(err, tC.errExpected) {
				t.Errorf("\nerrExpected:\n%v\nerr:\n%v\n", tC.errExpected, err)
				//return
			}
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("\nexpected:\n%v\nin:\n%v\n", tC.expected, result)
			}
		})
	}
}

func TestHash(t *testing.T) {
	testCases := []struct {
		nameTest string
		password string
		salt     []byte

		errExpected error
	}{
		{
			"test1",
			"password1",
			[]byte("saltTest"),
			nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.nameTest, func(t *testing.T) {
			e, err := arg.Create(tC.salt)
			if err != nil {
				t.Error(err)
			}
			result1 := e.Hesh(tC.password)
			result2 := e.Hesh(tC.password)
			if !reflect.DeepEqual(result1, result2) {
				t.Errorf("\nout1:\n%v\nout1:\n%v\n", result1, result2)
			}
		})
	}
}
