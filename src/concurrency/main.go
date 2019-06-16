package main

import (
	"fmt"
	"time"
)

func main() {
	pipelines := squareInt(multiplyTwo(generateInt(10)))

	for v := range pipelines {
		fmt.Println(time.Now())
		fmt.Println(v)
	}

	// outC := pipeline.New(generateIntInterface(10)).
	// 	Pipe(func(in interface{}) (interface{}, error) {
	// 		time.Sleep(2000 * time.Millisecond)
	// 		return (in.(int) * 2), nil
	// 	}).
	// 	Pipe(func(in interface{}) (interface{}, error) {
	// 		time.Sleep(2000 * time.Millisecond)
	// 		return fmt.Sprintf("%v Foo", in), nil
	// 	}).
	// 	Pipe(func(in interface{}) (interface{}, error) {
	// 		time.Sleep(2000 * time.Millisecond)
	// 		return fmt.Sprintf("%v Bar", in), nil
	// 	}).
	// 	Pipe(func(in interface{}) (interface{}, error) {
	// 		time.Sleep(2000 * time.Millisecond)
	// 		return fmt.Sprintf("%v Baz", in), nil
	// 	}).
	// 	Merge()

	// for v := range outC {
	// 	fmt.Println(time.Now())
	// 	fmt.Println(v)
	// }
}

func generateIntInterface(n int) chan interface{} {
	outC := make(chan interface{})

	go func() {
		defer close(outC)
		for i := 0; i < n; i++ {
			outC <- i
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	return outC
}

func generateInt(n int) chan int {
	fmt.Println("Start 1")
	outC := make(chan int)

	go func() {
		defer close(outC)
		for i := 0; i < n; i++ {
			outC <- i
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	return outC
}

func multiplyTwo(inC <-chan int) <-chan int {

	fmt.Println("Start 2ß")
	outC := make(chan int)

	go func() {
		defer close(outC)
		for v := range inC {
			outC <- v * 2
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	return outC
}

func squareInt(inC <-chan int) <-chan int {
	fmt.Println("Start 3")
	outC := make(chan int)

	go func() {
		defer close(outC)
		for v := range inC {
			outC <- v * v
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	return outC
}
