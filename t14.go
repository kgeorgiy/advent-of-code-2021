package main

import (
	"sort"
	"strings"
)

func solve14(fileName string, steps int) int {
	lines := readLines(fileName)

	pairs := map[string]int{}
	poly := lines[0]
	for i := range poly[1:] {
		pairs[poly[i:i+2]]++
	}

	var rules [][]string
	for _, line := range lines[1:] {
		parts := strings.Split(line, " -> ")
		rule1 := []string{parts[0], parts[0][:1] + parts[1]}
		rule2 := []string{parts[0], parts[1] + parts[0][1:]}
		rules = append(rules, rule1, rule2)
	}

	for i := 0; i < steps; i++ {
		next := map[string]int{}
		used := map[string]bool{}
		for _, rule := range rules {
			next[rule[1]] += pairs[rule[0]]
			used[rule[0]] = true
		}
		for pair, c := range pairs {
			if !used[pair] {
				next[pair] += c
			}
		}
		pairs = next
		total := 0
		for _, c := range pairs {
			total += c
		}
	}
	elements := map[uint8]int{poly[0]: 1, poly[len(poly)-1]: 1}
	for pair, c := range pairs {
		elements[pair[0]] += c
		elements[pair[1]] += c
	}

	var counts []int
	for _, c := range elements {
		counts = append(counts, c/2)
	}
	sort.Ints(counts)
	return counts[len(counts)-1] - counts[0]
}

func main() {
	println(solve14("14_1.in", 10))
	println(solve14("14_2.in", 10))
	println(solve14("14_1.in", 40))
	println(solve14("14_2.in", 40))
}
