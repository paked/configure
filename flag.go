package configeur

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// NewFlag returns a new instance of the Flag Checker, using os.Args as its
// flag source.
func NewFlag() *Flag {
	f := &Flag{
		args: os.Args,
	}

	return f
}

// Flag is an ULTRA simple flag parser for strings, ints and bools.
// It expects flags in the format --x=2 (where x is the var name
// and 2 is the value)
type Flag struct {
	args []string
}

func (f Flag) Setup() error {
	for _, arg := range f.args {
		if arg == "-h" {
			fmt.Println("PRINTING DEFAULTS HERE")
		}
	}

	return nil
}

func (f Flag) value(name string) (string, error) {
	for _, arg := range f.args {
		// --x=2 -> x=2
		ass := strings.TrimPrefix(arg, "--")
		if ass == arg {
			continue
		}

		// x=2 -> 2
		val := strings.TrimPrefix(ass, fmt.Sprintf("%v=", name))
		if val == ass {
			continue
		}

		return val, nil
	}

	return "", errors.New("could not find that value")
}

// Int returns an int if it exists in the set flags.
func (f Flag) Int(name string) (int, error) {
	v, err := f.value(name)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// Bool returns a bool if it exists in the set flags.
func (f *Flag) Bool(name string) (bool, error) {
	v, err := f.value(name)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(v)
}

// String returns a string if it exists in the set flags.
func (f Flag) String(name string) (string, error) {
	return f.value(name)
}
