package flags

import "github.com/cstockton/go-conv"

// IntValue flag type.
type IntValue struct {
	Value
	V *int
}

// Set converts and sets a value
func (val *IntValue) Set(v interface{}) error {
	converted, err := conv.Int(v)
	if err != nil {
		return err
	}

	*val.V = converted

	return nil
}

// Int creates new Int flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedence over all resolvers and the default value.
func (fs *FlagSet) Int(name, usage string, val int, r ...ResolverFunc) *int {
	fs.initFlagSet()

	v := IntValue{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.Int(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parseIntVals() {
	for i, val := range fs.Values {
		intVal, ok := val.(IntValue)
		if !ok {
			continue
		}

		if fs.hasArg(intVal.name) {
			continue
		}

		for _, r := range intVal.resolvers {
			if r(fs, intVal.name, 0, i) {
				break
			}
		}
	}
}
