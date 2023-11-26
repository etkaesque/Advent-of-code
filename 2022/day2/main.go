package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func partOne() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		x := strings.Contains(line, "X")
		y := strings.Contains(line, "Y")
		z := strings.Contains(line, "Z")

		a := strings.Contains(line, "A")
		b := strings.Contains(line, "B")
		c := strings.Contains(line, "C")

		if x {
			score += 1
		} else if y {
			score += 2
		} else if z {
			score += 3
		}

		if x && a || y && b || z && c {
			score += 3
		} else if x && c || y && a || z && b {
			score += 6
		}

	}

	fmt.Println(score)

}

func partTwo() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		x := strings.Contains(line, "X") //need to lose
		y := strings.Contains(line, "Y") //need to draw
		z := strings.Contains(line, "Z") // need to win

		a := strings.Contains(line, "A")
		b := strings.Contains(line, "B")
		c := strings.Contains(line, "C")

		if a && x || b && x || c && x {
			score += 0

			if a {
				score += 3
			} else if b {
				score += 1
			} else if c {
				score += 2
			}

		} else if a && y || b && y || c && y {
			score += 3

			if a {
				score += 1
			} else if b {
				score += 2
			} else if c {
				score += 3
			}

		} else if a && z || b && z || c && z {
			score += 6

			if a {
				score += 2
			} else if b {
				score += 3
			} else if c {
				score += 1
			}
		}

	}

	fmt.Println(score)

}

func main() {

	partOne()
	partTwo()

}
