# Go flags

[![Go](https://github.com/aneshas/flags/actions/workflows/go.yml/badge.svg)](https://github.com/aneshas/flags/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/aneshas/flags/badge.svg?branch=trunk)](https://coveralls.io/github/aneshas/flags?branch=trunk)
[![Go Report Card](https://goreportcard.com/badge/github.com/aneshas/flags)](https://goreportcard.com/report/github.com/aneshas/flags)
[![Go Reference](https://pkg.go.dev/badge/github.com/aneshas/flags.svg)](https://pkg.go.dev/github.com/aneshas/flags)

Configuration package inspired by this [talk](https://www.youtube.com/watch?v=PTE4VJIdHPg) / [article](https://peter.bourgon.org/go-for-industrial-programming/) by Peter Bourgon.

The guiding idea behind this package is that `flags are the best way to configure your program` and thus it provides a thin wrapper
around go standard flag package while providig an extra degree of configurability via different extendable / composable mechanisms such as
`environment variables`, `config files` etc... on an opt-in basis.

Two of these mechanisms are provided by default `env` for environment variables and `json` for flat json config giles. Others resolvers can be implemented additionaly.

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

Though where it really shines is when you use it in combination with different config mechanisms such as env flags or json config files:

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

As you can see, flag definitions accept additional arguments where you can set different mechanisms that will be evaluated in a sequence and the first one that yields a valid value will be used.

For example here:

```go
username = fs.String("username", "DB username", "root", json.ByName(), env.Named("UNAME"))
```

If no `-username` flag is provided, the package would next try to find the value as a json key in the provided config file (see below), if the value is not found it will try to read the environment variable `UNAME`. You can put as many resolvers as you want (even of the same kind).

**NOTE** If flag is provided explicitly on the command line it wil **always** take a precedence regardless if value exists in environment or config file for example.

One more thing to keep a note of is that the flags are evaluated in sequence. This means that if you are using flags themselves in order to provide some config for the flags package itself (eg. the config flag in the example above) and thus affecting values of other flags, you need to make sure to define those flags prior to any other flags that might be affected by them (eg. config flag must be parsed first in order to parse the correct config file needed by other flags).
