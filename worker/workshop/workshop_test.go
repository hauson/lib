package workshop

import (
	"testing"
	"strconv"
	"time"
	"math/rand"

	"github.com/hauson/lib/worker/mockjob"
	"github.com/hauson/lib/worker/sequence"
)

func TestWorkshop_AddJob(t *testing.T) {
	rand.Seed(time.Now().Unix())

	seqMaker := sequence.New(1000000)
	shop := New(3)
	for i := 0; i < 40; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		n := rand.Intn(5)
		job := mockjob.New("job"+strconv.Itoa(n), seqMaker.Next())
		shop.AddJob(job)
	}

	time.Sleep(2 * time.Minute)
}
