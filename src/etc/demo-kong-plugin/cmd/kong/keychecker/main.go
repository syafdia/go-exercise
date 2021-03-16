package main

import (
	"github.com/Kong/go-pdk/server"
	kongdelivery "github.com/syafdia/demo-kong-plugin/internal/delivery/kong"
)

func main() {
	server.StartServer(kongdelivery.NewKeyChecker, "0.1", 1000)
}
