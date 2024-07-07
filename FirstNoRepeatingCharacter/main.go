package main

import (
	"strings"
)

func FirstNonRepeatingCharacter(str string) int {

	for _, v := range str {
		result := strings.Count(str, string(v))

		if result == 1 {
			return strings.Index(str, string(v))
		}

	}
	return -1

}
