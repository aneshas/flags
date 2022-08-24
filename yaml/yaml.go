package yaml

import (
	"fmt"
	"log"
	"os"

	"github.com/aneshas/flags"
	"gopkg.in/yaml.v3"
)

const configKey = "flags_yaml_core_resolver"

type yamlConfig map[string]interface{}

// WithConfigFile sets json config file path
func WithConfigFile(path *string) flags.FlagSetOption {
	return func(fs *flags.FlagSet) {
		fs.Config[configKey] = path
	}
}

// ByName sets json resolver that will use flag name as json key
func ByName() flags.ResolverFunc {
	return newEnv("")
}

// Named sets json resolver that will use provided name as json key
func Named(name string) flags.ResolverFunc {
	return newEnv(name)
}

func newEnv(name string) flags.ResolverFunc {
	return func(fs *flags.FlagSet, flag string, t interface{}, i int) bool {
		if name == "" {
			name = flag
		}

		config := getConfig(fs)

		val, ok := config[name]
		if !ok {
			return false
		}

		err := fs.Set(i, val, t)
		if err != nil {
			log.Fatalf("json cannot set flag value: %v", err)
		}

		return true
	}
}

func getConfig(fs *flags.FlagSet) yamlConfig {
	config := make(yamlConfig)

	path, ok := fs.Config[configKey]
	if !ok {
		return config
	}

	file := *path.(*string)

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(fmt.Errorf("could not open config file %s: %w", file, err))
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(fmt.Errorf("could not parse config file: %w", err))
	}

	return config
}
