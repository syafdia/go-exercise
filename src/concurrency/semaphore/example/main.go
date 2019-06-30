package main

import (
	"fmt"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/semaphore"
)

func main() {
	sem := semaphore.New(3)
	doneC := make(chan bool, 1)
	N := 10

	for i := 1; i <= N; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			fmt.Println("XXX ->", v)
			time.Sleep(1 * time.Second)

			if v == N {
				doneC <- true
			}
		}(i)
	}

	<-doneC
}
