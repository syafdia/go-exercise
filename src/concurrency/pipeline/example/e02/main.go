package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/pipeline/example"
)

func main() {
	N := 3 // Foo '0',Foo '4',Foo '16' -> Sequence is not guarantee.
	startTime := time.Now()
	valC := make(chan string, N)

	fmt.Printf("Input: %d\n", N)

	for i := 0; i < N; i++ {
		go func(input int) {
			val := example.AddFoo(example.AddQuoute(example.Square(example.MultiplyTwo(input))))

			valC <- val
		}(i)
	}

	vals := []string{}
	for i := 0; i < N; i++ {
		vals = append(vals, <-valC)
	}

	result := strings.Join(vals, ",")

	fmt.Printf("Result: %s\n", result)
	fmt.Printf("Elapsed time without concurrency: %s", time.Since(startTime))
}
