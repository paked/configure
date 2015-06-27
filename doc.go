// Package configeur is an easy to use multi-layer configuration system.
//
// configeur makes use of Checkers, which are used to retrieve values from their respective data sources.
// There are three built in Checkers, Environment, Flag and JSON. Environment retrieves environment variables.
// Flag retrieves variables within the flags of a command. JSON retrieves values from a JSON file/blob. Checkers
// can be essentially thought of as "configuration middlewear", in fact parts of the package API was inspired by
// negroni (https://github.com/codegangsta/negroni, special thanks to codegangsta for the awesome package!) and the
// standard libraries flag package.
//
// It is very easy to create your own Checkers, all they have to do is satisfy the Checker interface.
// That is a, Int method, String method and a Bool method. To retrieve their respective data types.
// If you do create your own Checkers I would be more than happy to add a link to them somewhere in the repository.
//
// A standard use-case can be found in the example folder.
package configeur
