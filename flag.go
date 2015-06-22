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

// Flag is an ULTRA simple flag parser for strings and ints.
// it expects flags in the format --x=2 (where x is the var name
// and 2 is the value)
type Flag struct {
	args []string
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

func (f Flag) String(name string) (string, error) {
	return f.value(name)
}
