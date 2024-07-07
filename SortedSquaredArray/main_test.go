package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func SortedSquaredArrayTest(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expectedResult := []int{1, 4, 9, 16, 25, 36, 49, 64, 91}

	require.Equal(t, expectedResult, array)

}
