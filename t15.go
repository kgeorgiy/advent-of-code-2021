package main

func solve15(fileName string, resize int) int {
	base := parseInts2(readLines(fileName), "")

	weights := make([][]int, len(base)*resize)
	for r := range weights {
		weights[r] = make([]int, len(base[0])*resize)
		for c := range weights[r] {
			weights[r][c] = (base[r%len(base)][c%len(base[0])]+r/len(base)+c/len(base[0])-1)%9 + 1
		}
	}

	distances := make([][]int, len(weights))
	for i := range distances {
		distances[i] = make([]int, len(weights[i]))
	}

	distances[0][0] = 1
	for {
		best := 100000000
		bestR := 0
		bestC := 0
		for r, row := range distances {
			for c, value := range row {
				if value > 0 && best > value {
					best = value
					bestR = r
					bestC = c
				}
			}
		}
		if best == 100000000 {
			return -1 - distances[len(distances)-1][len(distances[0])-1]
		}
		//println(bestR, bestC, best)

		relaxDistances(distances, weights, best, bestR+1, bestC)
		relaxDistances(distances, weights, best, bestR-1, bestC)
		relaxDistances(distances, weights, best, bestR, bestC+1)
		relaxDistances(distances, weights, best, bestR, bestC-1)
		distances[bestR][bestC] = -distances[bestR][bestC]
	}
}

func relaxDistances(distances [][]int, weights [][]int, best int, r int, c int) {
	if 0 <= r && r < len(distances) && 0 <= c && c < len(distances[r]) {
		value := distances[r][c]
		if value == 0 || value > best+weights[r][c] {
			distances[r][c] = best + weights[r][c]
		}
	}
}

func main() {
	println(solve15("15_1.in", 1))
	println(solve15("15_2.in", 1))
	println(solve15("15_1.in", 5))
	println(solve15("15_2.in", 5))
}
