package main

import (
	"fmt"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/pipeline/example"
)

func main() {
	// single()
	multiple()
}

func single() {
	startTime := time.Now()

	valC := addFoo(addQuoute(square(multiplyTwo(generateInput(3)))))

	fmt.Printf("Result: %s\n", <-valC)

	fmt.Printf("Elapsed time without concurrency: %s", time.Since(startTime))
}

func multiple() {
	startTime := time.Now()

	input1 := 3
	input2 := 4
	input3 := 5

	valC1 := addFoo(addQuoute(square(multiplyTwo(generateInput(3)))))
	valC2 := addFoo(addQuoute(square(multiplyTwo(generateInput(4)))))
	valC3 := addFoo(addQuoute(square(multiplyTwo(generateInput(5)))))

	fmt.Printf("Input: %d, Result: %s\n", input1, <-valC1)
	fmt.Printf("Input: %d, Result: %s\n", input2, <-valC2)
	fmt.Printf("Input: %d, Result: %s\n", input3, <-valC3)

	fmt.Printf("Elapsed time without concurrency: %s", time.Since(startTime))
}

func generateInput(N int) chan int {
	resC := make(chan int)

	go func() {
		resC <- N
		close(resC)
	}()

	return resC
}

func multiplyTwo(valC <-chan int) chan int {
	resC := make(chan int)

	go func() {
		for val := range valC {
			fmt.Printf("multiplyTwo: %v\n", val)
			resC <- example.MultiplyTwo(val)
		}

		close(resC)
	}()

	return resC
}

func square(valC <-chan int) chan int {
	resC := make(chan int)

	go func() {
		for val := range valC {
			fmt.Printf("square: %v\n", val)
			resC <- example.Square(val)
		}

		close(resC)
	}()

	return resC
}

func addQuoute(valC <-chan int) chan string {
	resC := make(chan string)

	go func() {
		for val := range valC {
			fmt.Printf("addQuoute: %v\n", val)
			resC <- example.AddQuoute(val)
		}

		close(resC)
	}()

	return resC
}

func addFoo(valC <-chan string) chan string {
	resC := make(chan string)

	go func() {
		for val := range valC {
			fmt.Printf("addFoo: %v\n", val)
			resC <- example.AddFoo(val)
		}

		close(resC)
	}()

	return resC
}
