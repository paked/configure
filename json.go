package configeur

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

func NewJSON(r io.Reader) *JSON {
	return &JSON{
		values: make(map[string]string),
		read:   false,
		file:   r,
	}
}

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

func (j *JSON) String(name string) (string, error) {
	return j.value(name)
}
