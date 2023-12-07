package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf(
		"Example 1: %d\nPuzzle 1: %d\nExample 2: %d\nPuzzle 2: %d\n",
		runExample1(),
		runPuzzle1(),
		runExample2(),
		runPuzzle2(),
	)
}

func runExample1() int {
	return puzzle1Helper("example.txt")
}

func runPuzzle1() int {
	return puzzle1Helper("puzzle.txt")
}

var numRE = regexp.MustCompile("([0-9]+)")

func puzzle1Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read file: %s\n", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	timeNums := numRE.FindAllString(lines[0], -1)
	distanceNums := numRE.FindAllString(lines[1], -1)
	possibilities := 1
	for i := range timeNums {
		raceTime, err := strconv.Atoi(timeNums[i])
		if err != nil {
			log.Fatalf("failed to convert time %s to a number: %s\n", timeNums[i], err.Error())
		}
		raceRecord, err := strconv.Atoi(distanceNums[i])
		if err != nil {
			log.Fatalf("failed to convert distance %s to a number: %s\n", distanceNums[i], err.Error())
		}
		currentPossiblities := 0
		for i := 1; i < raceTime; i++ {
			distance := (raceTime - i) * i
			if distance > raceRecord {
				currentPossiblities += 1
			}
		}
		possibilities *= currentPossiblities
	}
	return possibilities
}

func runExample2() int {
	return puzzle2Helper("example.txt")
}

func runPuzzle2() int {
	return puzzle2Helper("puzzle.txt")
}

func puzzle2Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read file: %s\n", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	timeNums := numRE.FindAllString(lines[0], -1)
	distanceNums := numRE.FindAllString(lines[1], -1)
	raceTime, err := strconv.Atoi(strings.Join(timeNums, ""))
	if err != nil {
		log.Fatalf("failed to convert time %s to a number: %s\n", strings.Join(timeNums, ""), err.Error())
	}
	raceRecord, err := strconv.Atoi(strings.Join(distanceNums, ""))
	if err != nil {
		log.Fatalf("failed to convert distance %s to a number: %s\n", strings.Join(distanceNums, ""), err.Error())
	}
	possibilities := 0
	for i := 1; i < raceTime; i++ {
		distance := (raceTime - i) * i
		if distance > raceRecord {
			possibilities += 1
		}
	}
	return possibilities
}
