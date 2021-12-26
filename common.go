package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(fileName string) []string {
	println()
	println(fileName)
	b, err := ioutil.ReadFile("inputs/" + fileName)
	check(err)

	var result []string

	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result
}

func parseInt(value string) int {
	result, err := strconv.Atoi(value)
	check(err)
	return result
}

func parseInts(data string, separator string) []int {
	parts := strings.Split(data, separator)
	nums := make([]int, 0, len(parts))

	for _, part := range parts {
		if len(part) > 0 {
			nums = append(nums, parseInt(part))
		}
	}

	return nums
}

func parseInts2(lines []string, separator string) (result [][]int) {
	for _, line := range lines {
		result = append(result, parseInts(line, ""))
	}
	return result
}

func printInts(ints []int) {
	for _, v := range ints {
		print(v)
		print(" ")
	}
	println()
}

func printInts2(ints [][]int) {
	for _, row := range ints {
		printInts(row)
	}
	println()
}

func sumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func intMax(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func intMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func intMax1(values []int) (result int) {
	result = values[0]
	for _, v := range values {
		result = intMax(result, v)
	}
	return
}

func intSign(a int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -1
	} else {
		return 1
	}
}

func intAbs(a int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -a
	} else {
		return a
	}
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func intFoldLeft(values []int, zero int, f func(a, b int) int) int {
	for _, value := range values {
		zero = f(zero, value)
	}
	return zero
}

func intClamp(value int, min int, max int) int {
	return intMin(intMax(value, min), max)
}
