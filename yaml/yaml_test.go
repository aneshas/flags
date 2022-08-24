package yaml_test

import (
	"testing"

	"github.com/aneshas/flags"
	"github.com/aneshas/flags/yaml"
	"github.com/stretchr/testify/assert"
)

func TestShould_Parse_YAML_Values(t *testing.T) {
	var fs flags.FlagSet

	var (
		host     = fs.String("host", "DB host", "localhost", yaml.ByName())
		username = fs.String("username", "DB username", "root", yaml.Named("usr"))
		cfg      = fs.String("config", "JSON Config file", "./testdata/config.yaml")
	)

	fs.Parse(
		[]string{"cmd"},
		yaml.WithConfigFile(cfg),
	)

	assert.Equal(t, "google", *host)
	assert.Equal(t, "admin", *username)
}
