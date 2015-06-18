package configeur

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func NewEnvironment() *Environment {
	return &Environment{}
}

type Environment struct {
}

// value takes a string in normal flag syntax (hello-world) and changes
// it into a environment variable syntax (HELLO_WORLD). It returns the
// the value associated with that env.
func (e Environment) value(name string) (string, error) {
	n := e.process(name)
	v := os.Getenv(n)
	if v == "" {
		return v, errors.New("Value does not exist")
	}

	return v, nil
}

func (e Environment) process(name string) string {
	// name is in form hello-world
	// we need it in form HELLO_WORLD
	name = strings.ToUpper(name)
	name = strings.Replace(name, "-", "_", -1)

	return name
}

func (e Environment) Int(name string) (int, error) {
	v, err := e.value(name)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (e Environment) String(name string) (string, error) {
	return e.value(name)
}
