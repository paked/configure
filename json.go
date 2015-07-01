package configeur

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

// NewJSON returns an instance of the JSON checker. It takes a function
// which returns an io.Reader which will be called when the first value
// is recalled. The contents of the io.Reader MUST be decodable into JSON.
func NewJSON(gen func() (io.Reader, error)) *JSON {
	return &JSON{
		gen: gen,
	}
}

// NewJSONFromFile returns an instance of the JSON checker. It reads its
// data from a file which its location has been specified through the path
// parameter
func NewJSONFromFile(path string) *JSON {
	return NewJSON(func() (io.Reader, error) {
		return os.Open(path)
	})
}

// JSON represents the JSON Checker. It reads an io.Reader and then pulls a value out of a map[string]string.
type JSON struct {
	values map[string]interface{}
	gen    func() (io.Reader, error)
}

func (j *JSON) value(name string) (interface{}, error) {
	if j.values == nil {
		r, err := j.gen()
		if err != nil {
			return nil, err
		}

		dec := json.NewDecoder(r)
		j.values = make(map[string]interface{})

		err = dec.Decode(&j.values)
		if err != nil {
			return "", err
		}

	}

	val, ok := j.values[name]
	if !ok {
		return nil, errors.New("that variable does not exist")
	}

	return val, nil
}

// Int returns an int if it exists within the unmarshalled JSON io.Reader.
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

// Bool returns a bool if it exists within the unmarshalled JSON io.Reader.
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

// String returns a string if it exists within the unmarshalled JSON io.Reader.
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
