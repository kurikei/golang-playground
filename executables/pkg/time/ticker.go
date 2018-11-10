package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	c := make(chan time.Time)
	go simpleTicker(c)

	for {
		select {
		case now := <-t.C:
			fmt.Printf("time.Ticker: %#v\n", now.Format(time.RFC3339))
		case now := <-c:
			fmt.Printf("simpleTicker: %#v\n", now.Format(time.RFC3339))
		}
	}
}

func simpleTicker(c chan time.Time) {
	for {
		c <- time.Now()
		time.Sleep(time.Second)
	}
}
