package main

func solve20(fileName string, steps int) int {
	lines := readLines(fileName)
	transform := make([]int, len(lines[0]))
	for i, c := range lines[0] {
		if c == '#' {
			transform[i] = 1
		}
	}

	border := steps + 5
	var image = make([][]int, len(lines)-1+border*2)
	for i := range image {
		image[i] = make([]int, len(lines[1])+border*2)
	}

	for i, line := range lines[1:] {
		for j, c := range line {
			if c == '#' {
				image[i+border+1][j+border+1] = 1
			}
		}
	}

	//printInts2(image)
	for s := 0; s < steps; s++ {
		next := make([][]int, len(image))
		for i := range image {
			next[i] = make([]int, len(image[i]))
		}
		var inf int
		if image[0][0] == 0 {
			inf = transform[0]
		} else {
			inf = transform[511]
		}

		for c := 0; c < len(image[0]); c++ {
			next[0][c] = inf
			next[len(next)-1][c] = inf
		}
		for r := 1; r < len(image)-1; r++ {
			next[r][0] = inf
			next[r][len(next[0])-1] = inf

			for c := 1; c < len(image[r])-1; c++ {
				index := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						index = index*2 + image[r+dr][c+dc]
					}
				}
				next[r][c] = transform[index]
			}
		}
		image = next
		//for _, row := range next {
		//	for _, v := range row {
		//		print(string(".#"[v]))
		//	}
		//	println()
		//}
		//println()
	}

	sum := 0
	for _, row := range image {
		sum += sumInts(row)
	}
	return sum
}

func main() {
	println(solve20("20_1.in", 2))
	println(solve20("20_2.in", 2))
	println(solve20("20_1.in", 50))
	println(solve20("20_2.in", 50))
}
