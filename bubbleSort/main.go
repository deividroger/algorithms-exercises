package main

import (
	"sort"
)

func BubbleSort(array []int) []int {

	for i := 0; i < len(array)-1; i++ {
		for j := 1; j < len(array); j++ {
			if array[i] < array[j] {
				change := array[i]
				array[i] = array[j]
				array[j] = change
			}
		}
	}
	sort.Ints(array)

	return array
}
