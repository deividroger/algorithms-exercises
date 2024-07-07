package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func IsValidSequenceTest(t *testing.T) {

	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	require.True(t, IsValidSubSequece(array, sequence))
}
