package main

import (
	"log"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/workerpool"
)

func main() {
	wp := workerpool.NewWorkerPool(2, 4)

	wp.AddTask("TA", func() (int, error) {
		time.Sleep(2 * time.Second)
		return 1 + 1, nil
	})
	wp.AddTask("TB", func() (int, error) {
		time.Sleep(2 * time.Second)
		return 2 + 2, nil
	})
	wp.AddTask("TC", func() (int, error) {
		time.Sleep(2 * time.Second)
		return 3 + 3, nil
	})
	wp.AddTask("TD", func() (int, error) {
		time.Sleep(2 * time.Second)
		return 4 + 4, nil
	})
	wp.AddTask("TE", func() (int, error) {
		time.Sleep(2 * time.Second)
		return 5 + 5, nil
	})

	for {
		select {
		case task := <-wp.GetProcessedTask():
			log.Printf("[main] Got task ID: %s, result: %v, err: %v", task.ID, task.Result, task.Err)
		default:
			time.Sleep(1 * time.Second)
			log.Printf("[main] Waiting for result, total queued tasks %d", wp.GetTotalQueuedTask())
		}
	}
}
