package singleflight

import "sync"

type Singler struct {
	executors map[string]Executor
	sync.Mutex
}

func New() *Singler {
	return &Singler{
		executors: make(map[string]Executor),
	}
}

func (s *Singler) Flight(e Executor) {
	s.Lock()
	defer s.Unlock()

	preExecutor, ok := s.executors[e.Key()]
	if ok {
		if !preExecutor.Valid() {
			preExecutor.SetValid(true)
		}

		go func() {
			<-preExecutor.Done()
			e.CopyResult(preExecutor)
			e.Close()
		}()
		return
	}

	s.executors[e.Key()] = e
	go func() {
		e.Do()
		e.SetValid(false)
		<-e.Done()
	}()
}

type Executor interface {
	Key() string
	Done() <-chan int
	Result() (interface{}, error)
	CopyResult(Executor)
	Do()
	Valid() bool
	SetValid(bool)
	Close()
}
