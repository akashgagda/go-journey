package main

func sumAll(num []int) int {
	var finalSum int
	for _, v := range num {
		finalSum += v
	}
	return finalSum
}
