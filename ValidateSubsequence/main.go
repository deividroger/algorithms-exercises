package main

func IsValidSubSequece(array []int, sequence []int) bool {

	i := 0

	for _, v := range array {
		if v == sequence[i] {
			i++
		}
		if len(sequence) == i {
			return true
		}
	}

	return false
}
