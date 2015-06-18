package configeur

import (
	"os"
	"testing"
)

var e = NewEnvironment()

func TestProcess(t *testing.T) {
	if e.process("hello-world") != "HELLO_WORLD" {
		t.Fail()
	}
}

func TestEnvironment(t *testing.T) {
	os.Setenv("TESTING_TESTING", "xyz")

	if v, _ := e.String("testing-testing"); v != "xyz" {
		t.Fail()
	}

	if _, err := e.String("this-var-should-not-exist"); err == nil {
		t.Fail()
	}
}
