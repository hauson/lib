package main

import (
	"errors"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestMultiTask(t *testing.T) {
	const tasknum = 5
	rand.Seed(time.Now().Unix())

	runables := make([]Runable, tasknum)
	for i := 0; i < tasknum; i++ {
		runables[i] = NewRunner()
	}

	for i := 0; i < 3; i++ {
		// timeout encapsulate into runable
		start := time.Now()
		results := ConcurrentTasks(runables, 3*time.Second)
		t.Log(i, " times", " spend: ", time.Since(start).String())
		t.Fail()

		for _, result := range results {
			t.Log(result)
			t.Fail()
		}
	}
}

func TestSingleTask(t *testing.T) {
	var wg sync.WaitGroup
	resultCh := make(chan TaskResult, 1)
	wg.Add(1)
	go AyncTask(&wg, 4*time.Second, NewRunner(), resultCh)
	wg.Wait()
	close(resultCh)
	r := <-resultCh
	t.Log("done ", r)
	t.Fail()
}

//exec with block
func ConcurrentTasks(runables []Runable, timeout time.Duration) []TaskResult {
	wg := new(sync.WaitGroup)
	taskResultCh := make(chan TaskResult, len(runables))
	for _, r := range runables {
		wg.Add(1)
		go AyncTask(wg, timeout, r, taskResultCh)
	}

	wg.Wait()
	close(taskResultCh)

	taskResults := make([]TaskResult, 0, len(runables))
	for taskResult := range taskResultCh {
		taskResults = append(taskResults, taskResult)
	}
	return taskResults
}

type Runner struct {
	ch chan TaskResult
}

func NewRunner() *Runner {
	return &Runner{ch: make(chan TaskResult)}
}

func (r *Runner) Exec() {
	timeout := time.Duration(1+rand.Intn(3)) * time.Second
	time.Sleep(timeout)
	r.ch <- TaskResult{
		Result: 1.5,
		Type:   ResultFloat64,
	}
}

func (r *Runner) Result() <-chan TaskResult {
	return r.ch
}

func AyncTask(wg *sync.WaitGroup, timeout time.Duration, r Runable, outCh chan<- TaskResult) {
	defer wg.Done()

	go r.Exec()
	select {
	case <-time.After(timeout):
		outCh <- TaskResult{
			Error: errors.New("timeout"),
		}
	case res := <-r.Result():
		outCh <- res
	}
}

type Runable interface {
	Exec()
	Result() <-chan TaskResult
}

type TaskResult struct {
	Result interface{}
	Type   interface{}
	Error  error
}

type ResultType string

const (
	ResultFloat64 = "float64"
)
