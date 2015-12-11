package configure

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/hcl"
)

// NewHCL returns an instance of the HCL checker. It takes a function
// which returns an io.Reader which will be called when the first value
// is recalled. The contents of the io.Reader MUST be decodable into HCL or JSON.
func NewHCL(gen func() (io.Reader, error)) *HCL {
	return &HCL{
		gen: gen,
	}
}

// NewHCLFromFile returns an instance of the HCL checker. It reads its
// data from a file (file content can be HCL or JSON) which its location has been specified through the path
// parameter
func NewHCLFromFile(path string) *HCL {
	return NewHCL(func() (io.Reader, error) {
		return os.Open(path)
	})
}

// HCL represents the HCL Checker. It reads an io.Reader and then pulls a value out of a map[string]interface{}.
type HCL struct {
	values map[string]interface{}
	gen    func() (io.Reader, error)
}

// Setup initializes the HCL Checker
func (h *HCL) Setup() error {
	r, err := h.gen()
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	s := buf.String()

	//first parse the hcl file
	obj, err := hcl.Parse(s)
	if err != nil {
		return err
	}

	h.values = make(map[string]interface{})

	// then decode the object
	if err = hcl.DecodeObject(&h.values, obj); err != nil {
		return err
	}

	return nil
}

func (h *HCL) value(name string) (interface{}, error) {
	val, ok := h.values[name]
	if !ok {
		return nil, errors.New("variable does not exist")
	}

	return val, nil
}

// Int returns an int if it exists within the HCL io.Reader
func (h *HCL) Int(name string) (int, error) {
	v, err := h.value(name)
	if err != nil {
		return 0, err
	}

	f, ok := v.(float64)
	if !ok {
		i, ok := v.(int)
		if !ok {
			return v.(int), errors.New(fmt.Sprintf("%T unable", v))
		}

		return i, nil
	}

	return int(f), nil
}

// Bool returns a bool if it exists within the HCL io.Reader.
func (h *HCL) Bool(name string) (bool, error) {
	v, err := h.value(name)
	if err != nil {
		return false, err
	}

	return v.(bool), nil
}

// String returns a string if it exists within the HCL io.Reader.
func (h *HCL) String(name string) (string, error) {
	v, err := h.value(name)
	if err != nil {
		return "", err
	}

	return v.(string), nil
}
