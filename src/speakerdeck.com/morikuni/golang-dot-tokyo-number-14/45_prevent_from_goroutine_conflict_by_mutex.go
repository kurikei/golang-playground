package main

import (
	"fmt"
	"sync"
)

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=45
func main() {
	sizeOfSlice := 20
	xs := make([]int, sizeOfSlice)
	for i := 0; i < sizeOfSlice; i++ {
		xs[i] = i
	}

	var wg sync.WaitGroup
	wg.Add(len(xs))

	var mu sync.Mutex

	var double []int
	for _, x := range xs {
		go func(x int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()

			double = append(double, x*2)
		}(x)
	}

	wg.Wait()

	fmt.Printf("size of slice is %v, elements=%v", len(double), double) // 要素数が必ず20になる
}
