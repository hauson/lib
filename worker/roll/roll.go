package roll

import "sync"

type JobRoll struct {
	jobNames map[string]int
	*sync.RWMutex
}

func New() *JobRoll {
	return &JobRoll{
		jobNames: make(map[string]int),
		RWMutex:  new(sync.RWMutex),
	}
}

func (r *JobRoll) Enroll(name string) {
	r.Lock()
	defer r.Unlock()

	r.jobNames[name]++
}

func (r *JobRoll) Unroll(name string) {
	r.Lock()
	defer r.Unlock()

	r.jobNames[name]--
	if r.jobNames[name] <= 0 {
		delete(r.jobNames, name)
	}
}

func (r *JobRoll) Exist(name string) bool {
	r.RLock()
	defer r.RUnlock()

	return r.jobNames[name] > 0
}
