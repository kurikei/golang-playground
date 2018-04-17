package main

import (
	"fmt"
	"net/http"
	"runtime"
)

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=50
type result struct {
	Response *http.Response
	Err      error
}

func fetch(url string, c chan<- result) {
	res, err := http.Get(url)
	c <- result{Response: res, Err: err}
}

func printGoroutineNum() {
	fmt.Printf("%v\n", runtime.NumGoroutine())
}

func main() {
	printGoroutineNum()
	ch := make(chan result)

	go fetch("http://example.com", ch)
	printGoroutineNum() // increase goroutine

	go fetch("http://example.com", ch)
	printGoroutineNum() // increase goroutine

	fmt.Println(<-ch)
	printGoroutineNum()

	fmt.Println(<-ch)
	printGoroutineNum()

	fmt.Println(<-ch)
	printGoroutineNum()
	// goroutine leak occured
	// 参照：https://qiita.com/i_yudai/items/3336a503079ac5749c35
}
