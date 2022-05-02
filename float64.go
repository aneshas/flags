package flags

import "github.com/cstockton/go-conv"

// Float64Value flag type.
type Float64Value struct {
	Value
	V *float64
}

// Set converts and sets a value
func (val *Float64Value) Set(v interface{}) error {
	converted, err := conv.Float64(v)
	if err != nil {
		return err
	}

	*val.V = converted 

	return nil
}

// Float64 creates new Float64 flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedance over all resolvers and the default value.
func (fs *FlagSet) Float64(name, usage string, val float64, r ...ResolverFunc) *float64 {
	fs.initFlagSet()

	v := Float64Value{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.Float64(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parseFloat64Vals() {
	for i, val := range fs.Values {
		float64Val, ok := val.(Float64Value)
		if !ok {
			continue
		}

		if fs.hasArg(float64Val.name) {
			continue
		}

		for _, r := range float64Val.resolvers {
			if r(fs, float64Val.name, 1.0, i) {
				break
			}
		}
	}
}
