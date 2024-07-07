package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {
	actualArray := []int{8, 5, 2, 9, 5, 6, 3}
	expectedResult := []int{2, 3, 5, 5, 6, 8, 9}
	result := QuickSort(actualArray)

	require.ElementsMatch(t, expectedResult, result)

}
