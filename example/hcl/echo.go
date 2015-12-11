package main

import (
	"fmt"

	"github.com/paked/configure"
)

var (
	// set up a configure instance with no default stack
	conf = configure.New()
	// declare flags / things to configure
	amount   = conf.Int("amount", 0, "how many times you want the string repeated!")
	message  = conf.String("message", "Echo!", "a selected string")
	newlines = conf.Bool("newlines", true, "whether or not you want new lines")
)

func init() {
	// add configuration middlewears to the stack
	conf.Use(configure.NewHCLFromFile("echo.hcl"))
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

func usage() string {
	return "Echo is the best message echo-er available in your terminal!\nUse the amount flag to set how many times you to echo\nmessage for what you want to echo\nand newlines for whether you want breaks in between messages\n"
}
