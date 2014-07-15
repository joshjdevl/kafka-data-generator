package main

import (
	"os"
	"os/signal"
	"log"
	"github.com/wvanbergen/kafka/consumergroup"
)

func main() {
	consumerGroupName := "my_consumer_group_name2"
	kafkaTopic := "topic"
	zookeeper := []string{"149.204.61.37:2181"}

	consumer, consumerErr := consumergroup.JoinConsumerGroup(consumerGroupName, kafkaTopic, zookeeper, nil)
	if consumerErr != nil {
		log.Fatalln(consumerErr)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		consumer.Close()
	}()

	eventCount := 0

	stream := consumer.Stream()
	for {
		event, ok := <-stream
		if !ok {
			break
		}

		// Process event
		log.Println(string(event.Value))

		eventCount += 1
	}

	log.Printf("Processed %d events.", eventCount)

}

