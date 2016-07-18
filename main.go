package main

import (
	"log"

	"github.com/chazsmi/stock-service/config"
	"github.com/chazsmi/stock-service/cs"
	"github.com/chazsmi/stock-service/handlers"
	proto "github.com/chazsmi/stock-service/proto"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

func main() {
	// Create a new service. Include command line flags.
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

	c, err := config.ReadReturn(config.File)
	if err != nil {
		log.Fatal(err)
	}

	cs.Init(c)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	if err := broker.Init(broker.Option(func(o *broker.Options) {
		o.Addrs = []string{"192.168.99.100:32889"}
	})); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
}
