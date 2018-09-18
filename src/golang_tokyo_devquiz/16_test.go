package main

import (
	"fmt"
	"testing"
)

type Hex int

// テストが通るように以下のメソッドを修正させてください
func (h Hex) String() string {
	return fmt.Sprintf("%x", h)
}

// テストは変更しないでください
func Test(t *testing.T) {
	cases := map[string]struct {
		h Hex
		e string
	}{
		"100->64": {h: 100, e: "64"},
		"16->10":  {h: 16, e: "10"},
	}
	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			if s := tc.h.String(); s != tc.e {
				t.Errorf("want %q, got %q", tc.e, s)
			}
		})
	}
}
