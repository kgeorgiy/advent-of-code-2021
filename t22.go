package main

import (
	"sort"
	"strings"
)

func parseCoordRange(input string) (int, int) {
	parts := strings.Split(strings.Split(input, "=")[1], "..")
	return parseInt(parts[0]), parseInt(parts[1])
}

const dims = 3

type cuboid struct {
	coords [dims][2]int
	on     bool
}

func solve22(fileName string, size int) int {
	cs := [dims][]int{}
	var cuboids []cuboid
	for _, line := range readLines(fileName) {
		parts := strings.Split(line, ",")
		cube := [dims][2]int{}
		outside := false
		for i := 0; i < dims; i++ {
			cube[i][0], cube[i][1] = parseCoordRange(parts[i])
			outside = outside || cube[i][1] < -size || size < cube[i][0]
		}

		if outside {
			continue
		}
		for i := 0; i < dims; i++ {
			cube[i][0] = intClamp(cube[i][0], -size, size)
			cube[i][1] = intClamp(cube[i][1], -size, size) + 1
			cs[i] = append(cs[i], cube[i][0], cube[i][1])
		}
		cuboids = append(cuboids, cuboid{cube, parts[0][1] == 'n'})
	}
	for d := range cs {
		sort.Ints(cs[d])
		r := 0
		for _, v := range cs[d][1:] {
			if cs[d][r] != v {
				r++
				cs[d][r] = v
			}
		}
		cs[d] = cs[d][:r+1]
	}

	for i := range cuboids {
		for d := range cuboids[i].coords {
			for c := range cuboids[i].coords[d] {
				//found := false
				for j := range cs[d] {
					if cs[d][j] == cuboids[i].coords[d][c] {
						cuboids[i].coords[d][c] = j
						//found = true
						break
					}
				}
				//println(found)
			}
		}
	}

	//printInts(cs[0])
	//printInts(cs[1])
	//printInts(cs[2])
	//
	//for _, cube := range cuboids {
	//	for _, coord := range cube.coords {
	//		for _, v := range coord {
	//			print(v, " ")
	//		}
	//	}
	//	println()
	//}

	total := 0

	for i0 := range cs[0][1:] {
		inside0 := make([]cuboid, 0, len(cuboids))
		for _, cube := range cuboids {
			if insideCoord(cube, 0, i0) {
				inside0 = append(inside0, cube)
			}
		}
		for i1 := range cs[1][1:] {
			inside1 := make([]cuboid, 0, len(inside0))
			for _, cube := range inside0 {
				if insideCoord(cube, 1, i1) {
					inside1 = append(inside1, cube)
				}
			}
			for i2 := range cs[2][1:] {
				on := false
				for _, cube := range inside1 {
					if insideCoord(cube, 0, i0) && insideCoord(cube, 1, i1) && insideCoord(cube, 2, i2) {
						on = cube.on
					}
				}
				if on {
					//println(cs[0][i0], cs[0][i0+1], cs[1][i1], cs[1][i1+1], cs[2][i2], cs[2][i2+1])
					total += (cs[0][i0+1] - cs[0][i0]) * (cs[1][i1+1] - cs[1][i1]) * (cs[2][i2+1] - cs[2][i2])
				}
			}
		}
	}

	return total
}

func insideCoord(cube cuboid, c int, i int) bool {
	return cube.coords[c][0] <= i && i < cube.coords[c][1]
}

func main() {
	println(solve22("22_0.in", 50))
	println(solve22("22_1.in", 50))
	println(solve22("22_2.in", 50))

	println(solve22("22_3.in", 1000000))
	println(solve22("22_2.in", 1000000))
}
