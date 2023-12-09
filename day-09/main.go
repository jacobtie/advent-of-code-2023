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

type ListsContainer struct {
	lists [][]int
}

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

var numRE = regexp.MustCompile("-?[0-9]+")

func puzzle1Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read file %s: %s\n", fileName, err.Error())
	}
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		matches := numRE.FindAllString(line, -1)
		nums := make([]int, 0)
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				log.Fatalf("failed to convert %s to a number: %s\n", match, err.Error())
			}
			nums = append(nums, num)
		}
		sum += extrapolateNext(nums)
	}
	return sum
}

func extrapolateNext(nums []int) int {
	container := &ListsContainer{[][]int{nums}}
	extrapolateNextHelper(container, 1)
	currentLast := container.lists[0][len(container.lists[0])-1]
	nextRowLast := container.lists[1][len(container.lists[1])-1]
	return currentLast + nextRowLast
}

func extrapolateNextHelper(container *ListsContainer, depth int) {
	newList := make([]int, 0, len(container.lists[depth-1]))
	for i := 0; i < len(container.lists[depth-1])-1; i++ {
		newList = append(newList, container.lists[depth-1][i+1]-container.lists[depth-1][i])
	}
	allZeroes := true
	for _, num := range newList {
		if num != 0 {
			allZeroes = false
			break
		}
	}
	container.lists = append(container.lists, newList)
	if allZeroes {
		container.lists[depth] = append(container.lists[depth], 0)
		return
	}
	extrapolateNextHelper(container, depth+1)
	currentLast := container.lists[depth][len(container.lists[depth])-1]
	nextRowLast := container.lists[depth+1][len(container.lists[depth+1])-1]
	container.lists[depth] = append(container.lists[depth], currentLast+nextRowLast)
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
		log.Fatalf("failed to read file %s: %s\n", fileName, err.Error())
	}
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		matches := numRE.FindAllString(line, -1)
		nums := make([]int, 0)
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				log.Fatalf("failed to convert %s to a number: %s\n", match, err.Error())
			}
			nums = append(nums, num)
		}
		sum += extrapolatePrev(nums)
	}
	return sum
}

func extrapolatePrev(nums []int) int {
	container := &ListsContainer{[][]int{nums}}
	extrapolatePrevHelper(container, 1)
	currentFirst := container.lists[0][0]
	nextRowFirst := container.lists[1][0]
	return currentFirst - nextRowFirst
}

func extrapolatePrevHelper(container *ListsContainer, depth int) {
	newList := make([]int, 0, len(container.lists[depth-1]))
	for i := 0; i < len(container.lists[depth-1])-1; i++ {
		newList = append(newList, container.lists[depth-1][i+1]-container.lists[depth-1][i])
	}
	allZeroes := true
	for _, num := range newList {
		if num != 0 {
			allZeroes = false
			break
		}
	}
	container.lists = append(container.lists, newList)
	if allZeroes {
		container.lists[depth] = append([]int{0}, container.lists[depth]...)
		return
	}
	extrapolatePrevHelper(container, depth+1)
	currentFirst := container.lists[depth][0]
	nextRowFirst := container.lists[depth+1][0]
	container.lists[depth] = append([]int{currentFirst - nextRowFirst}, container.lists[depth]...)
}
