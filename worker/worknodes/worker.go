package worknodes

import (
	"fmt"

	"github.com/hauson/lib/consistenthash"
	"github.com/hauson/lib/worker/worker"
)

//WorkerNodes a worker list
type WorkerNodes struct {
	workers map[string]*worker.Worker
	router  *consistenthash.Map
}

//New return a container of workers
func New(workerNum int) *WorkerNodes {
	router := consistenthash.New(3, nil)
	workers := make(map[string]*worker.Worker, workerNum)
	for i := 0; i < workerNum; i++ {
		workerName := fmt.Sprintf("Worker:%d", i)
		workers[workerName] = worker.New(workerName, 1024)
		router.Add(workerName)
	}

	return &WorkerNodes{
		workers: workers,
		router:  router,
	}
}

//Close close all worker
func (w *WorkerNodes) Close() {
	for _, worker := range w.workers {
		worker.Close()
	}
}

//AddJob add a job to one of workers basis of key
func (w *WorkerNodes) AddJob(job worker.Job) {
	workerName := w.router.Get(job.Name())
	w.workers[workerName].AddJob(job)
}
