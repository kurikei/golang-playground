package main

import (
	"fmt"
	"sync"
)

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=34
func main() {
	// creating huge slice
	sizeOfSlice := 100
	xs := make([]int, sizeOfSlice)
	for i := 0; i < sizeOfSlice; i++ {
		xs[i] = i
	}

	var wg sync.WaitGroup
	wg.Add(len(xs))

	for _, x := range xs {
		go func() {
			defer wg.Done()

			fmt.Println(x) // 離散的な数値がprintされる
		}()
	}
	wg.Wait()
}
