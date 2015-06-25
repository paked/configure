package configeur

import "fmt"

// Checker is the interface which external "checkers" must comply to. If a
// retrieval fails the next Checker in the Configeur.stack stack will be called. Checker's are added to the stack through the configeur.Use method.
type Checker interface {
	// Int finds and retrieves an integer by name. If it finds the value it will
	// return it, and if it doesn't an error will be returned and the next Checker
	// in the stack will be called.
	Int(name string) (int, error)

	// String finds and retrieves an string by name. If it finds the value it will
	// return it, and if it doesn't an error will be returned and the next Checker
	// in the stack will be called.
	String(name string) (string, error)

	// Bool finds and retrieves a boolean value by name.
	Bool(name string) (bool, error)
}

// Configeur is a stack of Checkers which are used to retrieve configuration values. It aims
// to have a similar API as the flag package in the standard library. Checker's are evaluated
// in the same order they are added through the initalization and Configeur.Use functions.
type Configeur struct {
	options map[string]*option // A map of all of the provided options
	stack   []Checker          // A list of all the "middlewear" which is used to find a value
}

// Int declares an int value. A pointer is returned that will
// point to the value after c.Parse() has been called.
func (c *Configeur) Int(name string, def int, description string) *int {
	v := c.option(name, def, description, intType)
	i, ok := v.(*int)

	if !ok {
		fmt.Printf("could not retrieve pointer to value for name %v\n", name)
		return nil
	}

	return i
}

// String declares a string value. A pointer is returned that will
// point to the value after c.Parse() has been called.
func (c *Configeur) String(name, def, description string) *string {
	v := c.option(name, def, description, stringType)
	s, ok := v.(*string)

	if !ok {
		fmt.Printf("could not retrieve pointer to value for name %v\n", name)
		return nil
	}

	return s
}

// Bool declares a new boolean option. A pointer is returned to this
// value that will be filled after Configeur.Parse() has been called.
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
		var s string
		opt.value = &s
	case intType:
		var i int
		opt.value = &i
	case boolType:
		var b bool
		opt.value = &b
	}

	return opt.value
}

// Use adds a variable amounts of new Checkers to the stack.
func (c *Configeur) Use(checkers ...Checker) {
	c.stack = append(c.stack, checkers...)
}

// Parse populates the pointers returned through the Configeur.Int and Configeur.String
// methods.
func (c *Configeur) Parse() {
	for _, opt := range c.options {
		changed := false
		for _, checker := range c.stack {
			// TODO undupe
			switch opt.typ {
			case stringType:
				s, err := checker.String(opt.name)
				if err != nil {
					fmt.Println(err)
					continue
				}

				opt.set(s)
			case intType:
				i, err := checker.Int(opt.name)
				if err != nil {
					fmt.Println(err)
					continue
				}

				opt.set(i)
			case boolType:
				b, err := checker.Bool(opt.name)
				if err != nil {
					fmt.Println(err)
					continue
				}

				opt.set(b)
			}

			changed = true
			break
		}

		if !changed {
			opt.set(opt.def)

			opt.value = opt.def
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
			fmt.Println("that var doesnt fit thy value")
		}

		*z = value.(bool)
	case int:
		z, ok := o.value.(*int)
		if !ok {
			fmt.Println("that var doesnt fit thy value")
		}

		*z = value.(int)

	case string:
		z, ok := o.value.(*string)
		if !ok {
			fmt.Println("that var doesnt fit thy value")
		}

		*z = value.(string)
	}
}
