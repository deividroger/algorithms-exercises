package main

import "sort"

func SortedSquaredArray(array []int) []int {
	result := []int{}

	for _, v := range array {
		result = append(result, v*v)
	}

	sort.Ints(result)

	return result
}
