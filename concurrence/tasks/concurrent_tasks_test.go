package tasks

import (
	"testing"
	"time"

	"github.com/bytom/blockcenter/test/suit"
)

func TestMultiTask(t *testing.T) {
	type Args struct {
		TaskNum    int
		SpendTimes []time.Duration
		OutTime    time.Duration
	}
	suit.TestSuit{
		{
			Desc: "",
			Args: Args{
				TaskNum:    3,
				SpendTimes: []time.Duration{10 * time.Millisecond, 15 * time.Millisecond, 25 * time.Millisecond},
				OutTime:    20 * time.Millisecond,
			},
			WantResults:  2,
			WantErr:      "",
			IgnoreFileds: []string{},
		},
	}.Range(t, func(c *suit.TestCase) (interface{}, error) {
		args := c.Args.(Args)
		tasks := make([]TaskAble, args.TaskNum)
		for i := 0; i < args.TaskNum; i++ {
			tasks[i] = NewMockRunner(args.SpendTimes[i])
		}

		results := ConcurrentTasks(tasks, args.OutTime)
		var sucessTaskCnt int
		for _, result := range results {
			if result.Error == nil {
				sucessTaskCnt++
			}
		}
		return sucessTaskCnt, nil
	})
}

// ----------------mock ------------------------------------
type TimeTask struct {
	ch chan TaskResult
	du time.Duration
}

func NewMockRunner(du time.Duration) *TimeTask {
	return &TimeTask{
		ch: make(chan TaskResult),
		du: du,
	}
}

func (r *TimeTask) Exec() {
	time.Sleep(r.du)
	r.ch <- TaskResult{
		Result: 1,
		Type:   ResultInt,
	}
}

func (r *TimeTask) Name() string {
	return "mock task"
}

func (r *TimeTask) Result() <-chan TaskResult {
	return r.ch
}
