package main

import (
	"sort"
	"strings"
)

func digits() []string {
	return []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}
}

func solve08(fileName string, one bool) int {
	total := 0
	for _, line := range readLines(fileName) {
		parts := strings.Split(line, " | ")
		digs := restore(strings.Split(parts[0], " "), strings.Split(parts[1], " "))
		n := 0
		for _, d := range digs {
			if one && (d == 1 || d == 4 || d == 7 || d == 8) {
				total++
			} else {
				n = n*10 + d
			}
		}
		if !one {
			total += n
		}
	}
	return total
}

func restore(digs1 []string, number []string) (result []int) {
	fr := map[rune]rune{}

	var f rune
	counts1 := countSegments(digs1)
	for d, c := range counts1 {
		if c == 9 {
			f = d
			fr[d] = 'f'
		} else if c == 4 {
			fr[d] = 'e'
		}
	}

	digs2 := filterDigits(digs1, f)
	counts2 := countSegments(digs2)
	for d := range counts2 {
		if _, ok := fr[d]; !ok && counts1[d] == counts2[d] {
			fr[d] = 'b'
		}
	}

	fr[findSegment(digs2, 2, fr)] = 'c'
	a := findSegment(digs2, 3, fr)
	fr[a] = 'a'

	for d, c := range countSegments(filterDigits(digs2, a)) {
		if _, ok := fr[d]; !ok {
			if c == 5 {
				fr[d] = 'd'
			} else if c == 6 {
				fr[d] = 'g'
			}
		}
	}

	for _, n := range number {
		var runes []rune
		for _, d := range n {
			runes = append(runes, fr[d])
		}
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		s := string(runes)
		digs := digits()
		result = append(result, SliceIndex(len(digs), func(i int) bool { return digs[i] == s }))
	}
	return result
}

func findSegment(digits []string, length int, m map[rune]rune) rune {
	for _, digit := range digits {
		if len(digit) == length {
			for _, d := range digit {
				if _, ok := m[d]; !ok {
					return d
				}
			}
		}
	}
	panic(digits)
}

func filterDigits(digs []string, c rune) (result []string) {
	for _, digit := range digs {
		if strings.ContainsRune(digit, c) {
			result = append(result, digit)
		}
	}
	return result
}

func countSegments(digits []string) map[rune]int {
	result := map[rune]int{}
	for _, c := range "abcdefg" {
		count := 0
		for _, digit := range digits {
			if strings.ContainsRune(digit, c) {
				count++
			}
		}
		result[c] = count
	}
	return result
}

func main() {
	println(solve08("08_1.in", true))
	println(solve08("08_2.in", true))
	println(solve08("08_1.in", false))
	println(solve08("08_2.in", false))
}
