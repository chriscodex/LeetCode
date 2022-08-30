package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotal(t *testing.T) {
	c := require.New(t)
	testArr := []struct {
		str      string
		expected int
	}{
		{
			str:      "MDCXCV",
			expected: 1695,
		},
		{
			str:      "X",
			expected: 10,
		},
		{
			str:      "CMVI",
			expected: 906,
		},
		{
			str:      "XXII",
			expected: 22,
		},
		{
			str:      "XI",
			expected: 11,
		},
		{
			str:      "IX",
			expected: 9,
		},
		{
			str:      "CDLXXX",
			expected: 480,
		},
		{
			str:      "MCMXCIV",
			expected: 1994,
		},
		{
			str:      "XXXIV",
			expected: 34,
		},
	}
	for _, e := range testArr {
		result := romainToInt(e.str)
		c.Equal(e.expected, result)
	}
}
