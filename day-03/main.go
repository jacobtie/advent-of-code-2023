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

func puzzle1Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read file: %s", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	grid := make([][]string, len(lines))
	for i := range lines {
		grid[i] = strings.Split(lines[i], "")
	}
	sum := 0
	for i := range grid {
		currNumLength := 0
		for j := 0; j < len(grid[i])+1; j++ {
			if j == len(grid[i]) || !isInt(grid[i][j]) {
				if currNumLength == 0 {
					continue
				}
				numStr := strings.Join(grid[i][j-currNumLength:j], "")
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatalf("failed to convert number %s: %s", numStr, err.Error())
				}
				// Check: grid[i-1][j-currNumLength-1:j+1], grid[i][j-currNumLength-1], grid[i][j], grid[i+1][j-currNumLength-1:j+1]
				isValid := false
				if i-1 >= 0 {
					for k := j - currNumLength - 1; k < j+1; k++ {
						if k < 0 || k >= len(grid[i]) {
							continue
						}
						if isSymbol(grid[i-1][k]) {
							isValid = true
							break
						}
					}
				}
				if !isValid && j-currNumLength-1 > 0 && isSymbol(grid[i][j-currNumLength-1]) {
					isValid = true
				}
				if !isValid && j < len(grid[i]) && isSymbol(grid[i][j]) {
					isValid = true
				}
				if !isValid && i+1 < len(grid) {
					for k := j - currNumLength - 1; k < j+1; k++ {
						if k < 0 || k >= len(grid[i]) {
							continue
						}
						if isSymbol(grid[i+1][k]) {
							isValid = true
							break
						}
					}
				}
				if isValid {
					sum += num
				}
				currNumLength = 0
				continue
			}
			currNumLength += 1
		}
	}
	return sum
}

func runExample2() int {
	return puzzle2Helper("example.txt")
}

func runPuzzle2() int {
	return puzzle2Helper("puzzle.txt")
}

var numsRe = regexp.MustCompile("[0-9]+")

func puzzle2Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read file: %s", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	sum := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != '*' {
				continue
			}
			adjacentNums := make([]int, 0, 2)
			if i != 0 {
				matchedNumsAbove := numsRe.FindAllString(lines[i-1], -1)
				matchedNumAboveIdxs := numsRe.FindAllStringIndex(lines[i-1], -1)
				for k := range matchedNumAboveIdxs {
					if matchedNumAboveIdxs[k][0] > j+1 || matchedNumAboveIdxs[k][1] < j {
						continue
					}
					matchedNum, err := strconv.Atoi(matchedNumsAbove[k])
					if err != nil {
						log.Fatalf("failed to convert number %s: %s", matchedNumsAbove[k], err.Error())
					}
					adjacentNums = append(adjacentNums, matchedNum)
				}
			}
			if i != len(lines)-1 {
				matchedNumsBelow := numsRe.FindAllString(lines[i+1], -1)
				matchedNumBelowIdxs := numsRe.FindAllStringIndex(lines[i+1], -1)
				for k := range matchedNumBelowIdxs {
					if matchedNumBelowIdxs[k][0] > j+1 || matchedNumBelowIdxs[k][1] < j {
						continue
					}
					matchedNum, err := strconv.Atoi(matchedNumsBelow[k])
					if err != nil {
						log.Fatalf("failed to convert number %s: %s", matchedNumsBelow[k], err.Error())
					}
					adjacentNums = append(adjacentNums, matchedNum)
				}
			}
			matchedNumsSameRow := numsRe.FindAllString(lines[i], -1)
			matchedNumSameRowIdxs := numsRe.FindAllStringIndex(lines[i], -1)
			for k := range matchedNumSameRowIdxs {
				if matchedNumSameRowIdxs[k][0] == j+1 || matchedNumSameRowIdxs[k][1] == j {
					matchedNum, err := strconv.Atoi(matchedNumsSameRow[k])
					if err != nil {
						log.Fatalf("failed to convert number %s: %s", matchedNumsSameRow[k], err.Error())
					}
					adjacentNums = append(adjacentNums, matchedNum)
				}
			}
			if len(adjacentNums) != 2 {
				continue
			}
			sum += adjacentNums[0] * adjacentNums[1]
		}
	}
	return sum
}

func isInt(s string) bool {
	return strings.Contains("0123456789", s)
}

func isSymbol(s string) bool {
	return !strings.Contains("0123456789.", s)
}
