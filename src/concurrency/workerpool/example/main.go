package main

import (
	"log"
	"time"

	"github.com/syafdia/go-exercise/src/concurrency/workerpool"
)

func main() {
	log.SetFlags(log.Ltime)
	wp := workerpool.NewWorkerPool(3)
	wp.Run()

	wp.AddTasks([]*workerpool.Task{
		workerpool.NewTask("A", func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return 1 + 1, nil
		}),
		workerpool.NewTask("B", func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return 2 + 2, nil
		}),
		workerpool.NewTask("C", func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return 3 + 3, nil
		}),
		workerpool.NewTask("D", func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return 4 + 4, nil
		}),
		workerpool.NewTask("E", func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return 5 + 5, nil
		}),
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
