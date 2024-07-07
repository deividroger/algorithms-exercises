package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TwoNumberSumTest(t *testing.T) {
	expected := []int{-1, 11}
	output := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)

	require.ElementsMatch(t, expected, output)
}
