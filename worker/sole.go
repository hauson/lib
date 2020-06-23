package worker

import "sync"

type Sole struct {
	stop  func()
	mutex sync.Mutex
}

func (o *Sole) Close() {
	if o.stop != nil {
		o.stop()
		o.stop = nil
	}
}

func (o *Sole) Start(runner Runner) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.Close()
	go runner.Run()
	o.stop = runner.Stop
}

type Runner interface {
	Run()
	Stop()
}
