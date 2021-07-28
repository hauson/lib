package singleflight

import (
	"sync"
	"time"
	"fmt"

	"github.com/hauson/lib/container/prioqueue"
)

type Singler struct {
	delTimer  *time.Timer
	delQueue  *prioqueue.PrioQueue
	executors map[string]Executor
	sync.Mutex
}

func New() *Singler {
	return &Singler{
		delTimer:  time.NewTimer(1 * time.Hour),
		executors: make(map[string]Executor),
		delQueue:  prioqueue.New(),
	}
}

func (s *Singler) Flight(e Executor) {
	s.Lock()
	defer s.Unlock()

	preExecutor, ok := s.executors[e.Key()]
	if ok {
		go func() {
			preExecutor.Wait()
			e.CopyResult(preExecutor)
			e.Close()
		}()
		return
	}

	s.executors[e.Key()] = e
	s.delQueue.Add(&delItem{e})
	s.resetTimer()
	go func() {
		e.Do()
	}()
}

// Del excutor
func (s *Singler) Del(e Executor) {
	s.Lock()
	delete(s.executors, e.Key())
	fmt.Println("del executor:", e.Key())
	s.Unlock()
}

func (s *Singler) resetTimer() {
	if s.delTimer != nil {
		s.delTimer.Stop()
	}

	item, ok := s.delQueue.Root()
	if !ok {
		fmt.Println("reset timer: hour")
		s.delTimer.Reset(time.Hour)
		return
	}

	rootItem := item.(*delItem)
	d := rootItem.Expire().Sub(time.Now())
	s.delTimer = time.AfterFunc(d, func() {
		if root, ok := s.delQueue.PopRoot(); ok {
			item := root.(*delItem)
			s.Del(item.Executor)
		}

		s.resetTimer()
	})
}

type delItem struct {
	Executor
}

func (d *delItem) Less(other prioqueue.Item) bool {
	otherItem := other.(*delItem)
	return d.Expire().Before(otherItem.Expire())
}

type Executor interface {
	Key() string
	Wait()
	Result() (interface{}, error)
	CopyResult(Executor)
	Do()
	Close()
	Expire() time.Time
}
