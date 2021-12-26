package main

func solve25(fileName string) int {
	var right [][2]int
	var down [][2]int

	var board [][]int
	lines := readLines(fileName)
	rows := len(lines)
	cols := len(lines[0])
	for r, line := range lines {
		board = append(board, make([]int, cols))
		for c, v := range line {
			if v == '>' {
				right = append(right, [2]int{r, c})
				board[r][c] = len(right)
			} else if v == 'v' {
				down = append(down, [2]int{r, c})
				board[r][c] = -len(down)
			}
		}
	}

	for step := 1; ; step++ {
		//for _, row := range board {
		//	for _, v := range row {
		//		if v > 0 {
		//			print(">")
		//		} else if v < 0 {
		//			print("v")
		//		} else {
		//			print(".")
		//		}
		//	}
		//	println()
		//}
		//println()

		moved := false

		var active []int
		for i, c := range right {
			if board[c[0]][(c[1]+1)%cols] == 0 {
				active = append(active, i)
			}
		}
		for _, i := range active {
			board[right[i][0]][right[i][1]] = 0
			right[i][1] = (right[i][1] + 1) % cols
			board[right[i][0]][right[i][1]] = i + 1
		}
		moved = moved || len(active) > 0

		active = nil
		for i, c := range down {
			if board[(c[0]+1)%rows][c[1]] == 0 {
				active = append(active, i)
			}
		}
		for _, i := range active {
			board[down[i][0]][down[i][1]] = 0
			down[i][0] = (down[i][0] + 1) % rows
			board[down[i][0]][down[i][1]] = -(i + 1)
		}
		moved = moved || len(active) > 0

		if !moved {
			return step
		}
		println("    ", step)
	}
}

func main() {
	println(solve25("25_1.in"))
	println(solve25("25_2.in"))
}
