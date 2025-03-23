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
	if err != nil {
		panic(err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("[publisher] Will publish some messages")

	js.Publish(ctx, "hi.joe", buildGreetPayload("Joe"))
	time.Sleep(1 * time.Second)

	js.Publish(ctx, "hi.john", buildGreetPayload("John"))
	time.Sleep(1 * time.Second)

	js.Publish(ctx, "hi.doe", buildGreetPayload("Doe"))
	time.Sleep(1 * time.Second)

	log.Println("[publisher] Messages has been published")
}

func buildGreetPayload(name string) []byte {
	return []byte(fmt.Sprintf("Hello %s, %d", name, time.Now().Unix()))
}
