package worker

import (
	"log"
	"fmt"

	"github.com/lib/worker/roll"
)

// Worker a goroutine wrapper
type Worker struct {
	*roll.JobRoll
	name    string
	exitSig chan int
	jobs    chan Job
}

// New return a worker
func New(name string, jobsMax int) *Worker {
	worker := &Worker{
		name:    name,
		JobRoll: roll.New(),
		exitSig: make(chan int),
		jobs:    make(chan Job, jobsMax),
	}
	go worker.run()

	return worker
}

func (w *Worker) Name() string {
	return w.name
}

func (w *Worker) AddJob(job Job) {
	w.Enroll(job.Name())
	w.jobs <- job
}

func (w *Worker) JobsLen() int {
	return len(w.jobs)
}

// Close close the work goroutine
func (w *Worker) Close() {
	close(w.exitSig)
}

func (w *Worker) run() {
	for {
		select {
		case <-w.exitSig:
			return
		case job := <-w.jobs:
			w.doWithRecover(job)
		}
	}
}

func (w *Worker) doWithRecover(job Job) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("catch err ", err, " then recover")
		}
	}()
	defer w.Unroll(job.Name())

	fmt.Println(w.name + "-->" + job.Exec())
}
