package main

import (
	"fmt"
	"sync"
)

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=38
func main() {
	// creating huge slice
	xs := []int{1, 2, 3, 4, 5}
	}

	var wg sync.WaitGroup
	wg.Add(len(xs))

	for _, x := range xs {
		x := x // 別の変数に代入する
		go func() {
			defer wg.Done()

			fmt.Println(x) // 1から5までの値が順不同で出力される
		}()
	}
	wg.Wait()
}
