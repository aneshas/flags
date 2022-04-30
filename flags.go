package flags

import (
	"flag"
	"fmt"
	"os"
)

type FlagSet struct {
	EnvPrefix string
	Values    []interface{}
}

type Value struct {
	name      string
	resolvers []Resolver
}

func (fs *FlagSet) Parse(opts ...FlagSetOption) {
	for _, opt := range opts {
		opt(fs)
	}

	flag.Parse()
	fs.parseVals()
}

func (fs *FlagSet) parseVals() {
	fs.parseIntVals()
	fs.parseStringVals()
}

func hasArg(name string) bool {
	for _, arg := range os.Args {
		if arg == fmt.Sprintf("-%s", name) {
			return true
		}
	}

	return false
}

type Resolver func(*FlagSet, string, interface{}, int)

type FlagSetOption func(*FlagSet)
