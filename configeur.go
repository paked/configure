package configeur

// Checker is the interface which external adders must comply to
type Checker interface {
	Int(name string, stack []Checker) int       // Retrieves an int
	String(name string, stack []Checker) string // Retrieves a string
}

type Configeur struct {
	strings map[string]*string // The requested string values
	ints    map[string]*int    // The requested integer values
	stack   []Checker          // A list of all the "middlewear" which is used to find a value
}

// New creates a new configeur instance, immediately ready to be used.
// It takes a slice of Checker interfaces which will be used to retrieve values.
func New(stack []Checker) *Configeur {
	c := &Configeur{
		strings: make(map[string]*string),
		ints:    make(map[string]*int),
		stack:   stack,
	}

	return c
}

// Classic returns the default set of Checkers: Flag, JSON and Environment
func Classic(extras ...Checker) []Checker {
	checkers := []Checker{Flag{}, JSON{}, Environment{}}
	checkers = append(checkers, extras...)

	return checkers
}
