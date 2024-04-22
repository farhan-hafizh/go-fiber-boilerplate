package main

import (
	"flag"
	"go-fiber-boilerplate/cmd/server"
	"go-fiber-boilerplate/pkg/config"
)

var (
	mode = flag.String("mode", "development", "Application mode (development, production)")
)

func main() {
	flag.Parse()

	config.LoadAllConfigs(".env", *mode)

	server.Serve()
}
