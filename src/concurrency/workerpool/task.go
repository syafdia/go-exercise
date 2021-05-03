package workerpool

// Task will hold wrapped function which will
// be processed by Worker Pool.
type Task struct {
	ID       string
	Result   T
	Err      error
	executor Executor
}

func NewTask(id string, executor Executor) *Task {
	return &Task{
		ID:       id,
		executor: executor,
	}
}

// Execute will run wrapped function on Task instance
// and set the Result & Error property.
func (t *Task) Execute() {
	t.Result, t.Err = t.executor()
}
