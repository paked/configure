package configeur

import (
	"fmt"
	"testing"
)

func TestFlags(t *testing.T) {
	f := Flag{
		args: []string{"exec", "--x=2", "--y=four five six seven"},
	}

	if v, _ := f.String("x"); v != "2" {
		fmt.Println(v)
		t.Fail()
	}

	if v, _ := f.String("y"); v != "four five six seven" {
		fmt.Println(v)
		t.Fail()
	}
}
