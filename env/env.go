package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aneshas/flags"
)

func WithPrefix(prefix string) flags.FlagSetOption {
	return func(fs *flags.FlagSet) {
		fs.EnvPrefix = prefix
	}
}

func ByFlagName() flags.Resolver {
	return newEnv("")
}

func Named(name string) flags.Resolver {
	return newEnv(name)
}

func newEnv(name string) flags.Resolver {
	return func(fs *flags.FlagSet, flag string, t interface{}, i int) {
		if name == "" {
			name = strings.ToUpper(flag)
		}
		val := os.Getenv(fmt.Sprintf("%s%s", fs.EnvPrefix, name))

		if val == "" {
			return
		}

		switch t.(type) {
		case string:
			v := (fs.Values[i]).(flags.StringValue)
			*v.V = val
		case int:
			ival, err := strconv.Atoi(val)
			if err != nil {
				panic("unsupported type")
			}

			v := (fs.Values[i]).(flags.IntValue)
			*v.V = ival

		default:
			panic("unsupported flag type")
		}
	}
}
