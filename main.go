package main

import (
	"fiber-boilerplate/cmd/server"
	"fiber-boilerplate/pkg/config"
)

func main() {
	config.LoadAllConfigs(".env")

	server.Serve()
}
