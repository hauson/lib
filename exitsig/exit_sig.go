package exitsig

import "sync"

type ExitSig struct {
	wg       sync.WaitGroup
	ch       chan int
	callback func() error
}

func New(f func() error) *ExitSig {
	return &ExitSig{
		ch:       make(chan int),
		callback: f,
	}
}

type Sig <-chan int
type GoWithExitSig func(exitSig Sig)

func (s *ExitSig) GoFunc(g GoWithExitSig) {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()
		g(s.ch)
	}()
}

func (s *ExitSig) IsClose() bool {
	select {
	case <-s.ch:
		return true
	default:
		return false
	}
}

func (s *ExitSig) Close() error {
	select {
	case <-s.ch:
		return nil
	default:
		close(s.ch)
		s.wg.Wait()
		if s.callback == nil {
			return nil
		}
		return s.callback()
	}
}
