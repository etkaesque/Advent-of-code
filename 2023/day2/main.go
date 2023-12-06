package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkIfGamePossible(cubesSetsArr []string) bool {

	for _, set := range cubesSetsArr {

		cubesArr := strings.Split(set, ",")

		for _, cube := range cubesArr {

			fmt.Println("cubes", cube)
			cubeTrimed := strings.Trim(cube, " ")
			cube := strings.Split(cubeTrimed, " ")
			cubeNum, _ := strconv.Atoi(cube[0])
			cubeColor := cube[1]

			if strings.Contains(cubeColor, "blue") && cubeNum > 14 {
				return false
			}

			if strings.Contains(cubeColor, "red") && cubeNum > 12 {
				return false
			}

			if strings.Contains(cubeColor, "green") && cubeNum > 13 {
				return false
			}

		}

	}

	return true

}

func calculateSumOfPower(cubesSetsArr []string) int {

	redMin := 1
	greenMin := 1
	blueMin := 1

	for _, set := range cubesSetsArr {

		cubesArr := strings.Split(set, ",")

		for _, cube := range cubesArr {

			fmt.Println("cubes", cube)
			cubeTrimed := strings.Trim(cube, " ")
			cube := strings.Split(cubeTrimed, " ")
			cubeNum, _ := strconv.Atoi(cube[0])
			cubeColor := cube[1]

			if strings.Contains(cubeColor, "blue") && cubeNum >= blueMin {
				blueMin = cubeNum
			}

			if strings.Contains(cubeColor, "red") && cubeNum >= redMin {
				redMin = cubeNum
			}

			if strings.Contains(cubeColor, "green") && cubeNum >= greenMin {
				greenMin = cubeNum
			}

		}

	}

	powerSet := redMin * greenMin * blueMin

	return powerSet

}

func main() {
	file, err := os.Open("./input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)

	sum := 0
	powerSum := 0

	for scanner.Scan() {

		line := scanner.Text()

		game := strings.Split(line, ":")[0]
		game = strings.Split(game, " ")[1]
		gameId, _ := strconv.Atoi(game)

		cubesSets := strings.Split(line, ":")[1]
		cubesSetsArr := strings.Split(cubesSets, ";")

		isOk := checkIfGamePossible(cubesSetsArr)
		powerSet := calculateSumOfPower(cubesSetsArr)

		powerSum += powerSet

		if isOk {
			sum += gameId
		}

	}

	fmt.Println("Sum of game id's: ", sum)
	fmt.Println("Power sum of sets", powerSum)
}
