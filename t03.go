package main

func most(bit int, lines []string, selector bool) int {
	ones := 0
	for _, line := range lines {
		if line[bit] == '1' {
			ones++
		}
	}
	if (ones*2 > len(lines)) == selector {
		return 1
	} else {
		return 0
	}
}

func solve03one(fname string) {
	lines := readLines(fname)
	gamma := 0
	epsilon := 0
	for i := range lines[0] {
		gamma = gamma*2 + most(i, lines, true)
		epsilon = epsilon*2 + most(i, lines, false)
	}
	println(gamma, epsilon)
	println(gamma * epsilon)
}

func solve_03_2_2(lines []string, selector bool) int {
	current := lines
	for i := range lines[0] {
		next0 := []string{}
		next1 := []string{}
		ones := 0
		for _, line := range current {
			if line[i] == '1' {
				next1 = append(next1, line)
				ones++
			} else {
				next0 = append(next0, line)
			}
		}
		if (ones*2 >= len(current)) == selector {
			current = next1
		} else {
			current = next0
		}
		if len(current) == 1 {
			break
		}
	}

	value := 0
	for _, v := range current[0] {
		value = value*2 + int(v-'0')
	}
	println(selector, len(current), current[0], value)
	return value
}

func solve03two(fileName string) {
	lines := readLines(fileName)
	gamma := solve_03_2_2(lines, true)
	epsilon := solve_03_2_2(lines, false)
	println(gamma, epsilon)
	println(gamma * epsilon)
}

func main() {
	solve03one("03_1.in")
	solve03one("03_2.in")
	solve03two("03_1.in")
	solve03two("03_2.in")
}
