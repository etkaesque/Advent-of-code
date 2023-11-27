package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func partOne() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

		compMap := make(map[string]int)

		for _, char := range compartment1 {

			num := 0
			if unicode.IsUpper(char) {
				num = int(char) - 38

			} else {
				num = int(char) - 96
			}
			compMap[string(char)] = num
		}

		for _, char := range compartment2 {
			value, ok := compMap[string(char)]
			if ok {
				sum += value
				break
			}

		}

	}

	fmt.Println(sum)

}

func partTwo() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	group := []string{}
	lineCount := 1
	for scanner.Scan() {
		line := scanner.Text()

		group = append(group, line)
		if lineCount%3 == 0 {

			elfOneMap := make(map[string]int)

			for _, char := range group[0] {

				num := 0
				if unicode.IsUpper(char) {
					num = int(char) - 38

				} else {
					num = int(char) - 96
				}
				elfOneMap[string(char)] = num
			}

			elfTwoMap := make(map[string]int)

			for _, char := range group[1] {

				num := 0
				if unicode.IsUpper(char) {
					num = int(char) - 38

				} else {
					num = int(char) - 96
				}
				elfTwoMap[string(char)] = num
			}

			for _, char := range group[2] {
				_, isElfOneOk := elfOneMap[string(char)]
				value, isElfTwoOk := elfTwoMap[string(char)]

				if isElfOneOk && isElfTwoOk {
					sum += value
					break
				}

			}

			group = []string{}
		}

		lineCount++

	}

	fmt.Println(sum)

}

func main() {
	partOne()
	partTwo()

}
