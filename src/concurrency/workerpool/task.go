package workerpool

type Task struct {
	ID       string
	Result   T
	Err      error
	executor Executor
}

func (t *Task) Execute() {
	t.Result, t.Err = t.executor()
}
