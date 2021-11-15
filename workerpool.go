package fugue

type ID = int

type JobResult struct {
	Id     ID
	Result interface{}
}

type Job func() JobResult

type WorkerPool struct {
	workers chan ID
	jobs    chan Job
	results chan JobResult
}

type WorkerPoolInterface interface {
	New(numWorkers int)
	Run()
	getWorkerAndJob() (ID, Job)
	AddJob(job Job)
}

func (wp *WorkerPool) New(numWorkers int, buffer int) {
	wp.newChans(numWorkers, buffer)
	for worker_id := 0; worker_id < numWorkers; worker_id++ {
		wp.workers <- worker_id
	}
}

func (wp *WorkerPool) newChans(numWorkers int, buffer int) {
	wp.newWorkerChan(numWorkers)
	wp.newJobChan(buffer)
	wp.newResultChan(buffer)
}

func (wp *WorkerPool) newWorkerChan(numWorkers int) {
	workers_chan := make(chan ID, numWorkers)
	wp.workers = workers_chan
}

func (wp *WorkerPool) newJobChan(buffer int) {
	jobs_chan := make(chan Job, buffer)
	wp.jobs = jobs_chan
}

func (wp *WorkerPool) newResultChan(buffer int) {
	results_chan := make(chan JobResult, buffer)
	wp.results = results_chan
}

func (wp *WorkerPool) Run() {
	for {
		worker, job := wp.getWorkerAndJob()
		go wp.runJob(worker, job)
	}
}

func (wp *WorkerPool) getWorkerAndJob() (ID, Job) {
	worker := <-wp.workers
	job := <-wp.jobs
	return worker, job
}

func (wp *WorkerPool) runJob(worker ID, job Job) {
	result := job()
	wp.workers <- worker
	wp.results <- result
}

func (wp *WorkerPool) AddJob(job Job) {
	wp.jobs <- job
}
