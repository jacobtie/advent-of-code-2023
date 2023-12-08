package main

import (
	"fmt"
	"log"
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
	return puzzle1Helper("example1.txt")
}

func runPuzzle1() int {
	return puzzle1Helper("puzzle.txt")
}

var nodeRE = regexp.MustCompile("([0-9A-Z]{3})")
var stepToIdx = map[rune]int{'L': 0, 'R': 1}

func puzzle1Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read input file: %s\n", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	steps := lines[0]
	lines = lines[2:]
	adjacencyMap := make(map[string][]string)
	for _, line := range lines {
		matches := nodeRE.FindAllString(line, -1)
		adjacencyMap[matches[0]] = []string{matches[1], matches[2]}
	}
	numSteps := 0
	foundGoal := false
	currentNode := "AAA"
	for !foundGoal {
		for _, step := range steps {
			numSteps += 1
			currentNode = adjacencyMap[currentNode][stepToIdx[step]]
			if currentNode == "ZZZ" {
				foundGoal = true
				break
			}
		}
	}
	return numSteps
}

func runExample2() int64 {
	return puzzle2Helper("example2.txt")
}

func runPuzzle2() int64 {
	return puzzle2Helper("puzzle.txt")
}

func puzzle2Helper(fileName string) int64 {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read input file: %s\n", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	steps := lines[0]
	lines = lines[2:]
	adjacencyMap := make(map[string][]string)
	currentNodes := make([]string, 0)
	for _, line := range lines {
		matches := nodeRE.FindAllString(line, -1)
		adjacencyMap[matches[0]] = []string{matches[1], matches[2]}
		if strings.HasSuffix(matches[0], "A") {
			currentNodes = append(currentNodes, matches[0])
		}
	}
	nodeSteps := make([]int, len(currentNodes))
	for i := range currentNodes {
		foundGoal := false
		for !foundGoal {
			for _, step := range steps {
				nodeSteps[i] += 1
				currentNodes[i] = adjacencyMap[currentNodes[i]][stepToIdx[step]]
				if strings.HasSuffix(currentNodes[i], "Z") {
					foundGoal = true
					break
				}
			}
		}
	}
	lcm := getLCM(int64(nodeSteps[0]), int64(nodeSteps[1]))
	nodeSteps = nodeSteps[2:]
	for _, numSteps := range nodeSteps {
		lcm = getLCM(lcm, int64(numSteps))
	}
	return lcm
}

func getLCM(a, b int64) int64 {
	return (a * b) / getGCD(a, b)
}

func getGCD(a, b int64) int64 {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
