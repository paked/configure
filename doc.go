// Package configeur is a multi-layer configuration interface.
//
// configeur is based around the idea of middlewear, and is completely pluggable with
// anything that will fulfill the configeur.Checker interface. There are two built in Checkers,
// which can optionally be added to a configeur instance's stack. They are: Environment and JSON.
// Environment reads configuration values from your environment variables, and JSON will retrieve
// values from a JSON string provided through an io.Reader.
//
// As well as this, configeur aims to have a familiar API, which is why it mirrors the flag packages
// API as closely as possible.
package configeur
