package main

import (
	"fmt"
	"os"

	"github.com/aneshas/flags"
	"github.com/aneshas/flags/env"
	"github.com/aneshas/flags/json"
)

func main() {
	var fs flags.FlagSet

	var (
		host     = fs.String("host", "DB host", "localhost")
		username = fs.String("username", "DB username", "root", json.ByName(), env.Named("UNAME"))
		port     = fs.Int("port", "DB port", 3306, json.ByName(), env.ByName())
		cfg      = fs.String("config", "JSON Config file", "config.json", env.ByName())
	)

	fs.Parse(
		os.Args,
		env.WithPrefix("MYAPP_"),
		json.WithConfigFile(cfg),
	)

	fmt.Println("Host: ", *host)
	fmt.Println("Username: ", *username)
	fmt.Println("Port: ", *port)
}
