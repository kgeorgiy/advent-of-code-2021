package main

import "strings"

type bits struct {
	bits []int
	pos  int
}

func solve16(fileName string, eval bool) int {
	hex := make([][]int, 16)
	for i := 0; i < 16; i++ {
		hex[i] = []int{i / 8, i / 4 % 2, i / 2 % 2, i % 2}
	}

	indent := 0
	for _, line := range readLines(fileName) {
		println(line)
		inputBits := make([]int, len(line)*4)
		for i, c := range line {
			var v int32
			if c >= 'A' {
				v = c - 'A' + 10
			} else {
				v = c - '0'
			}
			for j := 0; j < 4; j++ {
				inputBits[i*4+j] = hex[v][j]
			}
		}
		input := &bits{inputBits, 0}

		println(readPacket(input, indent, eval))
		println()
	}
	return 0
}

func boolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

func readPacket(input *bits, indent int, eval bool) int {
	versions := getInt(input, 3)
	id := getInt(input, 3)
	value := 0
	if id == 4 {
		for {
			first := getInt(input, 1)
			value = value*16 + getInt(input, 4)
			if first == 0 {
				break
			}
		}
		println(strings.Repeat("    ", indent), versions, id, "=", value)
		if eval {
			return value
		}
	} else {
		println(strings.Repeat("    ", indent), versions, id, "=")
		var args []int
		if getInt(input, 1) == 0 {
			length := getInt(input, 15)
			limit := (*input).pos + length
			for (*input).pos < limit {
				args = append(args, readPacket(input, indent+1, eval))
			}
		} else {
			count := getInt(input, 11)
			for i := 0; i < count; i++ {
				args = append(args, readPacket(input, indent+1, eval))
			}
		}
		if eval {

			if id == 0 {
				return intFoldLeft(args, 0, func(a, b int) int { return a + b })
			} else if id == 1 {
				return intFoldLeft(args, 1, func(a, b int) int { return a * b })
			} else if id == 2 {
				return intFoldLeft(args, args[0], intMin)
			} else if id == 3 {
				return intFoldLeft(args, args[0], intMax)
			} else if id == 5 {
				return boolToInt(args[0] > args[1])
			} else if id == 6 {
				return boolToInt(args[0] < args[1])
			} else if id == 7 {
				return boolToInt(args[0] == args[1])
			}
		} else {
			for _, arg := range args {
				versions += arg
			}
		}
	}
	return versions
}

func getInt(bits *bits, count int) (result int) {
	for i := 0; i < count; i++ {
		result = (result << 1) + bits.bits[(*bits).pos+i]
	}
	(*bits).pos += count
	return result
}

func main() {
	println(solve16("16.in", false))
	println(solve16("16.in", true))
}
