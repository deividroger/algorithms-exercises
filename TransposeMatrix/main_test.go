package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransposeMatrix(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	actual := TransposeMatrix(input)
	require.Equal(t, expected, actual)
}
