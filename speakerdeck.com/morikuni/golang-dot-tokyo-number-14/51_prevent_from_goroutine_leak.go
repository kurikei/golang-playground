package main

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// https://speakerdeck.com/morikuni/golang-dot-tokyo-number-14?slide=51
func printGoroutineNum() {
	fmt.Printf("%v\n", runtime.NumGoroutine())
}

func fetch(ctx context.Context, url string, ch chan<- *http.Response) {
	res, _ := http.Get(url)
	select {
	case ch <- res:
	case <-ctx.Done(): // contextがdoneになったときに処理をやめる
	}
}

func main() {
	defer func() { // 最終的なgoroutineの数を出力する
		time.Sleep(time.Second) // cancelが完了するのを雑に待つ
		printGoroutineNum()     // cancelされたgoroutineの数(=2)だけ減る
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 完了後にcancelを実行し、contextをdone状態にする

	printGoroutineNum()

	ch := make(chan *http.Response)
	go fetch(ctx, "http://example.com", ch)
	printGoroutineNum() // increase goroutine num

	go fetch(ctx, "http://example.com", ch)
	printGoroutineNum() // increase goroutine num

	go fetch(ctx, "http://example.com", ch)
	printGoroutineNum() // increase goroutine num

	<-ch
	printGoroutineNum() // TODO: ここでもgoroutine の数が増える
}
