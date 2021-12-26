package main

func solve17(fileName string, first bool) int {
	for _, line := range readLines(fileName) {
		coords := parseInts(line, " ")
		println(solve17Case(coords[0], coords[2], coords[1], coords[3], first))
	}
	return 0
}

func solve17Case(minX int, minY int, maxX int, maxY int, first bool) int {
	count := 0
	for dy := 1000; dy >= minY; dy-- {
		for dx := -1000; dx <= 1000; dx++ {
			x := 0
			y := 0
			vx := dx
			vy := dy
			top := 0
			for y >= minY {
				top = intMax(top, y)
				if minX <= x && x <= maxX && minY <= y && y <= maxY {
					if first {
						return top
					}
					//println(dx, dy)
					count++
					break
				}
				x += vx
				y += vy
				vx -= intSign(vx)
				vy -= 1
			}
		}
	}
	return count
}

func main() {
	println(solve17("17.in", true))
	println(solve17("17.in", false))
}
