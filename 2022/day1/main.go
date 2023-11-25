package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	elf      int
	calories []int
	total    int
}

func parseFile(file string) []string {

	fileOutput, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	content := fmt.Sprintf(string(fileOutput))

	lines := strings.Split(content, "\n")

	return lines
}

func sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

func formatElvesToSlice(lines []string) []Elf {

	elves := []Elf{}
	calories := []int{}
	elfCount := 1
	for _, item := range lines {
		if item == "" {

			elf := Elf{
				elf:      elfCount,
				calories: calories,
				total:    sum(calories),
			}

			elves = append(elves, elf)
			calories = []int{}
			elfCount++

		} else {
			cal, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			calories = append(calories, cal)
		}
	}

	return elves
}

func findElfWithMostCalories(elves []Elf) {

	elfWithMostCalories := elves[0]
	for i := 1; i < len(elves); i++ {

		if elves[i].total > elfWithMostCalories.total {
			elfWithMostCalories = elves[i]
		}
	}

	fmt.Fprintf(os.Stdout, "%vth elf has %v calories\n", elfWithMostCalories.elf, elfWithMostCalories.total)

}

func findTopElvesWithMostCalories(slice []Elf, top int) {

	topSum := 0
	elves := slice

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].total > elves[j].total
	})

	topElves := elves[:top]

	for i := 0; i < len(topElves); i++ {
		topSum += topElves[i].total
	}

	fmt.Fprintf(os.Stdout, "Top %v elves carry %v calories\n", top, topSum)

}

func main() {
	content := parseFile("input.txt")
	elves := formatElvesToSlice(content)

	findElfWithMostCalories(elves)

	findTopElvesWithMostCalories(elves, 3)
}
