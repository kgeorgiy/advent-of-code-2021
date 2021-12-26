package main

func solve11(fileName string, steps int) int {
	energy := parseInts2(readLines(fileName), "")

	total := 0
	for i := 0; steps == -1 || i < steps; i++ {
		step := 0
		for r := range energy {
			for c := range energy[r] {
				step += incEnergy(energy, r, c)
			}
		}

		for r := range energy {
			for c := range energy[r] {
				if energy[r][c] == -1 {
					energy[r][c] = 0
				}
			}
		}
		if steps == -1 && step == len(energy)*len(energy[0]) {
			return i + 1
		}
		total += step
	}

	return total
}

func incEnergy(energy [][]int, r int, c int) int {
	if !(0 <= r && r < len(energy) && 0 <= c && c < len(energy[r])) {
		return 0
	}

	if energy[r][c] != -1 {
		energy[r][c]++
		if energy[r][c] == 10 {
			//println(r, c)
			energy[r][c] = -1

			total := 1
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					total += incEnergy(energy, r+dr, c+dc)
				}
			}
			return total
		}
	}
	return 0
}

func main() {
	println(solve11("11_1.in", 100))
	println(solve11("11_2.in", 100))
	println(solve11("11_1.in", -1))
	println(solve11("11_2.in", -1))
}
