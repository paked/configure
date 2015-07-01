package main

import (
	"fmt"

	"github.com/paked/configeur"
)

var (
	conf = configeur.New()
	name = conf.String("name", "Harrison", "The name you want to greet")
)

func init() {
	conf.Use(configeur.NewEnvironment())
	conf.Use(configeur.NewFlag())
}

func main() {
	conf.Parse()
	fmt.Printf("Hello, %v\n", *name)
}
