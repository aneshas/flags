package flags

import "github.com/cstockton/go-conv"

// BoolValue flag type.
type BoolValue struct {
	Value
	V *bool
}

// Set converts and sets a value
func (val *BoolValue) Set(v interface{}) error {
	converted, err := conv.Bool(v)
	if err != nil {
		return err
	}

	*val.V = converted

	return nil
}

// Bool creates new Bool flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedence over all resolvers and the default value.
func (fs *FlagSet) Bool(name, usage string, val bool, r ...ResolverFunc) *bool {
	fs.initFlagSet()

	v := BoolValue{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.Bool(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parseBoolVals() {
	for i, val := range fs.Values {
		boolVal, ok := val.(BoolValue)
		if !ok {
			continue
		}

		if fs.hasArg(boolVal.name) {
			continue
		}

		for _, r := range boolVal.resolvers {
			if r(fs, boolVal.name, false, i) {
				break
			}
		}
	}
}
