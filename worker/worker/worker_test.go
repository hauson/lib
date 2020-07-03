package worker

import (
	"testing"
	"time"
	"strconv"
	"github.com/lib/worker/mockjob"
)

func TestWorker_AddJob(t *testing.T) {
	worker := New("worker1", 5)
	for i := 0; i < 3; i++ {
		job := mockjob.New("job"+strconv.Itoa(i), uint64(i))
		worker.AddJob(job)
	}

	time.Sleep(8 * time.Second)
	worker.Close()
}
