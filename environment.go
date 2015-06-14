package configeur

type Environment struct {
}

func (e Environment) String(name string, stack []Checker) string {
	return ""
}

func (e Environment) Int(name string, stack []Checker) int {
	return 0
}
