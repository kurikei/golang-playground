package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t1000 := time.NewTicker(time.Second)
	defer t1000.Stop()

	t10 := time.NewTicker(10 * time.Millisecond)
	defer t10.Stop()

	rand.Seed(int64(time.Now().Nanosecond()))

	var number int64
	for {
		select {
		case <-t10.C:
			fmt.Printf(".")
			if number > 950 {
				fmt.Println("\nhit!!!")
				return
			}
		case <-t1000.C:
			number = rand.Int63n(1000)
		}
	}
}
