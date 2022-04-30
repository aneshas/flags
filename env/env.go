package env

import (
	"fmt"
	"log"
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

func ByFlagName() flags.ResolverFunc {
	return newEnv("")
}

func Named(name string) flags.ResolverFunc {
	return newEnv(name)
}

func newEnv(name string) flags.ResolverFunc {
	return func(fs *flags.FlagSet, flag string, t interface{}, i int) bool {
		if name == "" {
			name = strings.ToUpper(flag)
		}
		val := os.Getenv(fmt.Sprintf("%s%s", fs.EnvPrefix, name))

		if val == "" {
			return false
		}

		switch t.(type) {
		case string:
			v := (fs.Values[i]).(flags.StringValue)
			*v.V = val
		case int:
			ival, err := strconv.Atoi(val)
			if err != nil {
				// TODO - Call fs.usage
				log.Fatalf("cannot convert value to int: %v", val)
			}

			v := (fs.Values[i]).(flags.IntValue)
			*v.V = ival

		default:
			log.Fatalf("unsupported flag type: %t", t)
		}

		return true
	}
}
