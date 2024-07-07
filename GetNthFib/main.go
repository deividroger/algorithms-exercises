package main

func GetNthFib(n int) int {
	previous := 0
	current := 1

	for i := 0; i < n-1; i++ {
		previous, current = current, previous+current
	}
	return previous
}
