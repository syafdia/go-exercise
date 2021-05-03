package workerpool

// T is a type alias to accept any type.
type T = interface{}

// Executor is a type alias for Worker Pool parameter.
type Executor = func() (T, error)
