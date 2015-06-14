package configeur

type JSON struct {
}

func (j JSON) String(name string, stack []Checker) string {
	return ""
}

func (j JSON) Int(name string, stack []Checker) int {
	return 0
}
