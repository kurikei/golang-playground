package main

import (
	"fmt"
	"sync"
)

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=31
func main() {
	var wg sync.WaitGroup
	wg.Add(1) // カウンタをセット
	go func() {
		defer wg.Done() // Done = Add(-1) する
		fmt.Println("wait goroutine by using WaitGroup")
	}()
	wg.Wait() // カウンタが0になるまでWait
}
