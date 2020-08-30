package main

import (
	"fmt"
)

func LargestRange(array []int) []int {

	bestRange := []int{}
	longestLength := 0
	visited := map[int]bool{}
	for _, num := range array {
		visited[num] = true
	}
	for _, num := range array {
		if !visited[num] {
			continue
		}
		visited[num] = false
		currentLengnth := 1
		leftNum := num - 1
		rightNum := num + 1

		for visited[leftNum] {
			visited[leftNum] = false
			currentLengnth++
			leftNum--
		}
		for visited[rightNum] {
			visited[rightNum] = false
			currentLengnth += 1
			rightNum++
		}
		if currentLengnth > longestLength {
			longestLength = currentLengnth
			bestRange = []int{leftNum + 1, rightNum - 1}
		}
	}
	return bestRange
}

func main() {
	res := LargestRange([]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6, 8})
	fmt.Print(res)
}
