package main

import (
	"flag"
	"os"
	"text/template"
)

//go:generate sh -c "go run main.go -name=String -type=string > ../../string.go"
//go:generate sh -c "go run main.go -name=Int -type=int -val 0 > ../../int.go"
//go:generate sh -c "go run main.go -name=Int64 -type=int64 -val 'int64(0)' > ../../int64.go"
//go:generate sh -c "go run main.go -name=Uint -type=uint -val 'uint(0)' > ../../uint.go"
//go:generate sh -c "go run main.go -name=Uint64 -type=uint64 -val 'uint64(0)' > ../../uint64.go"
//go:generate sh -c "go run main.go -name=Bool -type=bool -val 'false' > ../../bool.go"
//go:generate sh -c "go run main.go -name=Float64 -type=float64 -val '1.0' > ../../float64.go"
type data struct {
	T     string
	TName string
	TVal  string
}

func main() {
	var d data

	flag.StringVar(&d.T, "type", "", "The type.")
	flag.StringVar(&d.TName, "name", "", "The name of the type.")
	flag.StringVar(&d.TVal, "val", `""`, "A value of the type.")

	flag.Parse()

	t := template.Must(template.New("flags").Parse(tpl))
	t.Execute(os.Stdout, d)
}

var tpl = `package flags

import "github.com/cstockton/go-conv"

// {{.TName}}Value flag type.
type {{.TName}}Value struct {
	Value
	V *{{.T}}
}

// Set converts and sets a value
func (val *{{.TName}}Value) Set(v interface{}) error {
	converted, err := conv.{{.TName}}(v)
	if err != nil {
		return err
	}

	*val.V = converted 

	return nil
}

// {{.TName}} creates new {{.TName}} flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedence over all resolvers and the default value.
func (fs *FlagSet) {{.TName}}(name, usage string, val {{.T}}, r ...ResolverFunc) *{{.T}} {
	fs.initFlagSet()

	v := {{.TName}}Value{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.{{.TName}}(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parse{{.TName}}Vals() {
	for i, val := range fs.Values {
		{{.T}}Val, ok := val.({{.TName}}Value)
		if !ok {
			continue
		}

		if fs.hasArg({{.T}}Val.name) {
			continue
		}

		for _, r := range {{.T}}Val.resolvers {
			if r(fs, {{.T}}Val.name, {{.TVal}}, i) {
				break
			}
		}
	}
}
`
