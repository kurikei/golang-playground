package main

import "fmt"

func panicRecover() {
	if r := recover(); r != nil {
		fmt.Printf("recovered by panic(%v)", r)
	}
}

func main() {
	defer panicRecover() // cannot recover panic in goroutine
	done := make(chan struct{})

	go func() {
		panic("panic occured!")
		close(done)
	}()
	<-done

	fmt.Println("done")
}
