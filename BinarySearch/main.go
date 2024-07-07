package main

func BinarySearch(array []int, target int) int {
	less := 0
	high := len(array) - 1

	for less <= high {
		middle := (less + high) / 2
		shoot := array[middle]

		if shoot == target {
			return shoot
		}

		if shoot > target {
			high = middle - 1
		} else {
			less = middle + 1
		}

	}

	return -1
}
