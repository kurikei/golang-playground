package main

import "fmt"

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=30
func main() {
	c := make(chan struct{})
	go func() {
		defer func() { c <- struct{}{} }() // struct{}{} の size は 0byte (unsafe.Sizeof(struct{}{})) なのでよく用いられる
		fmt.Println("wait goroutine by using chan")
	}()
	<-c
}
