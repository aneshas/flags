package flags_test

import (
	"testing"

	"github.com/aneshas/flags"
	"github.com/stretchr/testify/assert"
)

var args = []string{
	"cmd",
}

func TestShould_Convert_And_Set_Flag_Values(t *testing.T) {
	var (
		strVal    string  = "http://google.com"
		intVal    int     = 8080
		largeIVal int64   = 100000
		smallVal  uint    = 1
		largeUVal uint64  = 100000
		boolVal   bool    = true
		floatVal  float64 = 4.50
	)

	var (
		fs         flags.FlagSet
		strFlag    = fs.String("strFlag", "Flag usage", "")
		intFlag    = fs.Int("intFlag", "Flag usage", 0)
		largeIFlag = fs.Int64("int46Flag", "Flag usage", 0)
		smallFlag  = fs.Uint("uFlag", "Flag usage", 0)
		largeUFlag = fs.Uint64("u64Flag", "Flag usage", 0)
		boolFlag   = fs.Bool("boolFlag", "Flag usage", false)
		floatFlag  = fs.Float64("floatFlag", "Flag usage", 0.0)
	)

	fs.Parse(args)

	err := fs.Set(0, strVal, "")
	assert.NoError(t, err)

	err = fs.Set(1, intVal, intVal)
	assert.NoError(t, err)

	err = fs.Set(2, largeIVal, largeIVal)
	assert.NoError(t, err)

	err = fs.Set(3, smallVal, smallVal)
	assert.NoError(t, err)

	err = fs.Set(4, largeUVal, largeUVal)
	assert.NoError(t, err)

	err = fs.Set(5, boolVal, boolVal)
	assert.NoError(t, err)

	err = fs.Set(6, floatVal, floatVal)
	assert.NoError(t, err)

	assert.Equal(t, strVal, *strFlag)
	assert.Equal(t, intVal, *intFlag)
	assert.Equal(t, largeIVal, *largeIFlag)
	assert.Equal(t, smallVal, *smallFlag)
	assert.Equal(t, largeUVal, *largeUFlag)
	assert.Equal(t, boolVal, *boolFlag)
	assert.Equal(t, floatVal, *floatFlag)
}

func TestShould_Throw_Error_If_Setting_Non_Defined_Flag(t *testing.T) {
	var fs flags.FlagSet

	fs.Parse(args)

	err := fs.Set(0, "", "")

	assert.Error(t, err)
}
