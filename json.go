package configeur

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// NewJSON returns a new instance of the JSON Checker. It takes an io.Reader in it's
// parameters, which will be used to read the JSON content. The JSON must be marshallable
// into a map[string]string.
func NewJSON(r io.Reader) *JSON {
	return &JSON{
		values: make(map[string]interface{}),
		read:   false,
		file:   r,
	}
}

// JSON represents the JSON Checker. It reads an io.Reader and then pulls a value out of a map[string]string.
type JSON struct {
	values map[string]interface{}
	read   bool
	file   io.Reader
}

func (j *JSON) value(name string) (interface{}, error) {
	if !j.read {
		dec := json.NewDecoder(j.file)
		j.values = make(map[string]interface{})

		err := dec.Decode(&j.values)
		if err != nil {
			return "", err
		}

		j.read = true
	}

	val, ok := j.values[name]
	if !ok {
		return nil, errors.New("that variable does not exist")
	}

	return val, nil
}

// Int returns the integer if it exists within the unmarshalled JSON io.Reader.
func (j *JSON) Int(name string) (int, error) {
	v, err := j.value(name)
	if err != nil {
		return 0, err
	}

	f, ok := v.(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%T unable", v))
	}

	return int(f), nil
}

func (j *JSON) Bool(name string) (bool, error) {
	v, err := j.value(name)
	if err != nil {
		return false, err
	}

	b, ok := v.(bool)
	if !ok {
		return false, errors.New("unable to cast")
	}

	return b, nil
}

// String returns the integer if it exists within the unmarshalled JSON io.Reader.
func (j *JSON) String(name string) (string, error) {
	v, err := j.value(name)
	if err != nil {
		return "", err
	}

	s, ok := v.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("unable to cast %T", v))
	}

	return s, nil
}
