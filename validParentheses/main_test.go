package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	c := require.New(t)

	arrTest := []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
	}
}
