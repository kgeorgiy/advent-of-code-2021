package main

import "sort"

func readHeights(fileName string) [][]int {
	return parseInts2(readLines(fileName), "")
}

func getHeight(heights [][]int, r int, c int) int {
	if 0 <= r && r < len(heights) && 0 <= c && c < len(heights[r]) {
		return heights[r][c]
	} else {
		return 1_000_000
	}
}

func solve09one(fileName string) int {
	heights := readHeights(fileName)

	total := 0
	for r := range heights {
		for c := range heights[r] {
			min := intMin(
				intMin(getHeight(heights, r+1, c), getHeight(heights, r-1, c)),
				intMin(getHeight(heights, r, c+1), getHeight(heights, r, c-1)),
			)
			if heights[r][c] < min {
				//println(r, c)
				total += heights[r][c] + 1
			}
		}
	}
	//printInts2(heights)
	return total
}

func solve09two(fileName string) int {
	heights := readHeights(fileName)

	var sizes []int
	for r := range heights {
		for c := range heights[r] {
			if heights[r][c] != 9 {
				size := fill(heights, r, c)
				sizes = append(sizes, size)
			}
		}
	}

	sort.Ints(sizes)
	//printInts2(heights)
	n := len(sizes)
	return sizes[n-1] * sizes[n-2] * sizes[n-3]
}

func fill(heights [][]int, r int, c int) int {
	if getHeight(heights, r, c) < 9 {
		heights[r][c] = 9
		return 1 + fill(heights, r+1, c) + fill(heights, r-1, c) + fill(heights, r, c+1) + fill(heights, r, c-1)
	} else {
		return 0
	}
}

func main() {
	println(solve09one("09_1.in"))
	println(solve09one("09_2.in"))
	println(solve09two("09_1.in"))
	println(solve09two("09_2.in"))
}
