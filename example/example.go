package main

import (
	"fmt"

	"github.com/paked/configeur"
)

var (
	// set up a configeur instance with no stack
	conf = configeur.New()
	// declare flags / things to configure
	x = conf.Int("x", 0, "how many times you want the string repeated!")
	s = conf.String("s", "default bro", "a selected string")
	n = conf.Bool("n", true, "whether or not you want new lines")
)

func init() {
	// add configuration middlewears to the stack
	conf.Use(configeur.NewFlag())
	conf.Use(configeur.NewJSONFromFile("config.json"))
	conf.Use(configeur.NewEnvironment())
}

func main() {
	// populate the pointers
	conf.Parse()

	fmt.Printf("printing string `%v` %v time(s)\n", *s, *x)

	for i := 0; i < *x; i++ {
		fmt.Print(*s)

		if *n {
			fmt.Print("\n")
		}
	}
}
