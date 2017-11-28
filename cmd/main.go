package main

import (
	"../config"
	"../server"
	"github.com/nestgo/nest"
)

func main() {
	cfg := config.GetConfig()

	nest.Register(
		&server.GreeterServer{},
		
	)
	nest.Run(cfg.Address)
}
