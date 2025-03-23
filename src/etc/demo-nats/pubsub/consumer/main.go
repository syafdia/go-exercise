package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	handleError(err)

	js, err := jetstream.New(nc)
	handleError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// get existing stream handle
	s, err := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     "HI_STREAM",
		Subjects: []string{"hi.*"},
	})
	handleError(err)

	// retrieve consumer handle from a stream
	cons, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "HI_CONSUMER",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	handleError(err)

	log.Println("[publisher] Will consume from callback")

	cc, err := cons.Consume(func(msg jetstream.Msg) {
		log.Println("[consumer] Receive message via consume:", string(msg.Data()))
		msg.Ack()
	})
	handleError(err)

	defer cc.Stop()

	waitC := make(chan int, 1)
	<-waitC
}

func handleError(err error) {
	if err != nil {
		panic(fmt.Sprintf("[consumer]: got error, error: %s", err))
	}
}
