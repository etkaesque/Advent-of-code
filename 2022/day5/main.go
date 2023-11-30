package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	moveCount int
	moveTo    int
	moveFrom  int
}

func main() {

	partOne()

}

func reverseSlice(arg []string) []string {
	slice := arg
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func partOne() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal()
	}

	content := strings.Split(string(file), "\n\n")
	stack := strings.Split(content[0], "\n")
	moves := strings.Split(content[1], "\n")
	stackCol := strings.Split(strings.ReplaceAll(stack[len(stack)-1], " ", ""), "")
	stack = stack[:len(stack)-1]
	incrementIndex := 0
	stackMap := map[int][]string{}
	for _, col := range stackCol {
		stackSlice := []string{}
		colNum, _ := strconv.Atoi(col)

		for _, line := range stack {
			index := 1 + incrementIndex
			if incrementIndex > len(line) {
				index = len(line) - 2
			}
			char := string(line[index])
			if char != " " {
				stackSlice = append(stackSlice, char)
				stackMap[colNum] = stackSlice
			}

		}
		incrementIndex += 4
	}

	movesSlice := []Move{}
	for _, line := range moves {
		var moveCount, moveFrom, moveTo int

		fmt.Sscanf(line, "move %d from %d to %d", &moveCount, &moveFrom, &moveTo)

		move := Move{
			moveCount: moveCount,
			moveFrom:  moveFrom,
			moveTo:    moveTo,
		}
		movesSlice = append(movesSlice, move)

	}
	fmt.Println(stackMap)
	for _, move := range movesSlice {
		stackFrom := stackMap[move.moveFrom]
		stackTo := stackMap[move.moveTo]
		firstIndex := move.moveCount

		newStackFrom := make([]string, len(stackFrom[firstIndex:]))
		copy(newStackFrom, stackFrom[firstIndex:])

		appendTheseLetters := make([]string, len(stackFrom[0:firstIndex]))
		copy(appendTheseLetters, stackFrom[0:firstIndex])

		appendTheseLetters = reverseSlice(appendTheseLetters)

		newStackTo := make([]string, len(appendTheseLetters)+len(stackTo))
		newStackTo = append(appendTheseLetters, stackTo...)

		stackMap[move.moveFrom] = newStackFrom
		stackMap[move.moveTo] = newStackTo

		fmt.Println(stackFrom, stackTo, newStackFrom, newStackTo, appendTheseLetters)
	}

	fmt.Println(stackMap)

	for i := 0; i < len(stackMap); i++ {
		fmt.Println(stackMap[i+1][0]) // answer
	}

}
