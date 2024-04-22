package main

import (
	"fiber-boilerplate/cmd/server"
	"fiber-boilerplate/pkg/config"
	"flag"
)

var (
	mode = flag.String("mode", "development", "Application mode (development, production)")
)

func main() {
	flag.Parse()

	config.LoadAllConfigs(".env", *mode)

	server.Serve()
}
