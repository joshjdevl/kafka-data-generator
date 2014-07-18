package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	client, err := sarama.NewClient("client_id", []string{"149.204.61.37:49160"}, sarama.NewClientConfig())
	if err != nil {
		panic(err)
	} else {
		log.Println("> connected")
	}
	defer client.Close()

	producer, err := sarama.NewProducer(client, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	for {
		err = producer.SendMessage("topic", nil, sarama.StringEncoder("testing 123"))
		if err != nil {
			panic(err)
		} else {
			log.Println("> message sent")
		}
	}
}
