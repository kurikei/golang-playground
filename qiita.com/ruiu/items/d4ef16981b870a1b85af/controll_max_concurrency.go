package main

import (
	"fmt"
	"time"
)

func main() {
	maxConcurrencySize := 3
	c := make(chan bool, maxConcurrencySize)

	for i := 0; i < 30; i++ {
		c <- true // 最大並行数を超えたらここで待たされる

		go func(i int) {
			defer func() { <-c }()

			fmt.Printf("process index: %v\n", i)
			time.Sleep(time.Second)
		}(i)
	}
}
