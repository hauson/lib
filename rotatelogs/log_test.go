package rotatelogs

import (
	"time"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInitLogFile(t *testing.T) {
	if err := InitLogFile(""); err != nil {
		panic(err)
	}

	f := func() {
		for i := 0; i < 1000; i++ {
			logrus.WithFields(logrus.Fields{
				"animal": "walrus",
				"number": i,
			}).Info("A walrus appears")
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}

	for i := 0; i < 10; i++ {
		go f()
	}

	time.Sleep(11 * time.Minute)
}
