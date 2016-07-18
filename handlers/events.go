package handlers

import (
	"encoding/json"
	"log"

	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

var topic = "charlieplc.topic.stock"

type StockChange struct {
	Sku    string `json:"sku"`
	Amount int32  `json:"amount"`
}

func pub(stockUpdated StockChange) {
	j, err := json.Marshal(stockUpdated)
	if err != nil {
		log.Println("Trying to pub json marhsal failed - " + err.Error())
	}
	msg := &broker.Message{
		Body: j,
	}
	if err := broker.Publish(topic, msg); err != nil {
		log.Printf("[pub] failed: %v", err)
	} else {
		log.Println("[pub] published stock change:", string(msg.Body))
	}
}
