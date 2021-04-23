package main

import (
	"context"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "quickstart-events"
	brokerAddress = "kafka:9092"
)

func produce(ctx context.Context, string messageSend) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(i)),
		// create an arbitrary message payload for the value
		Value: []byte(messageSend),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

}

func main() {
	// create a new context
	ctx := context.Background()
	for i := 0; i < 20; i++ {
		go produce(ctx, "mensaje"+i)
	}
}
