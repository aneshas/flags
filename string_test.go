package flags_test

import (
	"testing"

	"github.com/aneshas/flags"
	"github.com/stretchr/testify/assert"
)

func TestShould_Use_Default_Flag_Value(t *testing.T) {
	val := "http://google.com"

	var (
		fs   flags.FlagSet
		host = fs.String("host", "DB host", val)
	)

	args := []string{
		"cmd",
	}

	fs.Parse(args)

	assert.Equal(t, val, *host)
}

func TestShould_Parse_From_Args(t *testing.T) {
	val := "http://localhost"

	var (
		fs   flags.FlagSet
		host = fs.String("host", "DB host", "http://google.com")
	)

	args := []string{
		"cmd",
		"-host",
		val,
	}

	fs.Parse(args)

	assert.Equal(t, val, *host)
}

func TestShould_Should_Use_Value_From_First_Resolver_That_Provides_It(t *testing.T) {
	val := "http://duckduckgo"

	var (
		fs   flags.FlagSet
		host = fs.String(
			"host",
			"DB host",
			"http://google.com",
			newStringMockResolver(val),
			newStringMockResolver("foo"),
			newStringMockResolver("foo"),
		)
	)

	args := []string{
		"cmd",
	}

	fs.Parse(args)

	assert.Equal(t, val, *host)
}

func TestShould_Should_Fall_Back_To_Arg_Value_If_Provided(t *testing.T) {
	val := "http://duckduckgo"

	var (
		fs   flags.FlagSet
		host = fs.String(
			"host",
			"DB host",
			"http://google.com",
			newStringMockResolver("foo"),
		)
	)

	args := []string{
		"cmd",
		"-host",
		val,
	}

	fs.Parse(args)

	assert.Equal(t, val, *host)
}

func newStringMockResolver(val string) flags.ResolverFunc {
	return func(fs *flags.FlagSet, flag string, t interface{}, i int) bool {
		v := (fs.Values[i]).(flags.StringValue)
		*v.V = val

		return true
	}
}
