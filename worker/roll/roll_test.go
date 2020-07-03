package roll

import (
	"testing"
)

func TestJobRoll(t *testing.T) {
	tests := []struct {
		desc       string
		names      []string
		deleteName string
		findName   string
		jobRoll    *JobRoll
		want       bool
	}{
		{
			desc:     "test case1",
			names:    []string{"job1", "job2", "job3", "job4", "job5"},
			findName: "job3",
			jobRoll:  New(),
			want:     true,
		},
		{
			desc:       "test case2",
			names:      []string{"job1", "job2", "job3", "job4", "job5"},
			deleteName: "job3",
			findName:   "job3",
			jobRoll:    New(),
			want:       false,
		},
	}

	for _, tt := range tests {
		for _, name := range tt.names {
			tt.jobRoll.Enroll(name)
		}
		tt.jobRoll.Unroll(tt.deleteName)
		if got := tt.jobRoll.Exist(tt.findName); got != tt.want {
			t.Fatalf("case: %s got:%v, want:%v", tt.desc, got, tt.want)
		}
	}
}
