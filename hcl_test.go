package configure

import (
	"io"
	"strings"
	"testing"
)

func TestHCL(t *testing.T) {
	src := `s = "hello"`

	hcl := NewHCL(func() (io.Reader, error) {
		return strings.NewReader(src), nil
	})

	hcl.Setup()

	if v, err := hcl.String("s"); v != "hello" {
		t.Errorf("hello %v 'hello' %v", v, err)
	}

	if _, err := hcl.String("this-message-does-not-exist"); err == nil {
		t.Error("hello2")
	}
}
