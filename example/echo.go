package main

import (
	"fmt"

	"github.com/paked/configeur"
)

var (
	// set up a configeur instance with no stack
	conf = configeur.New()
	// declare flags / things to configure
	amount   = conf.Int("amount", 0, "how many times you want the string repeated!")
	message  = conf.String("message", "Echo!", "a selected string")
	newlines = conf.Bool("newlines", true, "whether or not you want new lines")
)

func init() {
	// add configuration middlewears to the stack
	conf.Use(configeur.NewFlag())
	conf.Use(configeur.NewJSONFromFile("echo.json"))
	conf.Use(configeur.NewEnvironment())
}

func main() {
	// populate the pointers
	conf.Parse()

	for i := 0; i < *amount; i++ {
		fmt.Print(*message)

		if *newlines {
			fmt.Print("\n")
		}
	}
}
