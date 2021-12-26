package main

func solve19(fileName string) int {
	var scanners [][][3]int
	{
		var scanner [][3]int
		for _, line := range readLines(fileName) {
			if line[:3] == "---" {
				scanners = append(scanners, scanner)
				scanner = nil
			} else {
				ints := parseInts(line, ",")
				scanner = append(scanner, [3]int{ints[0], ints[1], ints[2]})
			}
		}
		scanners = append(scanners, scanner)
	}

	beacons := scanners[0]
	rest := len(scanners) - 1
	unified := make([][]int, len(scanners))
	unified[0] = []int{0, 0, 0}
	for {
		for i, scanner := range scanners {
			if unified[i] == nil {
				unified[i], beacons = unifyScanner(beacons, scanner)
				if unified[i] != nil {
					rest--
					if rest == 0 {
						println(len(beacons))
						best := 0
						for _, b1 := range unified {
							for _, b2 := range unified {
								distance := intAbs(b1[0]-b2[0]) + intAbs(b1[1]-b2[1]) + intAbs(b1[2]-b2[2])
								best = intMax(best, distance)
							}
						}
						return best
					}
				}
			}
		}
	}
}

func unifyScanner(global [][3]int, local [][3]int) ([]int, [][3]int) {
	points := map[[3]int]bool{}
	for _, b := range global {
		points[b] = true
	}

	for rx := 0; rx < 4; rx++ {
		for ry := 0; ry < 4; ry++ {
			for rz := 0; rz < 4; rz++ {
				for _, g := range global {
					for _, l := range local {
						count := 0
						point := [3]int{}
						for _, p := range local {
							for i := range point {
								point[i] = p[i] - l[i] + g[i]
								if points[point] {
									count++
								}
							}
						}
						if count >= 12 {
							for _, p := range local {
								for i := range point {
									point[i] = p[i] - l[i] + g[i]
								}
								if !points[point] {
									global = append(global, [3]int{point[0], point[1], point[2]})
								}
							}
							println(count, len(global))
							for i := range point {
								point[i] = g[i] - l[i]
							}
							//printInts(point[:])
							return point[:], global
						}
					}
				}
				rotate(local, 0, 1)
			}
			rotate(local, 0, 2)
		}
		rotate(local, 1, 2)
	}
	return nil, global
}

func rotate(local [][3]int, axis1 int, axis2 int) {
	for i := range local {
		c1 := local[i][axis1]
		c2 := local[i][axis2]
		local[i][axis1] = c2
		local[i][axis2] = -c1
	}
}

func main() {
	println(solve19("19_1.in"))
	println(solve19("19_2.in"))
}
