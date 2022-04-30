package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aneshas/flags"
)

const configKey = "flags_env_core_resolver"

func WithPrefix(prefix string) flags.FlagSetOption {
	return func(fs *flags.FlagSet) {
		fs.Config[configKey] = prefix
	}
}

func ByName() flags.ResolverFunc {
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

		var prefix string

		if p, ok := fs.Config[configKey]; ok {
			prefix = p.(string)
		}

		val := os.Getenv(fmt.Sprintf("%s%s", prefix, name))

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
