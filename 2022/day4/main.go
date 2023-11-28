package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)

	overlap := 0
	overlapAny := 0

	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")

		firstRange := strings.Split(pair[0], "-")
		secondRange := strings.Split(pair[1], "-")
		println(firstRange[0], firstRange[1])

		x, _ := strconv.Atoi(firstRange[0]) // 4
		y, _ := strconv.Atoi(firstRange[1]) // 8

		a, _ := strconv.Atoi(secondRange[0]) // 2
		b, _ := strconv.Atoi(secondRange[1]) // 6

		println(x, y)

		if (a <= x && b >= y) || (x <= a && y >= b) {
			overlap++
		}

		if (x <= a && a <= y) || (x <= b && b <= y) {
			overlapAny++
		} else if (a <= x && x <= b) || (a <= y && y <= b) {
			overlapAny++
		}

	}

	println(overlap)
	println(overlapAny)

}
