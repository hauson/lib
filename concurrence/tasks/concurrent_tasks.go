package tasks

import (
	"errors"
	"sync"
	"time"
)

const (
	//ResultFloat64  task flaot64 result
	ResultFloat64 = "float64"
	// ResultInt task int result
	ResultInt = "int"
)

// ResultType task result type
type ResultType string

//TaskAble can do as a task
type TaskAble interface {
	Name() string
	Exec()
	Result() <-chan TaskResult
}

// TaskResult task result
type TaskResult struct {
	Result interface{}
	Type   ResultType
	Error  error
	From   string
}

// ConcurrentTasks  concurrent exec task and block
func ConcurrentTasks(tasks []TaskAble, timeout time.Duration) []TaskResult {
	wg := new(sync.WaitGroup)
	taskResultCh := make(chan TaskResult, len(tasks))
	for _, task := range tasks {
		wg.Add(1)
		go SyncTask(wg, timeout, task, taskResultCh)
	}

	wg.Wait()
	close(taskResultCh)

	taskResults := make([]TaskResult, 0, len(tasks))
	for taskResult := range taskResultCh {
		taskResults = append(taskResults, taskResult)
	}
	return taskResults
}

// SyncTask sync exec task
func SyncTask(wg *sync.WaitGroup, timeout time.Duration, task TaskAble, outCh chan<- TaskResult) {
	defer wg.Done()

	go task.Exec()
	select {
	case <-time.After(timeout):
		outCh <- TaskResult{
			From:  task.Name(),
			Error: errors.New("timeout"),
		}
	case res := <-task.Result():
		outCh <- res
	}
}
