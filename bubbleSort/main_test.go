package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubleSortTest(t *testing.T) {

	array := []int{8, 5, 2, 9, 5, 6, 3}
	expectSequence := []int{2, 3, 5, 5, 6, 8, 9}
	result := BubbleSort(array)

	require.ElementsMatch(t, result, expectSequence)
}
