package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	c := require.New(t)

	arrTest := []struct {
		strs []string
		out  string
	}{
		{
			strs: []string{""},
			out:  "",
		},
		{
			strs: []string{"a"},
			out:  "a",
		},
		{
			strs: []string{"ab", "a"},
			out:  "a",
		},
		{
			strs: []string{"racecar", "dog", "car"},
			out:  "",
		},
	}

	for _, e := range arrTest {
		c.Equal(e.out, longestCommonPrefix(e.strs))
	}
}
