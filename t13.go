package main

import "strings"

func solve13(fileName string) int {
	lines := readLines(fileName)

	var marks [][]int
	folds := false
	for _, line := range lines {
		if line == "---" {
			folds = true
		} else if folds {
			parts := strings.Split(line, "=")
			if parts[0] == "fold along x" {
				marks = fold(marks, parseInt(parts[1]), 1000000)
			} else {
				marks = fold(marks, 100000, parseInt(parts[1]))
			}
			println(len(marks))
		} else {
			marks = append(marks, parseInts(line, ","))
		}
	}
	var output [][]rune
	for _, mark := range marks {
		for r := len(output); r <= mark[1]; r++ {
			output = append(output, nil)
		}
		r := mark[1]
		for c := len(output[r]); c <= mark[0]; c++ {
			output[r] = append(output[r], ' ')
		}
		output[r][mark[0]] = '#'
	}
	for _, row := range output {
		println(string(row))
	}
	return -1
}

func fold(marks [][]int, x int, y int) (result [][]int) {
	visited := map[int]map[int]bool{}
	for _, mark := range marks {
		if mark[0] > x {
			mark[0] = 2*x - mark[0]
		}
		if mark[1] > y {
			mark[1] = 2*y - mark[1]
		}
		if !visited[mark[0]][mark[1]] {
			if visited[mark[0]] == nil {
				visited[mark[0]] = map[int]bool{}
			}
			visited[mark[0]][mark[1]] = true
			result = append(result, mark)
		}
	}
	return result
}

func main() {
	println(solve13("13_1.in"))
	println(solve13("13_2.in"))
}
