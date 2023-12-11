package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"strings"
)

type Node struct {
	Symbol   rune
	IsGalaxy bool
	Row      int
	Col      int
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

func puzzle1Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read %s: %s\n", fileName, err.Error())
	}
	lines := strings.Split(string(input), "\n")
	originalGrid := make([][]*Node, 0)
	galaxies := make([]*Node, 0)
	rowHasGalaxy := make([]bool, len(lines))
	columnHasGalaxy := make([]bool, len(lines[0]))
	for i, row := range lines {
		originalGrid = append(originalGrid, make([]*Node, 0))
		for j, item := range row {
			isGalaxy := false
			if item == '#' {
				isGalaxy = true
				rowHasGalaxy[i] = true
				columnHasGalaxy[j] = true
			}
			node := &Node{Symbol: item, IsGalaxy: isGalaxy}
			originalGrid[i] = append(originalGrid[i], node)
			if isGalaxy {
				galaxies = append(galaxies, node)
			}
		}
	}
	expandedGrid := make([][]*Node, 0)
	rowNum := 0
	for i, row := range originalGrid {
		expandedGrid = append(expandedGrid, make([]*Node, 0))
		colNum := 0
		for j, node := range row {
			node.Row = rowNum
			node.Col = colNum
			expandedGrid[rowNum] = append(expandedGrid[rowNum], node)
			colNum += 1
			if !columnHasGalaxy[j] {
				expandedGrid[rowNum] = append(expandedGrid[rowNum], &Node{Symbol: '.', IsGalaxy: false, Row: rowNum, Col: colNum})
				colNum += 1
			}
		}
		rowNum += 1
		if !rowHasGalaxy[i] {
			newRow := make([]*Node, 0, len(expandedGrid[rowNum-1]))
			colNum := 0
			for j := 0; j < len(expandedGrid[rowNum-1]); j++ {
				newRow = append(newRow, &Node{Symbol: '.', IsGalaxy: false, Row: rowNum, Col: colNum})
				colNum += 1
			}
			expandedGrid = append(expandedGrid, newRow)
			rowNum += 1
		}
	}
	sum := 0
	for _, galaxyA := range galaxies {
		for _, galaxyB := range galaxies {
			if galaxyA == galaxyB {
				continue
			}
			sum += int(math.Abs(float64(galaxyA.Row)-float64(galaxyB.Row)) + math.Abs(float64(galaxyA.Col)-float64(galaxyB.Col)))
		}
	}
	return sum / 2
}

func runExample2() int64 {
	return puzzle2Helper("example.txt")
}

func runPuzzle2() int64 {
	return puzzle2Helper("puzzle.txt")
}

func puzzle2Helper(fileName string) int64 {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read %s: %s\n", fileName, err.Error())
	}
	lines := strings.Split(string(input), "\n")
	originalGrid := make([][]*Node, 0)
	galaxies := make([]*Node, 0)
	rowHasGalaxy := make([]bool, len(lines))
	columnHasGalaxy := make([]bool, len(lines[0]))
	for i, row := range lines {
		originalGrid = append(originalGrid, make([]*Node, 0))
		for j, item := range row {
			isGalaxy := false
			if item == '#' {
				isGalaxy = true
				rowHasGalaxy[i] = true
				columnHasGalaxy[j] = true
			}
			node := &Node{Symbol: item, IsGalaxy: isGalaxy}
			originalGrid[i] = append(originalGrid[i], node)
			if isGalaxy {
				galaxies = append(galaxies, node)
			}
		}
	}
	expandedGrid := make([][]*Node, 0)
	rowNum := 0
	for i, row := range originalGrid {
		expandedGrid = append(expandedGrid, make([]*Node, 0))
		colNum := 0
		for j, node := range row {
			node.Row = rowNum
			node.Col = colNum
			if !columnHasGalaxy[j] {
				colNum += 999_999
			}
			expandedGrid[i] = append(expandedGrid[i], node)
			colNum += 1
		}
		rowNum += 1
		if !rowHasGalaxy[i] {
			rowNum += 999_999
		}
	}
	var sum int64 = 0
	for _, galaxyA := range galaxies {
		for _, galaxyB := range galaxies {
			if galaxyA == galaxyB {
				continue
			}
			sum += int64(math.Abs(float64(galaxyA.Row)-float64(galaxyB.Row)) + math.Abs(float64(galaxyA.Col)-float64(galaxyB.Col)))
		}
	}
	return sum / 2
}
