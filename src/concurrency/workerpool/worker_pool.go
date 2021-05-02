package workerpool

import (
	"log"
)

type WorkerPool interface {
	AddTask(id string, executor Executor)
	GetProcessedTask() chan *Task
	GetTotalQueuedTask() int
}

type workerPool struct {
	maxWorker      int
	tasks          []*Task
	queuedTaskC    chan *Task
	processedTaskC chan *Task
}

func NewWorkerPool(maxWorker int, maxTask int) WorkerPool {
	wp := &workerPool{
		maxWorker:      maxWorker,
		queuedTaskC:    make(chan *Task, maxTask),
		processedTaskC: make(chan *Task, maxTask),

		tasks: []*Task{},
	}

	defer wp.run()
	return wp
}

func (wp *workerPool) AddTask(id string, executor Executor) {
	task := &Task{
		ID:       id,
		executor: executor,
	}
	wp.queuedTaskC <- task

	log.Printf("[WorkerPool] Task %s has been added", id)
}

func (wp *workerPool) GetTotalQueuedTask() int {
	return len(wp.queuedTaskC)
}

func (wp *workerPool) GetProcessedTask() chan *Task {
	return wp.processedTaskC
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				log.Printf("[WorkerPool] Worker %d start task %s", workerID, task.ID)

				task.Execute()
				wp.processedTaskC <- task

				log.Printf("[WorkerPool] Worker %d finish task %s", workerID, task.ID)
			}
		}(i + 1)
	}
}
