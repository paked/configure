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
}

// Configeur is a stack of Checkers which are used to retrieve configuration values. It aims
// to have a similar API as the flag package in the standard library. Checker's are evaluated
// in the same order they are added through the initalization and Configeur.Use functions.
type Configeur struct {
	strings map[string]*stringOption // The requested string values
	ints    map[string]*intOption    // The requested integer values
	stack   []Checker                // A list of all the "middlewear" which is used to find a value
}

// Int declares an int value. A pointer is returned that will
// point to the value after c.Parse() has been called.
func (c *Configeur) Int(name string, def int, description string) *int {
	opt := &intOption{
		name:        name,
		def:         def,
		description: description,
	}

	c.ints[name] = opt

	return &opt.value
}

// String declares a string value. A pointer is returned that will
// point to the value after c.Parse() has been called.
func (c *Configeur) String(name, def, description string) *string {
	opt := &stringOption{
		name:        name,
		def:         def,
		description: description,
	}

	c.strings[name] = opt

	return &opt.value
}

// Use adds a variable amounts of new Checkers to the stack.
func (c *Configeur) Use(checkers ...Checker) {
	c.stack = append(c.stack, checkers...)
}

// Parse populates the pointers returned through the Configeur.Int and Configeur.String
// methods.
func (c *Configeur) Parse() {
	for _, opt := range c.strings {
		changed := false
		for _, checker := range c.stack {
			v, err := checker.String(opt.name)
			if err != nil {
				continue
			}

			changed = true
			opt.value = v
			break
		}

		if !changed {
			fmt.Printf("set %v to default (%v)", opt, opt.def)
			opt.value = opt.def
		}
	}

	// so much duplicated code?
	for _, opt := range c.ints {
		changed := false
		for _, checker := range c.stack {
			v, err := checker.Int(opt.name)
			if err != nil {
				fmt.Println(err)
				continue
			}

			opt.value = v
			changed = true
			break
		}

		if !changed {
			fmt.Printf("set %v to default (%v)\n", opt.name, opt.def)
			opt.value = opt.def
		}
	}
}

// New returns a pointer to a new Configeur instance with a stack
// provided through the variadic stack variable.
func New(stack ...Checker) *Configeur {
	c := &Configeur{
		strings: make(map[string]*stringOption),
		ints:    make(map[string]*intOption),
		stack:   stack,
	}

	return c
}

type intOption struct {
	name        string
	description string
	def         int
	value       int
}

func (io intOption) String() string {
	return fmt.Sprintf("%v(%v)", io.name, io.description)
}

type stringOption struct {
	name        string
	description string
	def         string
	value       string
}

func (so stringOption) String() string {
	return fmt.Sprintf("[%v(%v)]", so.name, so.description)
}
