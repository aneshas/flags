package flags

import "github.com/cstockton/go-conv"

// UintValue flag type.
type UintValue struct {
	Value
	V *uint
}

// Set converts and sets a value
func (val *UintValue) Set(v interface{}) error {
	converted, err := conv.Uint(v)
	if err != nil {
		return err
	}

	*val.V = converted

	return nil
}

// Uint creates new Uint flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedence over all resolvers and the default value.
func (fs *FlagSet) Uint(name, usage string, val uint, r ...ResolverFunc) *uint {
	fs.initFlagSet()

	v := UintValue{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.Uint(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parseUintVals() {
	for i, val := range fs.Values {
		uintVal, ok := val.(UintValue)
		if !ok {
			continue
		}

		if fs.hasArg(uintVal.name) {
			continue
		}

		for _, r := range uintVal.resolvers {
			if r(fs, uintVal.name, uint(0), i) {
				break
			}
		}
	}
}
