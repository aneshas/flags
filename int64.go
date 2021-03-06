package flags

import "github.com/cstockton/go-conv"

// Int64Value flag type.
type Int64Value struct {
	Value
	V *int64
}

// Set converts and sets a value
func (val *Int64Value) Set(v interface{}) error {
	converted, err := conv.Int64(v)
	if err != nil {
		return err
	}

	*val.V = converted

	return nil
}

// Int64 creates new Int64 flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedence over all resolvers and the default value.
func (fs *FlagSet) Int64(name, usage string, val int64, r ...ResolverFunc) *int64 {
	fs.initFlagSet()

	v := Int64Value{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.Int64(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parseInt64Vals() {
	for i, val := range fs.Values {
		int64Val, ok := val.(Int64Value)
		if !ok {
			continue
		}

		if fs.hasArg(int64Val.name) {
			continue
		}

		for _, r := range int64Val.resolvers {
			if r(fs, int64Val.name, int64(0), i) {
				break
			}
		}
	}
}
