package main

import (
	"time"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/hauson/lib/ipsniff"
)

func main() {
	sniff, err := ipsniff.New()
	if err != nil {
		log.Fatal(" new sniff err")
	}

	sniff.Run()

	for {
		time.Sleep(10 * time.Second)
		for _, ip := range sniff.IPList() {
			fmt.Println(ip)
		}
	}
}
