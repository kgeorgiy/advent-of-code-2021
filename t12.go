package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func caveDfs(v string, edges map[string][]string, visits map[string]int, twice bool) int {
	if v == "end" {
		return 1
	}
	r, _ := utf8.DecodeRuneInString(v)
	//visits[v] == 0 || twice && visits
	if !unicode.IsUpper(r) {
		if visits[v] > 0 {
			if v == "start" || !twice {
				return 0
			}
			twice = false
		}
	}
	visits[v]++

	total := 0
	for _, u := range edges[v] {
		total += caveDfs(u, edges, visits, twice)
	}

	visits[v]--
	return total
}

func solve12(fileName string, twice bool) int {
	edges := map[string][]string{}
	for _, line := range readLines(fileName) {
		parts := strings.Split(line, "-")
		edges[parts[0]] = append(edges[parts[0]], parts[1])
		edges[parts[1]] = append(edges[parts[1]], parts[0])
	}
	return caveDfs("start", edges, map[string]int{}, twice)
}

func main() {
	println(solve12("12_1.in", false))
	println(solve12("12_2.in", false))
	println(solve12("12_3.in", false))
	println(solve12("12_4.in", false))
	println(solve12("12_1.in", true))
	println(solve12("12_2.in", true))
	println(solve12("12_3.in", true))
	println(solve12("12_4.in", true))
}
