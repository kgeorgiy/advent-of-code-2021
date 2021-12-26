package main

import (
	"strings"
)

type row struct {
	direction string
	count     int
}

func readFile(fileName string) (rows []row) {
	rows = []row{}
	for _, line := range readLines(fileName) {
		parts := strings.Split(line, " ")
		rows = append(rows, row{parts[0], parseInt(parts[1])})
	}

	return rows
}

func solve02one(fname string) {
	x := 0
	y := 0
	for _, row := range readFile(fname) {
		if row.direction == "up" {
			y -= row.count
		}
		if row.direction == "down" {
			y += row.count
		}
		if row.direction == "forward" {
			x += row.count
		}
	}
	println(x, y)
	println(x * y)
}

func solve02two(fname string) {
	x := 0
	y := 0
	aim := 0
	for _, row := range readFile(fname) {
		if row.direction == "up" {
			aim -= row.count
		}
		if row.direction == "down" {
			aim += row.count
		}
		if row.direction == "forward" {
			x += row.count
			y += aim * row.count
		}
		//        println(x, y, aim)
	}
	println(x, y)
	println(x * y)
}

func main() {
	solve02one("02_1.in")
	solve02one("02_2.in")
	solve02two("02_1.in")
	solve02two("02_2.in")
}
