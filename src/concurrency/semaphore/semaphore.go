package semaphore

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan int8
}

func New(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan int8, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- 1
}

func (s *semaphore) Release() {
	<-s.semC
}
