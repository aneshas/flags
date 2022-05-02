package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aneshas/flags"
)

const configKey = "flags_env_core_resolver"

// WithPrefix sets env variable prefix
func WithPrefix(prefix string) flags.FlagSetOption {
	return func(fs *flags.FlagSet) {
		fs.Config[configKey] = prefix
	}
}

// ByName sets env variable resolver that will use uppercase flag name as env variable name
func ByName() flags.ResolverFunc {
	return newEnv("")
}

// Named sets env variable resolver that will use provided name as env variable name
// (wont uppercase it)
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

		err := fs.Set(i, val, t)
		if err != nil {
			log.Fatalf("json cannot set flag value: %v", err)
		}

		return true
	}
}
