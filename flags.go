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
	fs        *flag.FlagSet
	args      []string
	EnvPrefix string
	Values    []interface{}
}

// FlagSetOption func mainly used for extending FlagSet configuration for
// different types of resolvers.
type FlagSetOption func(*FlagSet)

// Parse parses command line flags and runs all additional resolvers for
// each flag in sequence.
// If command line flag is set it takes precedance over all other resolvers
// which are skipped in that case.
func (fs *FlagSet) Parse(args []string, opts ...FlagSetOption) {
	for _, opt := range opts {
		opt(fs)
	}

	fs.args = args

	fs.fs.Parse(args[1:])
	fs.parseVals()
}

func (fs *FlagSet) parseVals() {
	fs.parseIntVals()
	fs.parseStringVals()
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
