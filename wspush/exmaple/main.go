package main

import (
	"fmt"
	"time"

	"github.com/lib/wspush/configs"
	"github.com/lib/exception"
	"github.com/lib/wspush/pusher"
)

func main() {
	cfg, err := configs.Load("cfg.json")
	exception.CheckError(err)

	pusher := pusher.New(cfg)
	go pusher.Run()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for i := 0; true; <-ticker.C {
		i++
		acc := "chx"
		if i%2 == 0 {
			acc = "shanshi"
		}
		pusher.Send(&ChainStatus{
			Accounts: []string{acc},
			Height:   uint64(i),
		})
	}
	fmt.Println("start ws pusher sucess!")
	select {}
}
