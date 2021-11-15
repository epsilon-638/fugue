package fugue

func NewWorkerPool(numWorkers int, buffer int) *WorkerPool {
	workerPool := new(WorkerPool)
	workerPool.New(numWorkers, buffer)
	return workerPool
}
