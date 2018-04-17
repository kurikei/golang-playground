package main

import "fmt"

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=28
func main() {
	go func() {
		fmt.Println("cannot wait for printing")
	}()
}
