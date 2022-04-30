package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aneshas/flags"
	"github.com/cstockton/go-conv"
)

const configKey = "flags_json_core_resolver"

type jsonConfig map[string]interface{}

func WithConfigFile(path *string) flags.FlagSetOption {
	return func(fs *flags.FlagSet) {
		fs.Config[configKey] = path
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
			name = flag
		}

		config := getConfig(fs)

		val, ok := config[name]
		if !ok {
			return false
		}

		switch t.(type) {
		case string:
			v := (fs.Values[i]).(flags.StringValue)
			*v.V = val.(string)

		case int:
			v := (fs.Values[i]).(flags.IntValue)

			got, err := conv.Int(val)
			if err != nil {
				log.Fatalf("cannot convert value to int: %v", val)
			}

			*v.V = got

		default:
			panic("unsupported flag type")
		}

		return true
	}
}

func getConfig(fs *flags.FlagSet) jsonConfig {
	config := make(jsonConfig)

	path, ok := fs.Config[configKey]
	if !ok {
		return config
	}

	file := *path.(*string)

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(fmt.Errorf("could not open config file %s: %w", file, err))
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(fmt.Errorf("could not parse config file: %w", err))
	}

	return config
}
