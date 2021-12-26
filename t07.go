package main

func solve07(fileName string, squared bool) int {
	coords := parseInts(readLines(fileName)[0], ",")
	maxC := intMax1(coords)

	best := 1000000000
	for i := 0; i <= maxC; i++ {
		sum := 0
		for _, c := range coords {
			delta := intAbs(c - i)
			if squared {
				sum += delta * (delta + 1) / 2
			} else {
				sum += delta
			}
		}
		best = intMin(best, sum)
	}
	return best
}

func main() {
	println(solve07("07_1.in", false))
	println(solve07("07_2.in", false))
	println(solve07("07_1.in", true))
	println(solve07("07_2.in", true))
}
