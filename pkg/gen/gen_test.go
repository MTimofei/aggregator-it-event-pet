package gen_test

import (
	"new/pkg/gen"
	"reflect"
	"testing"
)

// тест функции Rand8bytes
func TestRand8bytes(t *testing.T) {
	result := gen.Rand8bytes()
	if reflect.DeepEqual(result, []byte{}) && len(result) != 8 {
		t.Errorf("TestRand8bytes")
	}
}
