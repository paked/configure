package configeur

import (
	"fmt"
	"testing"
)

func value(v interface{}, err error) interface{} {
	return v
}

func test(received, expected interface{}, failFunc func()) {
	if received == expected {
		return
	}

	fmt.Printf("Failed test. Got '%[1]v' (type %[1]T), expected '%[2]v' (type %[2]T)\n", received, expected)

	failFunc()
}

func normalFailFunc(t *testing.T) func() {
	return func() {
		t.Fail()
	}
}
