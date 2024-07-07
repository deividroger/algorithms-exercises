package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstNonRepeatingCharacter(t *testing.T) {
	input := "abcdcaf"
	expected := 1
	actual := FirstNonRepeatingCharacter(input)
	require.Equal(t, expected, actual)
}
