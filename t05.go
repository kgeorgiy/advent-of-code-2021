package main

import (
	"strings"
)

type vent struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func solve05(fileName string, diags bool) int {
	var vents []vent = nil
	for _, line := range readLines(fileName) {
		parts := strings.Split(line, " -> ")
		coord0 := parseInts(parts[0], ",")
		coord1 := parseInts(parts[1], ",")
		vents = append(vents, vent{coord0[0], coord0[1], coord1[0], coord1[1]})
	}

	maxX := 0
	maxY := 0
	for _, vent := range vents {
		maxX = intMax(intMax(maxX, vent.x1), vent.x2)
		maxY = intMax(intMax(maxY, vent.y1), vent.y2)
	}

	field := make([][]int, maxX+1)
	for i := range field {
		field[i] = make([]int, maxY+1)
	}

	for _, vent := range vents {
		dx := intSign(vent.x2 - vent.x1)
		dy := intSign(vent.y2 - vent.y1)
		if diags || dx == 0 || dy == 0 {
			for x, y := vent.x1, vent.y1; x != vent.x2 || y != vent.y2; x, y = x+dx, y+dy {
				field[x][y]++
			}
			field[vent.x2][vent.y2]++
		}
	}

	count := 0
	for _, row := range field {
		for _, value := range row {
			if value >= 2 {
				count++
			}
		}
	}
	return count
}

func main() {
	println(solve05("05_1.in", false))
	println(solve05("05_2.in", false))
	println(solve05("05_1.in", true))
	println(solve05("05_2.in", true))
}
