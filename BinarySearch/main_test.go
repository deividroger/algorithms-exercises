package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BinarySearchTest(t *testing.T) {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 93}
	target := 31

	output := BinarySearch(array, target)

	require.Equal(t, target, output)
}
