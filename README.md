# Go flags

[![Go](https://github.com/aneshas/flags/actions/workflows/go.yml/badge.svg)](https://github.com/aneshas/flags/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/aneshas/flags/badge.svg?branch=trunk)](https://coveralls.io/github/aneshas/flags?branch=trunk)
[![Go Report Card](https://goreportcard.com/badge/github.com/aneshas/flags)](https://goreportcard.com/report/github.com/aneshas/flags)
[![Go Reference](https://pkg.go.dev/badge/github.com/aneshas/flags.svg)](https://pkg.go.dev/github.com/aneshas/flags)

Configuration package inspired by this [talk](https://www.youtube.com/watch?v=PTE4VJIdHPg) / [article](https://peter.bourgon.org/go-for-industrial-programming/) by Peter Bourgon.

# Disclaimer
Since the time of writing this package a user on reddit pointed out that Peter has actually went ahead and implemented the package himself although with slight changes in the original design. You can find his package [here](https://pkg.go.dev/github.com/peterbourgon/ff/v3#section-readme)

This package provides a wrapper of the Go standard library flag package configurable via resolvers e.g., environment variables and config files, on an opt-in basis.

Two resolvers are provided by default `env` for environment variables and `json` for JSON config files. Others resolvers can be implemented.

# Example Usage

As (almost) drop-in replacement for `flag` package:

```go
var fs flags.FlagSet

var (
	host     = fs.String("host", "DB host", "localhost")
	username = fs.String("username", "DB username", "root")
	port     = fs.Int("port", "DB port", 3306)
)

fs.Parse(os.Args)
```

Composing different config resolvers is possible. For example using CLI flags, env variables, or JSON config files:

```go
var fs flags.FlagSet

var (
	cfg      = fs.String("config", "JSON Config file", "config.json", env.ByName())
	host     = fs.String("host", "DB host", "localhost")
	username = fs.String("username", "DB username", "root", json.ByName(), env.Named("UNAME"))
	port     = fs.Int("port", "DB port", 3306, json.ByName(), env.ByName())
)

fs.Parse(
	os.Args,
	env.WithPrefix("MYAPP_"), // Optional prefix
	json.WithConfigFile(cfg), // Path to config file
)
```

Flag definitions accept additional arguments where resolver precendence can be set. The first resolver that yields a valid value will be used.

For example here:

```go
username = fs.String("username", "DB username", "root", json.ByName(), env.Named("UNAME"))
```

If no `-username` flag is provided, the package would then try to find the value as a JSON key in the provided config file (see below), if the value is not found it will try to read the environment variable `UNAME`.

So, the order of evaluation is (first value found wins):\
`flag provided via command line` -> `resolvers in sequence (left ot right)` -> `default value`

**NOTE** If flag is provided explicitly on the command line it wil **always** take a precedence regardless if value exists in environment or config file for example.

Flags are evaluated in sequence. This means that if you are using flags themselves in order to provide some config for the flags package itself (eg. the config flag in the example above) and thus affecting values of other flags, you need to make sure to define those flags prior to any other flags that might be affected by them (eg. config flag must be parsed first in order to parse the correct config file needed by other flags).
