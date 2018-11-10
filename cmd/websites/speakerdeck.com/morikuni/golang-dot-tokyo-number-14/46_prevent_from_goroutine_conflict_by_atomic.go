package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	sizeOfSlice := 100
	xs := make([]int, sizeOfSlice)
	for i := 0; i < sizeOfSlice; i++ {
		xs[i] = i
	}

	var wg sync.WaitGroup
	wg.Add(sizeOfSlice)

	var sum *int64
	for x := range xs {
		// sum += int64(xs[x])
		go func(i int) {
			defer wg.Done()

			// これだと NG
			// sum += int64(xs[i])
			// これも NG
			atomic.AddInt64(sum, int64(xs[i]))
		}(x)
	}

	wg.Done()
	fmt.Println(sum)
}
