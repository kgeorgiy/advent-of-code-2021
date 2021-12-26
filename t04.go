package main

type bingoBoard struct {
	field [][]int
	rows  []int
	cols  []int
	won   bool
}

func solve04(fileName string, first bool) int {
	numbers, fields := readBingo(fileName)

	rest := len(fields)
	for _, n := range numbers {
		for i := range fields {
			field := &fields[i]
			if !field.won && play(field, n) {
				rest--
				if first || rest == 0 {
					return bingoSum(*field, n)
				}
			}
		}
	}
	return 0
}

func bingoSum(field bingoBoard, n int) int {
	sum := 0
	for r := range field.field {
		for c := range field.field[r] {
			if field.field[r][c] >= 0 {
				sum += field.field[r][c]
			}
		}
	}

	println(n, sum)
	return n * sum
}

func readBingo(fileName string) ([]int, []bingoBoard) {
	size := 5
	lines := readLines(fileName)
	p := 0
	numbers := parseInts(lines[p], ",")
	p++

	var fields []bingoBoard = nil
	for p < len(lines) {
		field := bingoBoard{make([][]int, size, size), make([]int, size), make([]int, size), false}
		for i := 0; i < size; i++ {
			field.field[i] = parseInts(lines[p], " ")
			p++
		}
		fields = append(fields, field)
	}
	return numbers, fields
}

func play(field *bingoBoard, n int) bool {
	for r := range field.field {
		for c := range field.field[r] {
			if field.field[r][c] == n {
				field.field[r][c] = -1
				field.rows[r]++
				field.cols[c]++
				if field.rows[r] == len(field.field) || field.cols[c] == len(field.field[r]) {
					field.won = true
					return true
				}
			}
		}
	}
	return false
}

func main() {
	println(solve04("04_1.in", true))
	println(solve04("04_2.in", true))
	println(solve04("04_1.in", false))
	println(solve04("04_2.in", false))
}
