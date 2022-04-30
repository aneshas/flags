package flags

// StringValue flag type.
type StringValue struct {
	Value
	V *string
}

// String creates new String flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedance over all resolvers and the default value.
func (fs *FlagSet) String(name, usage string, val string, r ...ResolverFunc) *string {
	fs.initFlagSet()

	v := StringValue{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.String(name, val, usage),
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

		if fs.hasArg(stringVal.name) {
			continue
		}

		for _, r := range stringVal.resolvers {
			if r(fs, stringVal.name, "", i) {
				break
			}
		}
	}
}
