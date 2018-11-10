package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// ポインタpの指す先にvの値を入れる関数。
// 不適切な値があったらエラーを返す。
// 変更するのはSetの内部だけ。
func Set(p, v interface{}) error {
	pv := reflect.ValueOf(p)
	if pv.Kind() != reflect.Ptr {
		return errors.New("p is not a pointer")
	}

	vv := reflect.ValueOf(v)
	if !reflect.ValueOf(p).Elem().CanSet() {
		return errors.New("cannot assign value to p")
	}
	if !reflect.TypeOf(v).AssignableTo(reflect.TypeOf(p).Elem()) {
		return errors.New("cannot assign v to p")
	}

	// Elemでポインタの指してる先を取得して代入
	pv.Elem().Set(vv)
	return nil

}

type myint int

func (n myint) String() string {
	return fmt.Sprint(int(n))
}

func TestGolangTokyoDevQuiz18(t *testing.T) {
	var (
		v1 int
		v2 fmt.Stringer
	)

	cases := map[string]struct {
		p, v      interface{}
		expectErr bool
	}{
		"ok":       {&v1, 100, false},
		"stringer": {&v2, myint(100), false},
		"myint":    {&v1, myint(100), true},
		"nil":      {nil, 100, true},
		"string":   {&v1, "hoge", true},
		"float64":  {&v1, 1.11, true},
		"func":     {&v1, func() {}, true},
	}

	for tn, tc := range cases {
		tc := tc
		t.Run(tn, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatal("panicが起きた:", r)
				}
			}()

			err := Set(tc.p, tc.v)
			switch {
			case tc.expectErr && err == nil:
				t.Fatal("期待したエラーが発生しませんでした")
			case !tc.expectErr && err != nil:
				t.Fatal("予想外のエラー:", err)
			case err == nil:
				v := reflect.ValueOf(tc.p).Elem().Interface()
				if !reflect.DeepEqual(v, tc.v) {
					t.Errorf("want %v but got %v", tc.v, v)
				}
			}
		})
	}
}
