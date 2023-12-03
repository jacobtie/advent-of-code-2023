package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"regexp"
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

var lineRE = regexp.MustCompile("^Card +[0-9]+: (.*)$")

func puzzle1Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read file: %s", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		matches := lineRE.FindStringSubmatch(line)
		lists := strings.Split(matches[1], " | ")
		firstList := strings.Split(lists[0], " ")
		secondList := strings.Split(lists[1], " ")
		matchedNums := make(map[string]struct{})
		for _, num1 := range firstList {
			if num1 == "" {
				continue
			}
			for _, num2 := range secondList {
				if num1 == num2 {
					matchedNums[num1] = struct{}{}
				}
			}
		}
		sum += int(math.Pow(2, float64(len(matchedNums)-1)))
	}
	return sum
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
		log.Fatalf("failed to read file: %s", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	frontier := make([]int, 0)
	processed := 0
	memo := make(map[int]int)
	for i := range lines {
		frontier = append(frontier, i)
	}
	for len(frontier) > 0 {
		currentCardIdx := frontier[0]
		frontier = frontier[1:]
		processed += 1
		if _, ok := memo[currentCardIdx]; !ok {
			matches := lineRE.FindStringSubmatch(lines[currentCardIdx])
			lists := strings.Split(matches[1], " | ")
			firstList := strings.Split(lists[0], " ")
			secondList := strings.Split(lists[1], " ")
			matchedNums := make(map[string]struct{})
			for _, num1 := range firstList {
				if num1 == "" {
					continue
				}
				for _, num2 := range secondList {
					if num1 == num2 {
						matchedNums[num1] = struct{}{}
					}
				}
			}
			memo[currentCardIdx] = len(matchedNums)
		}
		matchedNumsCount := memo[currentCardIdx]
		for i := currentCardIdx + 1; i < currentCardIdx+matchedNumsCount+1; i++ {
			frontier = append(frontier, i)
		}
	}
	return processed
}
