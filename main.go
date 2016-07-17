package main

import (
	"log"

	"github.com/chazsmi/stock-service/config"
	"github.com/chazsmi/stock-service/cs"
	"github.com/chazsmi/stock-service/handlers"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("charlieplc.api.stock"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:   "config_file_path",
				EnvVar: "CONFIG_PATH",
				Usage:  "Config file path",
			},
		),
		micro.Action(func(c *cli.Context) {
			config.File = c.String("config_file_path")
		}),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(handlers.Stock),
		),
	)

	service.Init()

	cs.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
