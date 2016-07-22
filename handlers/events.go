package handlers

import (
	"log"

	"github.com/chazsmi/stock-service/proto"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/protobuf/proto"
)

var topic = "charlieplc.topic.stock"

func pub(stockUpdated *stock.StockReadResponse) {
	p, err := proto.Marshal(stockUpdated)
	if err != nil {
		log.Println("Trying to pub json marhsal failed - " + err.Error())
	}
	msg := &broker.Message{
		Header: map[string]string{"ContentType": "application/x-protobuf"},
		Body:   p,
	}
	if err := broker.Publish(topic, msg); err != nil {
		log.Printf("[pub] failed: %v", err)
	} else {
		log.Println("[pub] published stock change:", string(msg.Body))
	}
}
