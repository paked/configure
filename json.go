package configeur

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// NewJSON returns a new instance of the JSON Checker. It takes an io.Reader in it's
// parameters, which will be used to read the JSON content. The JSON must be marshallable
// into a map[string]string.
func NewJSON(r io.Reader) *JSON {
	return &JSON{
		values: make(map[string]string),
		read:   false,
		file:   r,
	}
}

// JSON represents the JSON Checker. It reads an io.Reader and then pulls a value out of a map[string]string.
type JSON struct {
	values map[string]string
	read   bool
	file   io.Reader
}

func (j *JSON) value(name string) (string, error) {
	if !j.read {
		dec := json.NewDecoder(j.file)
		j.values = make(map[string]string)

		err := dec.Decode(&j.values)
		if err != nil {
			return "", err
		}

		j.read = true
	}

	val, ok := j.values[name]
	if !ok {
		return "", errors.New("that variable does not exist")
	}

	return val, nil
}

// Int returns the integer if it exists within the unmarshalled JSON io.Reader.
func (j *JSON) Int(name string) (int, error) {
	v, err := j.value(name)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// String returns the integer if it exists within the unmarshalled JSON io.Reader.
func (j *JSON) String(name string) (string, error) {
	return j.value(name)
}
