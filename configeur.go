package configeur

import "fmt"

// Checker is the interface which external adders must comply to
type Checker interface {
	Int(name string) (int, error)       // Get the integer value from the source
	String(name string) (string, error) // Get a string value from a source
}

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

// Use adds any amount of new checkers. Useful for when you want to configure
// a Checker before adding it to the stack.
func (c *Configeur) Use(checkers ...Checker) {
	c.stack = append(c.stack, checkers...)
}

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

// New creates a new configeur instance, immediately ready to be used.
// It takes a slice of Checker interfaces which will be used to retrieve values.
func New(stack ...Checker) *Configeur {
	c := &Configeur{
		strings: make(map[string]*stringOption),
		ints:    make(map[string]*intOption),
		stack:   stack,
	}

	return c
}

// Classic returns the default set of Checkers: Environemnt
func Classic(extras ...Checker) []Checker {
	checkers := []Checker{Environment{}}
	checkers = append(checkers, extras...)

	return checkers
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
