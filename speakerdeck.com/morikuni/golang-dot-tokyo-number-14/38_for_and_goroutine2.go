package main

import (
	"fmt"
	"sync"
)

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=38
func main() {
	xs := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	wg.Add(len(xs))

	for _, x := range xs {
		go func(x int) {
			defer wg.Done()

			fmt.Println(x) // 1から5までの値が順不同で出力される
		}(x)
	}
	wg.Wait()
}
