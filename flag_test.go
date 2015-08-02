package configure

import (
	"testing"
)

func TestFlagStrings(t *testing.T) {
	f, ff := createFlag(t, "--x=hello", `--z=hello world`, "--y=22")

	test(value(f.String("x")), "hello", ff)
	test(value(f.String("z")), "hello world", ff)
	test(value(f.String("y")), "22", ff)
}

func TestFlagsInts(t *testing.T) {
	f, ff := createFlag(t, "--x=2", "--z=-1")

	test(value(f.Int("x")), 2, ff)
	test(value(f.Int("z")), -1, ff)
}

func TestFlagsBools(t *testing.T) {
	f, ff := createFlag(t, "--x=T", "--z=F")

	test(value(f.Bool("x")), true, ff)
	test(value(f.Bool("z")), false, ff)
}

func createFlag(t *testing.T, args ...string) (Flag, func()) {
	f := Flag{
		args: append([]string{"executable"}, args...),
	}

	ff := normalFailFunc(t)

	return f, ff
}
