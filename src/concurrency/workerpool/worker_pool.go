package workerpool

import "log"

// WorkerPool is a contract for Worker Pool implementation
type WorkerPool interface {
	Run()
	AddTasks(tasks []*Task)
	GetProcessedTask() chan *Task
	GetTotalQueuedTask() int
}

type workerPool struct {
	maxWorker      int
	taskC          chan *Task
	queuedTaskC    chan *Task
	processedTaskC chan *Task
}

// NewWorkerPool will create an instance of WorkerPool.
func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker:      maxWorker,
		queuedTaskC:    make(chan *Task),
		processedTaskC: make(chan *Task),
	}

	return wp
}

func (wp *workerPool) Run() {
	wp.run()
}

func (wp *workerPool) AddTask(id string, executor Executor) {
	go func() {
		task := &Task{
			ID:       id,
			executor: executor,
		}
		wp.queuedTaskC <- task
	}()

	log.Printf("[WorkerPool] Task %s has been added", id)
}

func (wp *workerPool) AddTasks(tasks []*Task) {
	go func() {
		for _, task := range tasks {
			wp.queuedTaskC <- task
		}

	}()
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
