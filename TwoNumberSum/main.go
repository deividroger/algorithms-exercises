package main

func TwoNumberSum(array []int, target int) []int {
	results := []int{}

	if len(array) == 1 {
		return results
	}

	for index, currentValue := range array {
		for _, currentValue2 := range array[index+1:] {
			if currentValue+currentValue2 == target {
				results = append(results, currentValue, currentValue2)
				return results
			}
		}
	}

	return results
}
