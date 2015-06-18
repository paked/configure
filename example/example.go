package main

import (
	"fmt"
	"os"

	"github.com/paked/configeur"
)

var (
	// set up a configeur instance with no stack
	conf = configeur.New()
	// declare flags / things to configure
	x = conf.Int("x", 0, "how many times you want the string repeated!")
	s = conf.String("s", "default bro", "a selected string")
)

func init() {
	config, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	// add configuration middlewears to the stack
	conf.Use(configeur.NewJSON(config))
	conf.Use(configeur.NewEnvironment())
}

func main() {
	conf.Parse()

	fmt.Printf("printing string `%v` %v time(s)\n", *s, *x)

	for i := 0; i < *x; i++ {
		fmt.Println(*s)
	}
}
