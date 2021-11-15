package fugue

import (
	"log"
	"testing"
)

func testJob() JobResult {
	var result interface{} = "result"
	return JobResult{
		Id:     0,
		Result: result,
	}
}

func TestWokerPoolNew(t *testing.T) {
	workerPool := new(WorkerPool)
	workerPool.New(10, 10)
	result := len(workerPool.workers)
	expected := 10
	if result != expected {
		t.Fatalf("Expected number workers: %d but got %d", expected, result)
	}
}

func TestWorkerPoolAddJob(t *testing.T) {
	workerPool := new(WorkerPool)
	workerPool.New(1, 1)
	log.Print(workerPool)
	log.Print("made it here")
	workerPool.AddJob(testJob)
	log.Print("Made it here")

	job := <-workerPool.jobs

	jobResult := job()
	result := jobResult.Result.(string)
	expected := "result"
	if result != expected {
		t.Fatalf("Expected job result: %s but got %s", expected, result)
	}
}

func TestWorkerPoolGetWorkerAndJob(t *testing.T) {
	workerPool := new(WorkerPool)
	workerPool.New(10, 10)
	workerPool.AddJob(testJob)
	worker, job := workerPool.getWorkerAndJob()
	jobResult := job()
	result := jobResult.Result.(string)
	expected := "result"
	if worker != 0 || result != expected {
		t.Fatalf("Expected job result: worker (%d) result (%s) but got worker (%d) result (%s)", 0, expected, worker, result)
	}
}

func TestWorkerPoolrunJob(t *testing.T) {
	workerPool := new(WorkerPool)
	workerPool.New(10, 10)
	for i := 0; i < 10; i++ {
		workerPool.AddJob(testJob)
	}
	for i := 0; i < 10; i++ {
		worker, job := workerPool.getWorkerAndJob()
		workerPool.runJob(worker, job)
	}
	for i := 0; i < 10; i++ {
		jobResult := <-workerPool.results
		result := jobResult.Result.(string)
		if result != "result" {
			t.Fatal("Failed")
		}
	}
}
