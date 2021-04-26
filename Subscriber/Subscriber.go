package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "quickstart-events"
	brokerAddress = "localhost:9092"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)

		if err != nil {
			panic("could not read message " + err.Error())
		}

		/*
			b := []byte(string(msg.Value))
			resp, err := http.Post("http://35.222.55.115:8080/nuevoRegistro", "application/json",
				bytes.NewBuffer(b))

			if err != nil {
				fmt.Print(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
		*/

		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}

func main() {
	for {
		Block{
			Try: func() {
				consume(context.Background())
			},
			Catch: func(e Exception) {
				fmt.Printf("Caught %v\n", e)
			},
		}.Do()
	}
}
