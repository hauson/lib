package mockjob

import (
	"time"
	"fmt"
	"math/rand"
)

type Job struct {
	name string
	seq  uint64
}

func New(name string, seq uint64) *Job {
	return &Job{
		name: name,
		seq:  seq,
	}
}

func (j *Job) Exec() string {
	time.Sleep(3 * time.Millisecond)
	if rand.Intn(5) == 2 {
		panic("mock panic")
	}

	return j.name + " mock job exec" + fmt.Sprintf(" seq:%d", j.seq)
}

func (j *Job) Name() string {
	return j.name
}

func (j *Job) Seq() uint64 {
	return j.seq
}
