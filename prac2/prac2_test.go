package main

import (
	"reflect"
	"testing"
)

var (
	a      = []int{1, 2, 3, 4, 5, 6, 7}
	b      = []int{2, 3, 5}
	wanted = []int{1, 4, 6, 7}
)

func TestFunc(t *testing.T) {
	got := Func(a, b)
	if !reflect.DeepEqual(got, wanted) {
		t.Fail()
	}
	t.Logf("got: %+v", got)
}
