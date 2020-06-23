package worker

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/lib/consistenthash"
)

//Worker a goroutine wrapper
type Worker struct {
	exitSig chan int
	inJobs  chan func()
}

//NewWorker return a worker
func NewWorker() *Worker {
	return &Worker{
		exitSig: make(chan int),
		inJobs:  make(chan func(), 1024),
	}
}

//Run start goroutine, usage go Run()
func (w *Worker) Run() {
	for {
		select {
		case <-w.exitSig:
			return
		case job := <-w.inJobs:
			w.execJob(job)
		}
	}
}

func (w *Worker) addJob(job func()) {
	w.inJobs <- job
}

//execJob exec job with recover
func (w *Worker) execJob(jobFun func()) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("worker execJob caught Exception:%v \n", err)
		}
	}()

	jobFun()
}

//Close close the work goroutine
func (w *Worker) Close() {
	close(w.exitSig)
}

//Executor a worker list
type Executor struct {
	workers map[string]*Worker
	jobCh   chan func()
	router  *consistenthash.Map
}

//NewExecutor return a container of workers
func NewExecutor(workerNum int) *Executor {
	router := consistenthash.New(3, nil)
	workers := make(map[string]*Worker, workerNum)
	for i := 0; i < workerNum; i++ {
		workerName := fmt.Sprintf("Worker:%d", i)
		workers[workerName] = NewWorker()
		router.Add(workerName)
	}

	return &Executor{
		workers: workers,
		router:  router,
	}
}

//Run start all worker
func (w *Executor) Run() {
	for _, worker := range w.workers {
		go worker.Run()
	}
}

//Close close all worker
func (w *Executor) Close() {
	for _, worker := range w.workers {
		worker.Close()
	}
}

//AddJob add a job to one of workers basis of key
func (w *Executor) AddJob(key string, job func()) {
	workerName := w.router.Get(key)
	w.workers[workerName].addJob(job)
}
