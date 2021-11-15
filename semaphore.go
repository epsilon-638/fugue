package fugue

type Semaphore struct {
	sema chan struct{}
}

type SemaphoreInterface interface {
	Acquire()
	Release()
}
