package main

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.

	winners := make(map[string]int)

	for index, competidors := range competitions { //position 0: home team, position: 1 away team
		if results[index] == 0 { //away team won
			addOrUpdatePostionInMap(winners, competidors[1], 3)

		} else if results[index] == 1 { //home team won

			addOrUpdatePostionInMap(winners, competidors[0], 3)
		}
	}

	return getWinner(winners)
}

func getWinner(winners map[string]int) string {
	var maxKey string
	valueValue := 0

	for key, value := range winners {
		if value > valueValue {
			valueValue = value
			maxKey = key
		}
	}
	return maxKey
}

func addOrUpdatePostionInMap(winners map[string]int, key string, pontuation int) {
	if val, ok := winners[key]; ok {
		winners[key] = val + pontuation
	} else {
		winners[key] = pontuation
	}
}
