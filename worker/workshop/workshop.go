package workshop

import (
	"fmt"
	"github.com/lib/worker/worker"
)

type Workshop struct {
	workers []*worker.Worker
}

func New(workCnt int) *Workshop {
	workers := make([]*worker.Worker, workCnt)
	for i := 0; i < workCnt; i++ {
		workers[i] = worker.New(fmt.Sprintf("work%d", i), 30)
	}

	workshop := &Workshop{workers: workers}
	return workshop
}

func (shop *Workshop) AddJob(job worker.Job) {
	var optWorker *worker.Worker
	if worker, ok := shop.selectWorkerByKey(job.Name()); ok {
		optWorker = worker
	} else {
		optWorker = shop.selectWorkerByLVS()
	}

	optWorker.AddJob(job)
}

func (shop *Workshop) selectWorkerByKey(jobName string) (*worker.Worker, bool) {
	for _, worker := range shop.workers {
		if worker.Exist(jobName) {
			return worker, true
		}
	}
	return nil, false
}

func (shop *Workshop) selectWorkerByLVS() *worker.Worker {
	index, jobsMin := 0, shop.workers[0].JobsLen()
	for i := 1; i < len(shop.workers); i++ {
		jobsCnt := shop.workers[i].JobsLen()
		if jobsCnt < jobsMin {
			index, jobsMin = i, jobsCnt
		}
	}
	return shop.workers[index]
}
