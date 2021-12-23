package main

import (
	"fmt"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/pipeline/example"
)

func main() {
	startTime := time.Now()

	for i := 0; i < 3; i++ {
		val := example.AddFoo(example.AddQuoute(example.Square(example.MultiplyTwo(i))))
		fmt.Printf("Input: %d, Output %s\n", i, val)

	}

	fmt.Printf("Elapsed time without concurrency: %s\n", time.Since(startTime))
}
