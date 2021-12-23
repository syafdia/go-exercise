package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/pipeline/example"
)

func main() {
	Ns := []int{3, 4, 5}
	startTime := time.Now()

	fmt.Printf("Multiple input: %d\n", Ns)

	valCs := withFanOut(Ns)
	valC := withFanIn(valCs)

	vals := []string{}
	for val := range valC {
		vals = append(vals, val)
	}

	result := strings.Join(vals, " | ")

	fmt.Printf("All results: %s\n", result)
	fmt.Printf("Elapsed time without concurrency: %s", time.Since(startTime))
}

// 1 -> N
func withFanOut(Ns []int) []chan string {
	outCs := []chan string{}
	for _, N := range Ns {
		outCs = append(outCs, addFoo(addQuoute(square(multiplyTwo(generateInput(N))))))
	}

	return outCs
}

// 3 -> 1
func withFanIn(valCs []chan string) chan string {
	var wg sync.WaitGroup

	outC := make(chan string)

	for _, valC := range valCs {
		wg.Add(1)

		go func(vC chan string) {
			for val := range vC {
				outC <- val
			}

			wg.Done()
		}(valC)
	}

	go func() {
		wg.Wait()
		close(outC)
	}()

	return outC
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
