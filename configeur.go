package configeur

import (
	"fmt"
	"os"
)

// Checker is the interface which external checkers must comply to. In the configeur
// implementation, if one Checker fails the next one in the stack should be called.
type Checker interface {
	// Setup initializes the Checker
	Setup() error
	// Int attempts to get an int value from the data source.
	Int(name string) (int, error)

	// String attempts to get a string value from the data source.
	String(name string) (string, error)

	// Bool attempts to get a bool value from the data source.
	Bool(name string) (bool, error)
}

// Configeur is a stack of Checkers which are used to retrieve configuration values. It has a
// similar API as the flag package in the standard library, and is also partially
// inspired by negroni. Checkers are evaluated in the same order they are added.
type Configeur struct {
	options map[string]*option // A map of all of the provided options
	stack   []Checker          // A list of all the "middlewear" which is used to find a value
}

// Int defines an int flag with a name, default and description. The return value
// is a pointer which will be populated with the value of the flag.
func (c *Configeur) Int(name string, def int, description string) *int {
	v := c.option(name, def, description, intType)
	i, ok := v.(*int)

	if !ok {
		fmt.Printf("could not retrieve pointer to value for name %v\n", name)
		return nil
	}

	return i
}

// String defines a string flag with a name, default and description. The return value
// is a pointer which will be populated with the value of the flag.
func (c *Configeur) String(name, def, description string) *string {
	v := c.option(name, def, description, stringType)
	s, ok := v.(*string)

	if !ok {
		fmt.Printf("could not retrieve pointer to value for name %v\n", name)
		return nil
	}

	return s
}

// Bool defines a bool flag with a name, default and description. The return value
// is a pointer which will be populated with the value of the flag.
func (c *Configeur) Bool(name string, def bool, description string) *bool {
	v := c.option(name, def, description, boolType)
	b, ok := v.(*bool)

	if !ok {
		fmt.Printf("could not retrieve pointer to that value\n")
		return nil
	}

	return b
}

// option returns a pointer of a type specified through the valueType parameter.
//
// note: You could potentially find the value to be pointed to through the def
// parameter, but this would pose an issue when one is not provided.
func (c *Configeur) option(name string, def interface{}, description string, typ valueType) interface{} {
	opt := &option{
		name:        name,
		def:         def,
		description: description,
		typ:         typ,
	}

	c.options[name] = opt

	switch opt.typ {
	case stringType:
		opt.value = new(string)
	case intType:
		opt.value = new(int)
	case boolType:
		opt.value = new(bool)
	}

	return opt.value
}

// Use adds a variable amount of Checkers onto the stack.
func (c *Configeur) Use(checkers ...Checker) {
	c.stack = append(c.stack, checkers...)
}

func (c *Configeur) setup() {
	for _, checker := range c.stack {
		err := checker.Setup()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Forced to abort because %T is failing to setup: %v\n", checker, err)
			os.Exit(1)
		}
	}
}

// Parse populates all of the defined arguments with their values provided by
// the stacks Checkers.
func (c *Configeur) Parse() {
	c.setup()
	for _, opt := range c.options {
		changed := false
		for _, checker := range c.stack {
			switch opt.typ {
			case stringType:
				s, err := checker.String(opt.name)
				if err != nil {
					continue
				}

				opt.set(s)
			case intType:
				i, err := checker.Int(opt.name)
				if err != nil {
					continue
				}

				opt.set(i)
			case boolType:
				b, err := checker.Bool(opt.name)
				if err != nil {
					continue
				}

				opt.set(b)
			}

			changed = true
			break
		}

		if !changed {
			opt.set(opt.def)
		}
	}
}

// New returns a pointer to a new Configeur instance with a stack
// provided through the variadic stack variable.
func New(stack ...Checker) *Configeur {
	c := &Configeur{
		options: make(map[string]*option),
		stack:   stack,
	}

	return c
}

type valueType int

func (vt valueType) String() string {
	name := ""
	switch vt {
	case intType:
		name = "Integer"
	case stringType:
		name = "String"
	case boolType:
		name = "Boolean"
	default:
		return "NOT A valueType"
	}

	return name
}

const (
	intType valueType = iota + 1
	stringType
	boolType
)

type option struct {
	name        string
	description string
	def         interface{}
	value       interface{}
	typ         valueType
}

func (o option) String() string {
	return fmt.Sprintf("%v(%v)[%v]", o.name, o.description, o.typ)
}

func (o *option) set(value interface{}) {
	switch value.(type) {
	case bool:
		z, ok := o.value.(*bool)
		if !ok {
			return
		}

		*z = value.(bool)
	case int:
		z, ok := o.value.(*int)
		if !ok {
			return
		}

		*z = value.(int)

	case string:
		z, ok := o.value.(*string)
		if !ok {
			return
		}

		*z = value.(string)
	}
}
