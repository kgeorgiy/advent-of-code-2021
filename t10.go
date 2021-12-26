package main

import (
	"sort"
	"strings"
)

func solve10(fileName string, incomplete bool) int {
	penalties := []int{3, 57, 1197, 25137}
	total := 0
	var scores []int
	for _, line := range readLines(fileName) {
		var open []int
		corrupted := false
		for _, c := range line {
			index := strings.IndexRune("([{<)]}>", c)
			if index < 4 {
				open = append(open, index)
			} else {
				top := len(open) - 1
				if open[top] != index-4 {
					if !incomplete {
						total += penalties[index-4]
					}
					corrupted = true
					break
				} else {
					open = open[:len(open)-1]
				}
			}
		}
		if incomplete && !corrupted {
			score := 0
			for i := len(open) - 1; i >= 0; i-- {
				score = score*5 + open[i] + 1
			}
			scores = append(scores, score)
		}
	}

	if incomplete {
		sort.Ints(scores)
		return scores[len(scores)/2]
	} else {
		return total
	}
}

func main() {
	println(solve10("10_1.in", false))
	println(solve10("10_2.in", false))
	println(solve10("10_1.in", true))
	println(solve10("10_2.in", true))
}
