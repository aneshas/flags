package env_test

import (
	"os"
	"testing"

	"github.com/aneshas/flags"
	"github.com/aneshas/flags/env"
	"github.com/stretchr/testify/assert"
)

func TestShould_Parse_Env_Values(t *testing.T) {
	var fs flags.FlagSet

	h := "google"
	u := "admin"

	os.Setenv("APP_HOST", h)
	os.Setenv("APP_USR", u)

	var (
		host     = fs.String("host", "DB host", "localhost", env.ByName())
		username = fs.String("username", "DB username", "root", env.Named("USR"))
	)

	fs.Parse(
		[]string{"cmd"},
		env.WithPrefix("APP_"),
	)

	assert.Equal(t, h, *host)
	assert.Equal(t, u, *username)
}
