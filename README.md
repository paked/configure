# configeur
configeur is an easy to use multi-layer configuration system.

# Contributing
If you notice something that you feel is broken or missing in configeur feel free to open up an issue so that we can discuss it more. While small changes could be immediately put into a PR, I believe it saves everyones time to discuss major changes before implementing them.

# Checkers
| Name | Location | Initialiser |Description|
|---|---|---|---|
|Environment| ***[builtin]*** http://github.com/paked/configeur |`NewEnvironment()` | Environment checks the os environment variables for values |
|JSON|***[builtin]*** http://github.com/paked/configeur |`NewJSON(io.Reader)`| Retrieves values from an `io.Reader` containing JSON |
|Flag|***[builtin]*** http://github.com/paked/configeur |`NewFlag()`| Retrieve flagged values from `os.Args` in a `--x=y` format|

If you write your own Checker I would *LOVE* to see it, create a PR with a new entry in the table!

# Usage
[Usage documentation can be found on godoc](http://godoc.org/pkg/github.com/paked/configeur).

An example can be found in the [example folder](http://github.com/paked/configeur/blob/master/example/).

###### designed by [Harrison Shoebridge](http://github.com/paked)
