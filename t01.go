package main

func solve_01(file string, step int) {
	data := readLines(file)

	result := 0
	for i := step; i < len(data); i++ {
		if data[i-step] < data[i] {
			result++
		}
	}
	println(result)
}

func main() {
	solve_01("01_1.in", 1)
	solve_01("01_2.in", 1)
	solve_01("01_1.in", 3)
	solve_01("01_2.in", 3)
}
