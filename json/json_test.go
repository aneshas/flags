package json_test

import (
	"testing"

	"github.com/aneshas/flags"
	"github.com/aneshas/flags/json"
	"github.com/stretchr/testify/assert"
)

func TestShould_Parse_JSON_Values(t *testing.T) {
	var fs flags.FlagSet

	var (
		host     = fs.String("host", "DB host", "localhost", json.ByName())
		username = fs.String("username", "DB username", "root", json.Named("usr"))
		cfg      = fs.String("config", "JSON Config file", "./testdata/config.json")
	)

	fs.Parse(
		[]string{"cmd"},
		json.WithConfigFile(cfg),
	)

	assert.Equal(t, "google", *host)
	assert.Equal(t, "admin", *username)
}
