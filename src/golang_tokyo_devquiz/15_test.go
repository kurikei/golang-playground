package main

import "testing"

type MyInt int

// テストが通るように以下のメソッドを修正させてください
func (n MyInt) Inc() {
	n++
}

// テストは変更しないで下さい
func TestMyInt_Inc(t *testing.T) {
	var n MyInt
	n.Inc()
	if n != 1 {
		t.Errorf("want %d, got %d", 1, n)
	}
	n = 10
	n.Inc()
	if n != 11 {
		t.Errorf("want %d, got %d", 11, n)
	}
}
