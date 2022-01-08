package main

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "topic-1",
		Balancer: &kafka.LeastBytes{},
	}

	defer w.Close()

	err := w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("Key-A"),
		Value: []byte("Hello World!"),
	})

	if err != nil {
		log.Fatal("failed when writing message, err:", err)
	}
}
