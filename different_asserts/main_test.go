package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumPositiveNum(t *testing.T) {
	v, _ := SumPositiveNum([]int{1, 2})
	assert.Equal(t, v, 2, "they should be equal")
}
