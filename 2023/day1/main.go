package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	partTwo()

}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	for scanner.Scan() {

		line := scanner.Text()

		lineSlice := strings.Split(line, "")
		fmt.Println(lineSlice)
		number := ""

		for i, item := range line {
			IsDigit := unicode.IsDigit(rune(item))

			if IsDigit {
				number += string(item)
			}

			if i == len(lineSlice)-1 {
				var firstDigit string
				var secondDigit string

				digits := strings.Split(number, "")
				len := len(digits)

				firstDigit = digits[0]
				secondDigit = digits[len-1]

				combination := firstDigit + secondDigit

				numberInt, _ := strconv.Atoi(combination)
				fmt.Println(numberInt)
				number = ""
				numbers = append(numbers, numberInt)

			}

		}

	}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Println(sum)

}

type letterNum struct {
	word  string
	value int
}

type numberPlace struct {
	index  int
	number string
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {

		line := scanner.Text()

		letterNums := map[string]letterNum{
			"one": {
				word:  "one",
				value: 1,
			},
			"two": {
				word:  "two",
				value: 2,
			},
			"three": {
				word:  "three",
				value: 3,
			},
			"four": {
				word:  "four",
				value: 4,
			},
			"five": {
				word:  "five",
				value: 5,
			},
			"six": {
				word:  "six",
				value: 6,
			},
			"seven": {
				word:  "seven",
				value: 7,
			},
			"eight": {
				word:  "eight",
				value: 8,
			},
			"nine": {
				word:  "nine",
				value: 9,
			},
		}

		lowestIndex := len(line) - 1
		lowestWord := ""

		for _, mapItem := range letterNums {
			index := strings.Index(line, mapItem.word)

			if index == -1 {
				continue
			} else if index < lowestIndex {
				lowestIndex = index
				lowestWord = mapItem.word
			}
		}

		if lowestIndex == len(line)-1 {
			lowestIndex = -1
		}

		highestIndex := -1
		highestWord := ""

		for _, mapItem := range letterNums {
			index := strings.LastIndex(line, mapItem.word)

			if index > highestIndex {
				highestIndex = index
				highestWord = mapItem.word
			}

		}

		numbers := []numberPlace{}

		for i, item := range line {
			IsDigit := unicode.IsDigit(rune(item))

			if IsDigit {
				numbers = append(numbers, numberPlace{
					index:  i,
					number: string(item),
				})
			}
		}

		digitOne := numbers[0].number
		digitTwo := numbers[len(numbers)-1].number

		if lowestIndex != -1 && lowestIndex < numbers[0].index {
			digitOne = strconv.Itoa(letterNums[lowestWord].value)

		}

		if highestIndex != -1 && highestIndex > numbers[len(numbers)-1].index {
			digitTwo = strconv.Itoa(letterNums[highestWord].value)
		}

		fmt.Println("lowestIndex", lowestIndex, "highestIndex", highestIndex)

		numbers = []numberPlace{}

		combination := digitOne + digitTwo
		numberInt, _ := strconv.Atoi(combination)

		fmt.Println(combination, "digitOne", digitOne, "digitTwo", digitTwo, line)

		sum += numberInt

	}

	println(sum)

}
