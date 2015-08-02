package configure

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// NewEnvironment creates a new instance of the Environment Checker.
func NewEnvironment() *Environment {
	return &Environment{}
}

// Environment is a configeur.Checker. It retrieves values from the host OS's environment variables.
// In this process, it will take a flag-like name, and convert it to a environmnet-like name. The process
// for this is to change all characters to upper case, and then replace hyphons with underscores.
type Environment struct {
}

func (e Environment) Setup() error {
	return nil
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

// Int returns an int if it exists in the set environment variables.
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

// Bool returns a bool if it exists in the set environment variables.
func (e *Environment) Bool(name string) (bool, error) {
	v, err := e.value(name)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(v)
}

// String returns a string if it exists in the set environment variables.
func (e Environment) String(name string) (string, error) {
	return e.value(name)
}
