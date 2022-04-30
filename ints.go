package flags

import "flag"

type IntValue struct {
	Value
	V *int
}

func (fs *FlagSet) Int(name, usage string, val int, r ...Resolver) *int {
	v := IntValue{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: flag.Int(name, val, usage),
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

		if hasArg(intVal.name) {
			continue
		}

		for _, r := range intVal.resolvers {
			r(fs, intVal.name, 0, i)
		}
	}
}
