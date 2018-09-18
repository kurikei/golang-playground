package main

import (
	"fmt"
	"testing"
)

type ErrNotEven struct {
	N int
}

func (err *ErrNotEven) Error() string {
	return fmt.Sprintf("%d is not even", err.N)
}

// IsAllEvenだけを編集し、テストが通るようにしてください
func IsAllEven(ns ...int) error {
	var err *ErrNotEven
	for n := range ns {
		if n%2 != 0 {
			err = &ErrNotEven{N: n}
			break
		}
	}
	return err
}

func Test(t *testing.T) {
	cases := map[string]struct {
		ns       []int
		hasError bool
	}{
		"1,2,3": {ns: []int{0, 1, 2}, hasError: true},
		"1,3,5": {ns: []int{1, 3, 5}, hasError: true},
		"0,2,4": {ns: []int{2, 4, 6}, hasError: false},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			err := IsAllEven(tc.ns...)
			switch {
			case err != nil && !tc.hasError:
				t.Error("unexpected error:", err)
			case err == nil && tc.hasError:
				t.Error("expected error has not occurred")
			}
		})
	}
}
