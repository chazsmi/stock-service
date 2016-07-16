package main

import (
	"log"

	"github.com/chazsmi/stock-service/config"
	"github.com/chazsmi/stock-service/cs"
	"github.com/chazsmi/stock-service/handlers"
	proto "github.com/chazsmi/stock-service/proto"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("service.stock"),
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

	proto.RegisterStockHandler(service.Server(), &handlers.Stock{})

	service.Init()

	cs.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
