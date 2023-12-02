package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type draw struct {
	red   int
	blue  int
	green int
}

const (
	PUZZLE_1_RED_LIMIT   = 12
	PUZZLE_1_BLUE_LIMIT  = 14
	PUZZLE_1_GREEN_LIMIT = 13
)

var lineRE = regexp.MustCompile("^Game ([0-9]+): (.+)$")
var drawRE = regexp.MustCompile("^([0-9]+) ((red)|(blue)|(green))$")

func main() {
	fmt.Printf("Example 1: %d\nPuzzle 1: %d\nExample 2: %d\nPuzzle 2: %d\n", solveExample1(), solvePuzzle1(), solveExample2(), solvePuzzle2())
}

func solveExample1() int {
	return puzzle1Helper("example1.txt")
}

func solvePuzzle1() int {
	return puzzle1Helper("puzzle.txt")
}

func puzzle1Helper(fileName string) int {
	input, _ := os.ReadFile(path.Join("inputs", fileName))
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		isValid := true
		lineMatches := lineRE.FindStringSubmatch(line)
		gameID, _ := strconv.Atoi(lineMatches[1])
		results := lineMatches[2]
		rounds := strings.Split(results, "; ")
		for _, round := range rounds {
			roundDraw := draw{}
			hands := strings.Split(round, ", ")
			for _, hand := range hands {
				handMatches := drawRE.FindStringSubmatch(hand)
				color := handMatches[2]
				count, _ := strconv.Atoi(handMatches[1])
				switch color {
				case "red":
					roundDraw.red += count
				case "blue":
					roundDraw.blue += count
				case "green":
					roundDraw.green += count
				}
			}
			if roundDraw.red > PUZZLE_1_RED_LIMIT || roundDraw.blue > PUZZLE_1_BLUE_LIMIT || roundDraw.green > PUZZLE_1_GREEN_LIMIT {
				isValid = false
				break
			}
		}
		if isValid {
			sum += gameID
		}
	}
	return sum
}

func solveExample2() int {
	return puzzle2Helper("example2.txt")
}

func solvePuzzle2() int {
	return puzzle2Helper("puzzle.txt")
}

func puzzle2Helper(fileName string) int {
	input, _ := os.ReadFile(path.Join("inputs", fileName))
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		lineMatches := lineRE.FindStringSubmatch(line)
		results := lineMatches[2]
		rounds := strings.Split(results, "; ")
		minDraw := draw{}
		for _, round := range rounds {
			roundDraw := draw{}
			hands := strings.Split(round, ", ")
			for _, hand := range hands {
				handMatches := drawRE.FindStringSubmatch(hand)
				color := handMatches[2]
				count, _ := strconv.Atoi(handMatches[1])
				switch color {
				case "red":
					roundDraw.red += count
				case "blue":
					roundDraw.blue += count
				case "green":
					roundDraw.green += count
				}
			}
			if roundDraw.red > minDraw.red {
				minDraw.red = roundDraw.red
			}
			if roundDraw.blue > minDraw.blue {
				minDraw.blue = roundDraw.blue
			}
			if roundDraw.green > minDraw.green {
				minDraw.green = roundDraw.green
			}
		}
		sum += minDraw.red * minDraw.blue * minDraw.green
	}
	return sum
}
