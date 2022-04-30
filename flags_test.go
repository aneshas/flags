package flags_test

import (
	"testing"

	"github.com/aneshas/flags"
	"github.com/aneshas/flags/env"
)

func Test(t *testing.T) {
	var fs flags.FlagSet

	var (
		host     = fs.String("host", "DB host", "http://google.com")
		username = fs.String("host", "DB username", "root", env.Named("USERNAME"))
	)

	fs.Parse(
		env.WithPrefix("MYAPP_"),
	)

	_ = host
	_ = username
}
