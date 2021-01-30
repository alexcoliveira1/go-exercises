package main

import (
	"testing"
)

func TestIntMinDiff(t *testing.T) {
	v := IntMin(0, 1)
	if v != 0 {
		t.Fail()
		t.Log(
			"For", 0, 1,
			"expected", 0,
			"got", v,
		)
	}
	v = IntMin(2, 1)
	if v != 1 {
		t.Fail()
		t.Log(
			"For", 2, 1,
			"expected", 1,
			"got", v,
		)
	}
}

func TestIntMinSame(t *testing.T) {
	v := IntMin(1, 1)
	if v != 1 {
		t.Fail()
		t.Log(
			"For", 1, 1,
			"expected", 1,
			"got", v,
		)
	}
}
