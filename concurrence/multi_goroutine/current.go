package main

import (
	"sync"
)

// Runner able exec with params
type Runner interface {
	Exec() (Result, error)
}

// Result exec result, and can merge another result
type Result interface {
	Data() interface{}
	Append(Result) error
}

func execRunners(runners []Runner, wg *sync.WaitGroup, exitCh <-chan int, errCh chan<- error, resultCh chan<- Result) {
	defer wg.Done()

	var result Result
	for _, runner := range runners {
		select {
		case <-exitCh:
			return
		default:
		}

		aResult, err := runner.Exec()
		if err != nil {
			errCh <- err
		}

		if result == nil {
			result = aResult
		} else {
			result.Append(aResult)
		}
	}
	resultCh <- result
}

func divide(goNum int, runners []Runner) [][]Runner {
	ss := make([][]Runner, 0, goNum)
	per, left := len(runners)/goNum, runners
	for i := 0; i < goNum; i++ {
		var batch []Runner
		if i == goNum-1 {
			batch = left
		} else {
			batch, left = left[:per], left[per:]
		}

		ss = append(ss, batch)
	}
	return ss
}

// ExecByMultiGo exec by multi goroutine
func ExecByMultiGo(goNum int, runners []Runner) (Result, error) {
	resultCh := make(chan Result, goNum)
	errCh := make(chan error, goNum)
	doneCh := make(chan int)
	exitCh := make(chan int)

	wg := new(sync.WaitGroup)
	wg.Add(goNum)

	go func() {
		wg.Wait()
		close(doneCh)
		close(resultCh)
	}()

	for _, batch := range divide(goNum, runners) {
		go execRunners(batch, wg, exitCh, errCh, resultCh)
	}

	select {
	case <-doneCh:
		var res Result
		for itemResult := range resultCh {
			if res == nil {
				res = itemResult
			} else {
				res.Append(itemResult)
			}
		}

		return res, nil
	case err := <-errCh:
		close(exitCh)
		return nil, err
	}
}
