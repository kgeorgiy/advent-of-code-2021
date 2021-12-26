package main

func solve06(fileName string, days int) int {
	ages := parseInts(readLines(fileName)[0], ",")
	total := len(ages)

	births := make([]int, days+10)
	for _, age := range ages {
		births[age]++
	}

	for d := 0; d < days; d++ {
		total += births[d]
		births[d+7] += births[d]
		births[d+9] += births[d]
		//println(d, total)
	}

	//print()
	return total
}

func main() {
	println(solve06("06_1.in", 80))
	println(solve06("06_2.in", 80))
	println(solve06("06_1.in", 256))
	println(solve06("06_2.in", 256))
}
