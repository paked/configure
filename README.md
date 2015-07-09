# configeur
Configeur is a Go package that gives you easy configuration of your project through redundancy. It has an API inspired by [negroni](http://godoc.org/github.com/codegangsta/negroni) and the [flag](http://godoc.org/flag) package.

## What is it?
Configeur aims to be the `github.com/codegangsta/negroni` of configuration. It is a `Checker` manager, in the same way `negroni` managers `net/http` middlewear. A `Checker` is a way of retrieving configuration values from a source, these can be easily made by completing the [Checker interface](http://godoc.org/github.com/paked/configeur#Checker). The idea is that you as a developer provide Configeur with a selection of `Checker`'s, either built in or not and it will iterate over them attempting to find values defined by the developer. If a `Checker` is successful in its retrieval, then Configeur will stop the iteration for that value. If it is not then Configeur will attempt the next `Checker` in chronological order.

# Getting Started
After you have installed Go (and have made sure to correctly setup your [GOPATH](http://golang.org/doc/code.html#GOPATH)) create a new `.go` file, maybe `hello.go`.
```go
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
  fmt.Printf("Hello, %v", *name)
}
```
If you run this code with
```
go run hello.go
```
`Hello, Harrison` will be printed to your terminal, not that interesting right... read on!

### Stage One : Declaration
```go
var (
  conf = configeur.New()
  name = conf.String("name", "Harrison", "The name you want to greet")
)
```
The declaration stage is important because it defines exactly what you CAN configure! First, `conf` is created which is your key to adding Checkers and retrieving variables. After that you begin properly declaring your variables, in this example only a string is declared but in practice you can use any number of `String`'s, `Int`'s or `Bool`'s. The variables returned by these methods are pointers to their respective types.

### Stage Two : Configuration
```go
func init() {
  conf.Use(configeur.NewEnvironment())
  conf.Use(configeur.NewFlag())
}
```
The configuration stage is where you configure `configeur` by adding Checkers to the stack. Checkers are objects which will attempt to retrieve your variables from their respective data sources. When a `Checker` fails the next one in the stack is called, the stack is in the same order that the `Checker`'s were added in. You can configure `configeur` anytime before you call the `conf.Parse()` function, but the `init()` function provides a reliable place to do so.

### Stage Three : Usage
```go
func main() {
  conf.Parse()
  fmt.Printf("Hello, %v", *name)
}
```
The final stage is where you can actually use the variables you have declared. After using `conf.Parse()` your variables should then be populated and accesible by dereferencing it (`name`).

## Execution
If you were to run this code in its current state it would print `Hello, Harrison` because `Harrison` is the default value provided in the declaration stage. But if you provide `--name=Johny` when you execute the command it will print `Hello, Johny`. At this point `configeur` is behaving like the default `flag` package through the `Flag` Checker. Now, run `export NAME=Jarvis` in your command line and execute the program again and ommit the entire `--name=`command line flag. You will see a `Hello, Jarvis`, as `configeur` has fallen back upon the `Environment` Checker. Note that, if you provide both means of input the environment variable will be used, as it has higher priority as it was added before the `Flag` Checker in the configuration stage. This works with any number of Checkers from any source, as long as the fulfil [the `Checker` interface](http://godoc.org/github.com/paked/configeur#Checker).

## Further Reading
[More package documentation can be found on godoc](http://godoc.org/pkg/github.com/paked/configeur).

A more complicated example can be found in the [example folder](http://github.com/paked/configeur/blob/master/example/), it uses all three variable types (`Int`, `Bool` and `String`) and all three of the default `Checker`'s (`JSON`, `Environment` and `Flag`).

# Contributing
If you notice something that you feel is broken or missing in configeur feel free to open up an issue so that we can discuss it more. While small changes could be immediately put into a PR, I believe it saves everyones time to discuss major changes before implementing them. Contributions are welcome and appreciated.

# Checkers
| Name | Location | Initialiser |Description|
|---|---|---|---|
|Environment| ***[builtin]*** http://github.com/paked/configeur |`NewEnvironment()` | Environment checks the os environment variables for values |
|JSON|***[builtin]*** http://github.com/paked/configeur |`NewJSON(io.Reader)`| Retrieves values from an `io.Reader` containing JSON |
|Flag|***[builtin]*** http://github.com/paked/configeur |`NewFlag()`| Retrieve flagged values from `os.Args` in a `--x=y` format|

If you write your own Checker I would *LOVE* to see it, create a PR with a new entry in the table!

# Note
As you may have noticed, I am not *particularly* great at english. If you notice a way to de-garble a few of my sentences be sure to let me know... Not only I, but future readers will be greatful too :)

###### designed and implemented by [Harrison Shoebridge](http://github.com/paked)
