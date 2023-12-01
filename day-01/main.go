package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

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
	sum := 0
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		firstChar := ""
		lastChar := ""
		for _, char := range line {
			if strings.Contains("123456789", string(char)) {
				if firstChar == "" {
					firstChar = string(char)
				}
				lastChar = string(char)
			}
		}
		fullNumStr := firstChar + lastChar
		fullNum, _ := strconv.Atoi(fullNumStr)
		sum += fullNum
	}
	return sum
}

func solveExample2() int {
	return puzzle2Helper("example2.txt")
}

func solvePuzzle2() int {
	return puzzle2Helper("puzzle.txt")
}

var numsAsStrs = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func puzzle2Helper(fileName string) int {
	input, _ := os.ReadFile(path.Join("inputs", fileName))
	sum := 0
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		firstChar := ""
		lastChar := ""
		for i, char := range line {
			if strings.Contains("123456789", string(char)) {
				if firstChar == "" {
					firstChar = string(char)
				}
				lastChar = string(char)
				continue
			}
			for numAsStr, num := range numsAsStrs {
				if i+len(numAsStr) > len(line) {
					continue
				}
				if line[i:i+len(numAsStr)] == numAsStr {
					if firstChar == "" {
						firstChar = string(num)
					}
					lastChar = string(num)
				}
			}
		}
		fullNumStr := firstChar + lastChar
		fullNum, _ := strconv.Atoi(fullNumStr)
		sum += fullNum
	}
	return sum
}
