package arg_test

import (
	"new/insert/authorization/ecode"
	"new/insert/authorization/ecode/arg"
	"reflect"
	"testing"
)

var _ ecode.ECode = (*arg.Hash)(nil)

func TestNew(t *testing.T) {
	result := arg.New()

	if reflect.DeepEqual(*result, arg.Hash{}) {
		t.Error("mathed New not valid")
	}
}
