package configeur

type Flag struct {
}

func (f Flag) String(name string, stack []Checker) string {
	return ""
}

func (f Flag) Int(name string, stack []Checker) int {
	return 0
}
