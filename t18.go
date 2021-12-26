package main

import "strconv"

type snail struct {
	left, right interface{}
}

func snailToString(value interface{}) string {
	if v, ok := value.(snail); ok {
		return "[" + snailToString(v.left) + "," + snailToString(v.right) + "]"
	} else {
		return strconv.Itoa(value.(int))
	}
}

func snailMagnitude(value interface{}) int {
	if v, ok := value.(snail); ok {
		return 3*snailMagnitude(v.left) + 2*snailMagnitude(v.right)
	} else {
		return value.(int)
	}
}

func snailParse(input string) (interface{}, string) {
	if input[0] != '[' {
		return int(input[0] - '0'), input[1:]
	}
	left, input := snailParse(input[1:])
	right, input := snailParse(input[1:])
	return snail{left, right}, input[1:]
}

func solve18(fileName string, max bool) int {
	var numbers []interface{}
	for _, line := range readLines(fileName) {
		if line == "---" {
			solve18Case(numbers, max)
			numbers = nil
		} else {
			n, _ := snailParse(line)
			numbers = append(numbers, n)
		}
	}
	return solve18Case(numbers, max)
}

func solve18Case(numbers []interface{}, max bool) int {
	if max {
		best := 0
		for i := range numbers {
			for j := range numbers {
				if i != j {
					best = intMax(best, snailMagnitude(snailReduce(snail{numbers[i], numbers[j]})))
				}
			}
		}
		println(best)
		return best
	} else {
		sum := numbers[0]
		for _, n := range numbers[1:] {
			sum = snailReduce(snail{sum, n})
			//println("   ", snailToString(n), snailToString(sum))
		}
		println(snailMagnitude(sum), snailToString(sum))
		return snailMagnitude(sum)
	}
}

func snailReduce(value interface{}) interface{} {
	for {
		var e bool
		for {
			//old := snailToString(value)
			e, _, value, _ = snailExplode(value, 0)
			//println("    ", old, "->", snailToString(value))
			if !e {
				break
			}
		}
		//old := snailToString(value)
		e, value = snailSplit(value)
		//println("    ", old, "-->", snailToString(value))
		if !e {
			return value
		}
	}
}

func snailSplit(value interface{}) (bool, interface{}) {
	if v, ok := value.(snail); ok {
		s, m := snailSplit(v.left)
		if s {
			return true, snail{m, v.right}
		}
		s, m = snailSplit(v.right)
		if s {
			return true, snail{v.left, m}
		}
		return false, v
	} else {
		v := value.(int)
		if v > 9 {
			//println(true, snailToString(snail{v / 2, v - v/2}))
			return true, snail{v / 2, v - v/2}
		} else {
			return false, v
		}
	}
}

func snailExplode(value interface{}, depth int) (bool, int, interface{}, int) {
	//println(depth, snailToString(value))
	if v, ok := value.(snail); ok {
		if depth == 4 {
			//println(true, v.left.(int), 0, v.right.(int))
			return true, v.left.(int), 0, v.right.(int)
		} else {
			e, l, m, r := snailExplode(v.left, depth+1)
			if e {
				//println(snailToString(m), snailToString(v.right))
				//println(snailToString(snailAddLeft(v.right, r)))
				//println(true, l, snailToString(snail{m, snailAddLeft(v.right, r)}), 0)
				return true, l, snail{m, snailAddLeft(v.right, r)}, 0
			}
			e, l, m, r = snailExplode(v.right, depth+1)
			if e {
				//println(snailToString(v.left), snailToString(m))
				//println(snailToString(snailAddRight(v.left, l)))
				//println(true, 0, snailToString(snail{snailAddRight(v.left, l), m}), r)
				return true, 0, snail{snailAddRight(v.left, l), m}, r
			}
			//println("->", false, 0, snailToString(v), 0)
			return false, 0, v, 0
		}
	} else if v, ok := value.(int); ok {
		//println("->", false, 0, v, 0)
		return false, 0, v, 0
	} else {
		panic(value)
	}
}

func snailAddLeft(value interface{}, delta int) interface{} {
	//if delta == 0 {
	//	return value
	//}
	//println("snailAddLeft <", snailToString(value), delta)
	if v, ok := value.(snail); ok {
		//println("snailAddLeft >", snailToString(snail{snailAddLeft(v.left, delta), v.right}))
		return snail{snailAddLeft(v.left, delta), v.right}
	} else {
		//println("snailAddLeft >", snailToString(value))
		return value.(int) + delta
	}
}

func snailAddRight(value interface{}, delta int) interface{} {
	//if delta == 0 {
	//	return value
	//}
	if v, ok := value.(snail); ok {
		return snail{v.left, snailAddRight(v.right, delta)}
	} else {
		return value.(int) + delta
	}
}

func main() {
	println(solve18("18_1.in", false))
	println(solve18("18_2.in", false))
	println(solve18("18_1.in", true))
	println(solve18("18_2.in", true))
}
