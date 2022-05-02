package flags

import (
	"flag"
	"fmt"
)

// ResolverFunc represents a resolver for a primitive flag type.
// Used for implementing additional config mechanisms (see env/ and json/).
type ResolverFunc func(*FlagSet, string, interface{}, int) bool

// Value represents a config value
type Value struct {
	name      string
	resolvers []ResolverFunc
}

// FlagSet represents flags container.
type FlagSet struct {
	fs     *flag.FlagSet
	args   []string
	Values []interface{}
	Config map[string]interface{}
}

// Set sets a flag value for flag at i index
func (fs *FlagSet) Set(i int, v interface{}, t interface{}) error {
	if i > (len(fs.Values) - 1) {
		return fmt.Errorf("no flag at index %d", i)
	}

	switch t.(type) {
	case string:
		val := (fs.Values[i]).(StringValue)
		*val.V = v.(string)

	case int:
		val := (fs.Values[i]).(IntValue)
		val.Set(v)

	case int64:
		val := (fs.Values[i]).(Int64Value)
		val.Set(v)

	case uint:
		val := (fs.Values[i]).(UintValue)
		val.Set(v)

	case uint64:
		val := (fs.Values[i]).(Uint64Value)
		val.Set(v)

	case bool:
		val := (fs.Values[i]).(BoolValue)
		val.Set(v)

	case float64:
		val := (fs.Values[i]).(Float64Value)
		val.Set(v)

	default:
		return fmt.Errorf("unsupported flag type")
	}

	return nil
}

// FlagSetOption func mainly used for extending FlagSet configuration for
// different types of resolvers.
type FlagSetOption func(*FlagSet)

// Parse parses command line flags and runs all additional resolvers for
// each flag in sequence.
// If command line flag is set it takes precedance over all other resolvers
// which are skipped in that case.
func (fs *FlagSet) Parse(args []string, opts ...FlagSetOption) {
	if fs.Config == nil {
		fs.Config = make(map[string]interface{})
	}

	for _, opt := range opts {
		opt(fs)
	}

	fs.args = args

	if len(fs.Values) == 0 {
		return
	}

	fs.fs.Parse(args[1:])
	fs.parseVals()
}

func (fs *FlagSet) parseVals() {
	fs.parseStringVals()
	fs.parseIntVals()
	fs.parseInt64Vals()
	fs.parseUintVals()
	fs.parseUint64Vals()
	fs.parseBoolVals()
	fs.parseFloat64Vals()
}

func (fs *FlagSet) initFlagSet() {
	if fs.fs == nil {
		fs.fs = flag.NewFlagSet("", flag.ExitOnError)
	}
}

func (fs *FlagSet) hasArg(name string) bool {
	for _, arg := range fs.args {
		if arg == fmt.Sprintf("-%s", name) {
			return true
		}
	}

	return false
}
