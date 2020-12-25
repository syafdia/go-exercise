package main

import (
	"fmt"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/semaphore"
)

func main() {
	sem := semaphore.New(3)
	doneC := make(chan bool, 1)
	totProcess := 10

	for i := 1; i <= totProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningProcess(v)

			if v == totProcess {
				doneC <- true
			}
		}(i)
	}

	<-doneC
}

func longRunningProcess(taskID int) {
	fmt.Println(time.Now().Format("15:04:05"), "Running task with ID", taskID)
	time.Sleep(2 * time.Second)
}
