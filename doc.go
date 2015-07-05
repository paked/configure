// Package configeur is an easy to use multi-layer configuration system.
//
// Examples can be found in the example folder (http://github.com/paked/configeur/blob/master/examples/)
// as well as a getting started guide in the main README file (http://github.com/paked/configeur).
//
// configeur makes use of Checkers, which are used to retrieve values from their respective data sources.
// There are three built in Checkers, Environment, Flag and JSON. Environment retrieves environment variables.
// Flag retrieves variables within the flags of a command. JSON retrieves values from a JSON file/blob. Checkers
// can be essentially thought of as "middlewear for configuration", in fact parts of the package API was inspired by
// negroni (https://github.com/codegangsta/negroni, the awesome net/http middlewear manager) and the standard library's flag package.
//
// It is very easy to create your own Checkers, all they have to do is satisfy the Checker interface.
//  type Checker interface {
//	  Int(name string) (int, error)
// 	  Bool(name string) (int, error)
// 	  String(name string) (string, error)
// 	  Setup() error
//  }
// That is an, Int method, String method and a Bool method. These functions are used to retrieve their respective data types. A setup method is
// also required, where the Checker should initialize itself and throw
// any errors.
//
// If you do create your own Checkers I would be more than happy to add a link to the README in the github repository.
//
package configeur
