package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func GetNthFibTest(t *testing.T) {
	expected := 5
	output := GetNthFib(6)

	require.Equal(t, expected, output)
}
