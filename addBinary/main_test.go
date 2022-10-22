package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddBinary(t *testing.T) {
	c := require.New(t)

	arrTest := []struct {
		p1  string
		p2  string
		out string
	}{
		{
			p1:  "11",
			p2:  "1",
			out: "100",
		},
		{
			p1:  "1010",
			p2:  "1011",
			out: "10101",
		},
	}

	for _, e := range arrTest {
		c.Equal(e.out, addBinary(e.p1, e.p2))
	}
}
