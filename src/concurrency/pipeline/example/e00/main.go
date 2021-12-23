package main

import (
	"fmt"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/pipeline"
	"github.com/syafdia/go-exercise/src/concurrency/pipeline/example"
)

func main() {
	N := 5
	startTime := time.Now()

	for i := 0; i < N; i++ {
		result := example.AddFoo(example.AddQuoute(example.Square(example.MultiplyTwo(i))))
		fmt.Printf("Result: %s\n", result)
	}

	fmt.Printf("Elapsed time without concurrency: %s", time.Since(startTime)) // ~40 seconds

	outC := pipeline.New(func(inC chan interface{}) {
		defer close(inC)
		for i := 0; i < N; i++ {
			inC <- i
		}
	}).
		Pipe(func(in interface{}) (interface{}, error) {
			return example.MultiplyTwo(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return example.Square(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return example.AddQuoute(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return example.AddFoo(in.(string)), nil
		}).
		Merge()

	startTimeC := time.Now()
	for result := range outC {
		fmt.Printf("Result: %s\n", result)
	}

	fmt.Printf("Elapsed time with concurrency: %s", time.Since(startTimeC)) // ~16 seconds
}
