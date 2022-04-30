package main

import (
	"fmt"

	"github.com/aneshas/flags"
	"github.com/aneshas/flags/env"
	"github.com/aneshas/flags/json"
)

func main() {
	var fs flags.FlagSet

	var (
		host     = fs.String("host", "DB host", "localhost")
		username = fs.String("username", "DB username", "root", env.ByFlagName(), env.Named("FOO"))
		port     = fs.Int("port", "DB port", 3306, env.ByFlagName())
		cfg      = fs.String("config", "JSON Config file", "config.json", env.ByFlagName()) // This is convenient - we can also choose how to provide or override config files
	)

	fs.Parse(
		env.WithPrefix("MYAPP_"),
		json.WithJSONConfig(*cfg),
	)

	fmt.Println("Host: ", *host)
	fmt.Println("Username: ", *username)
	fmt.Println("Port: ", *port)
}
