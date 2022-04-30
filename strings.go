package flags

import "flag"

type StringValue struct {
	Value
	V *string
}

func (fs *FlagSet) String(name, usage string, val string, r ...Resolver) *string {
	v := StringValue{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: flag.String(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parseStringVals() {
	for i, val := range fs.Values {
		stringVal, ok := val.(StringValue)
		if !ok {
			continue
		}

		if hasArg(stringVal.name) {
			continue
		}

		for _, r := range stringVal.resolvers {
			r(fs, stringVal.name, "", i)
		}
	}
}
