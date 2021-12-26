package main

import (
	"strings"
)

func solve24(fileName string, min bool) int {
	vars := "wxyz"
	state := map[[4]int32]int{}
	state[[4]int32{}] = 0
	no := 0
	total := 1
	input := 0
	for _, line := range readLines(fileName) {
		parts := strings.Split(line, " ")

		switch command := parts[0]; command {
		case "inp":
			target := strings.Index(vars, parts[1])
			nextState := map[[4]int32]int{}
			input++
			var start, end int
			if input >= 2 {
				start = 1
				end = 9
			} else if min {
				start = 7
				end = 7
			} else {
				start = 9
				end = 9
			}
			total *= 10 - start
			for i := start; i <= end; i++ {
				ii := int32(i)
				for prev, best := range state {
					next := [4]int32{prev[0], prev[1], prev[2], prev[3]}
					next[target] = ii
					relax(min, nextState, next, best*10+i)
				}
			}
			state = nextState
		case "mul":
			state = binary(min, state, parts[1], parts[2], func(a int32, b int32) (int32, bool) { return a * b, true })
		case "add":
			state = binary(min, state, parts[1], parts[2], func(a int32, b int32) (int32, bool) { return a + b, true })
		case "mod":
			state = binary(min, state, parts[1], parts[2], func(a int32, b int32) (int32, bool) {
				if a < 0 || b <= 0 {
					return 0, false
				} else {
					return a % b, true
				}
			})
		case "div":
			state = binary(min, state, parts[1], parts[2], func(a int32, b int32) (int32, bool) {
				if b == 0 {
					return 0, false
				} else {
					return a / b, true
				}
			})
		case "eql":
			state = binary(min, state, parts[1], parts[2], func(a int32, b int32) (int32, bool) {
				if a == b {
					return 1, true
				} else {
					return 0, true
				}
			})
		default:
			println("Unknown command " + command)
			panic(nil)
		}
		//for values, best := range State {
		//	println(values[0], values[1], values[2], values[3], best)
		//}

		no++
		println(input, no, total, len(state), total/len(state))
	}

	result := -1
	for prev, best := range state {
		if prev[3] == 0 && best > result {
			result = best
		}
	}

	return result
}

func binary(min bool, state map[[4]int32]int, targetS string, sourceS string, f func(int32, int32) (int32, bool)) map[[4]int32]int {
	nextState := map[[4]int32]int{}
	target := strings.Index("wxyz", targetS)
	source := strings.Index("wxyz", sourceS)
	var ok bool
	if source >= 0 {
		for prev, best := range state {
			next := [4]int32{prev[0], prev[1], prev[2], prev[3]}
			next[target], ok = f(next[target], prev[source])
			if ok {
				relax(min, nextState, next, best)
			}
		}
	} else {
		source := int32(parseInt(sourceS))
		for prev, best := range state {
			next := [4]int32{prev[0], prev[1], prev[2], prev[3]}
			next[target], ok = f(next[target], source)
			if ok {
				relax(min, nextState, next, best)
			}
		}
	}
	return nextState
}

func relax(min bool, state map[[4]int32]int, next [4]int32, newValue int) {
	if value, ok := state[next]; !ok || ((value >= newValue) == min) {
		state[next] = newValue
	}
}

func main() {
	//println(solve24("24_1.in", false))
	println(solve24("24_1.in", true))
}
