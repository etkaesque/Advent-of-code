package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Symbol struct {
	line  int
	index int
}

func findNumber(mid string, right string, left string, partTwo bool) (int, bool) {

	n := 0

	if mid != "x" && left != "x" && right != "x" || mid != "x" && left != "x" || mid != "x" && right != "x" {
		n, _ = strconv.Atoi(mid)

	} else if mid == "x" && left != "x" && right != "x" {
		nr, _ := strconv.Atoi(left)
		nl, _ := strconv.Atoi(right)

		if partTwo {
			n = nr * nl
			return n, true
		} else {
			n = nr + nl
			return n, false
		}

	} else if right != "x" && mid == "x" {
		n, _ = strconv.Atoi(right)

	} else if left != "x" && mid == "x" {
		n, _ = strconv.Atoi(left)

	} else if mid != "x" {
		n, _ = strconv.Atoi(mid)

	}

	return n, false

}

func partOne() {

	symbols := []Symbol{}
	lines := make(map[int]map[int]string)
	file, err := os.Open("./input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)
	l := 1

	for scanner.Scan() {

		line := scanner.Text()
		lines[l] = make(map[int]string)
		number := ""

		for position, char := range line {
			var isCharSymbol bool
			if (char >= 33 && char <= 45) || (char == 47) || (char >= 60 && char <= 64) {
				isCharSymbol = true
			}

			if unicode.IsDigit(char) {
				number += string(char)
			}

			if unicode.IsDigit(char) && position == len(line)-1 || !unicode.IsDigit(char) {

				for j := range number {

					lines[l][position-j-1] = number
				}
				lines[l][position] = "x"
				number = ""

			}

			if isCharSymbol {
				symbols = append(symbols, Symbol{
					line:  l,
					index: position,
				})
			}

		}

		l++
	}

	sum := 0

	for _, sym := range symbols {
		h, lo, m := 0, 0, 0

		highLineMid := lines[sym.line-1][sym.index]
		highLineLeft := lines[sym.line-1][sym.index-1]
		highLineRight := lines[sym.line-1][sym.index+1]

		lowLineMid := lines[sym.line+1][sym.index]
		lowLineLeft := lines[sym.line+1][sym.index-1]
		lowLineRight := lines[sym.line+1][sym.index+1]

		midLineLeft := lines[sym.line][sym.index-1]
		midLineRight := lines[sym.line][sym.index+1]

		ml, mr := 0, 0

		if midLineLeft != "x" && midLineRight != "x" {
			ml, _ = strconv.Atoi(midLineLeft)
			mr, _ = strconv.Atoi(midLineRight)
			m = ml + mr
		} else if midLineRight != "x" {
			m, _ = strconv.Atoi(midLineRight)
		} else if midLineLeft != "x" {
			m, _ = strconv.Atoi(midLineLeft)
		}

		h, _ = findNumber(highLineMid, highLineRight, highLineLeft, false)
		lo, _ = findNumber(lowLineMid, lowLineRight, lowLineLeft, false)

		sum += m + lo + h

	}

	fmt.Println(sum)

}

func partTwo() {

	symbols := []Symbol{}
	lines := make(map[int]map[int]string)
	file, err := os.Open("./input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)
	l := 1

	for scanner.Scan() {

		line := scanner.Text()
		lines[l] = make(map[int]string)
		number := ""

		for position, char := range line {
			var isCharSymbol bool
			if char == '*' {
				isCharSymbol = true
			}

			if unicode.IsDigit(char) {
				number += string(char)
			}

			if unicode.IsDigit(char) && position == len(line)-1 || !unicode.IsDigit(char) {

				for j := range number {

					lines[l][position-j-1] = number
				}
				lines[l][position] = "x"
				number = ""

			}

			if isCharSymbol {
				symbols = append(symbols, Symbol{
					line:  l,
					index: position,
				})
			}

		}

		l++
	}

	gearRatio := 0

	for _, sym := range symbols {
		h, lo, m := 0, 0, 0

		highLineMid := lines[sym.line-1][sym.index]
		highLineLeft := lines[sym.line-1][sym.index-1]
		highLineRight := lines[sym.line-1][sym.index+1]

		lowLineMid := lines[sym.line+1][sym.index]
		lowLineLeft := lines[sym.line+1][sym.index-1]
		lowLineRight := lines[sym.line+1][sym.index+1]

		midLineLeft := lines[sym.line][sym.index-1]
		midLineRight := lines[sym.line][sym.index+1]

		ml, mr := 0, 0

		if midLineLeft != "x" && midLineRight != "x" {
			ml, _ = strconv.Atoi(midLineLeft)
			mr, _ = strconv.Atoi(midLineRight)
			m = ml * mr
			gearRatio += m
			continue
		} else if midLineRight != "x" {
			m, _ = strconv.Atoi(midLineRight)
		} else if midLineLeft != "x" {
			m, _ = strconv.Atoi(midLineLeft)
		}

		h, ajdH := findNumber(highLineMid, highLineRight, highLineLeft, true)
		lo, ajdLo := findNumber(lowLineMid, lowLineRight, lowLineLeft, true)

		if ajdH {
			gearRatio += h
			continue
		}

		if ajdLo {
			gearRatio += lo
			continue
		}

		if h == 0 && lo == 0 && m == 0 {
			continue
		}

		if h > 0 && lo == 0 && m == 0 {
			continue
		}

		if lo > 0 && h == 0 && m == 0 {
			continue
		}

		if m > 0 && h == 0 && lo == 0 {
			continue
		}

		if lo == 0 {
			lo = 1
		}

		if h == 0 {
			h = 1
		}

		if m == 0 {
			m = 1
		}
		gearRatio += h * lo * m
	}
	fmt.Println(gearRatio)
}

func main() {
	partOne()
	partTwo()
}
