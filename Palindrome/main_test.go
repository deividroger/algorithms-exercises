package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	output := IsPalindrome("abcdcba")

	require.Equal(t, true, output)
}

func TestIsNotPalindrome(t *testing.T) {
	output := IsPalindrome("aafabcdcba")

	require.Equal(t, false, output)

}
