package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	c := require.New(t)
	arrOut := []struct {
		nums   []int
		target int
		out    []int
	}{
		{
			nums:   []int{2, 7, 11, 17},
			target: 9,
			out:    []int{0, 1},
		}, {
			nums:   []int{3, 2, 4},
			target: 6,
			out:    []int{1, 2},
		},
	}
}
