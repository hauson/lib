package client

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := New("127.0.0.1", "3000")
	if err := client.Connect(); err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	go client.ReadLoop()

	if err := client.reqLogin("chx"); err != nil {
		panic(err)
	}

	if err := client.reqSubscribe("chain_status"); err != nil {
		panic(err)
	}

	go client.HeartbeatLoop(60 * time.Second)
	<-make(chan int)
}
