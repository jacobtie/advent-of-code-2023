package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type Node struct {
	Symbol   rune
	Adjacent []*Node
	IsLoop   bool
	Row      int
	Col      int
}

func (n *Node) String() string {
	adjSymbols := make([]string, 0)
	for _, adj := range n.Adjacent {
		adjSymbols = append(adjSymbols, string(adj.Symbol))
	}
	return fmt.Sprintf("%c: {[%s], %t}", n.Symbol, strings.Join(adjSymbols, ", "), n.IsLoop)
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
	return puzzle1Helper("example1.txt")
}

func runPuzzle1() int {
	return puzzle1Helper("puzzle.txt")
}

func puzzle1Helper(fileName string) int {
	input, err := os.ReadFile(path.Join("inputs", fileName))
	if err != nil {
		log.Fatalf("failed to read file %s: %s\n", fileName, err.Error())
	}
	lines := strings.Split(string(input), "\n")
	grid := make([][]*Node, 0)
	for i, line := range lines {
		grid = append(grid, make([]*Node, 0))
		for _, char := range line {
			grid[i] = append(grid[i], &Node{Symbol: char, Adjacent: make([]*Node, 0)})
		}
	}
	var sNode *Node
	for i := range grid {
		for j := range grid[i] {
			currNode := grid[i][j]
			if currNode.Symbol == 'S' {
				sNode = currNode
			}
			if strings.Contains("S|LJ", string(currNode.Symbol)) && i != 0 {
				aboveNode := grid[i-1][j]
				switch aboveNode.Symbol {
				case '|':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				case '7':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				case 'F':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				}
			}
			if strings.Contains("S-J7", string(currNode.Symbol)) && j != 0 {
				leftNode := grid[i][j-1]
				switch leftNode.Symbol {
				case '-':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				case 'L':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				case 'F':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				}
			}
			if strings.Contains("S-LF", string(currNode.Symbol)) && j != len(grid[i])-1 {
				rightNode := grid[i][j+1]
				switch rightNode.Symbol {
				case '-':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				case '7':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				case 'J':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				}
			}
			if strings.Contains("S|7F", string(currNode.Symbol)) && i != len(grid)-1 {
				belowNode := grid[i+1][j]
				switch belowNode.Symbol {
				case '|':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				case 'L':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				case 'J':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				}
			}
		}
	}
	var steps int
	for _, nextNode := range sNode.Adjacent {
		var foundGoal bool
		steps, foundGoal = visitNode1(nextNode, sNode, 1)
		if foundGoal {
			break
		}
	}
	return steps / 2
}

func visitNode1(currNode, parentNode *Node, numSteps int) (int, bool) {
	if currNode.Symbol == 'S' {
		return numSteps, true
	}
	for _, nextNode := range currNode.Adjacent {
		if nextNode == parentNode {
			continue
		}
		steps, foundGoal := visitNode1(nextNode, currNode, numSteps+1)
		if foundGoal {
			return steps, true
		}
	}
	return 0, false
}

func runExample2() int {
	return puzzle2Helper("example2.txt")
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
	grid := make([][]*Node, 0)
	lineNum := 0
	for i, line := range lines {
		grid = append(grid, make([]*Node, 0))
		colNum := 0
		for j, char := range line {
			grid[lineNum] = append(grid[lineNum], &Node{Symbol: char, Adjacent: make([]*Node, 0), IsLoop: false, Row: lineNum, Col: colNum})
			colNum += 1
			if j != len(line)-1 {
				grid[lineNum] = append(grid[lineNum], &Node{Symbol: 'X', Adjacent: make([]*Node, 0), IsLoop: false, Row: lineNum, Col: colNum})
				colNum += 1
			}
		}
		lineNum += 1
		if i != len(lines)-1 {
			grid = append(grid, make([]*Node, 0))
			colNum := 0
			for j := 0; j < len(line); j++ {
				grid[lineNum] = append(grid[lineNum], &Node{Symbol: 'X', Adjacent: make([]*Node, 0), IsLoop: false, Row: lineNum, Col: colNum})
				colNum += 1
				if j != len(line)-1 {
					grid[lineNum] = append(grid[lineNum], &Node{Symbol: 'X', Adjacent: make([]*Node, 0), IsLoop: false, Row: lineNum, Col: colNum})
					colNum += 1
				}
			}
			lineNum += 1
		}
	}
	var sNode *Node
	for i := range grid {
		if i%2 == 1 {
			continue
		}
		for j := range grid[i] {
			if j%2 == 1 {
				continue
			}
			currNode := grid[i][j]
			if currNode.Symbol == 'S' {
				sNode = currNode
			}
			if strings.Contains("S|LJ", string(currNode.Symbol)) && i != 0 {
				aboveNode := grid[i-2][j]
				switch aboveNode.Symbol {
				case '|':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				case '7':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				case 'F':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, aboveNode)
				}
			}
			if strings.Contains("S-J7", string(currNode.Symbol)) && j != 0 {
				leftNode := grid[i][j-2]
				switch leftNode.Symbol {
				case '-':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				case 'L':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				case 'F':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, leftNode)
				}
			}
			if strings.Contains("S-LF", string(currNode.Symbol)) && j != len(grid[i])-1 {
				rightNode := grid[i][j+2]
				switch rightNode.Symbol {
				case '-':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				case '7':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				case 'J':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, rightNode)
				}
			}
			if strings.Contains("S|7F", string(currNode.Symbol)) && i != len(grid)-1 {
				belowNode := grid[i+2][j]
				switch belowNode.Symbol {
				case '|':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				case 'L':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				case 'J':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				case 'S':
					currNode.Adjacent = append(currNode.Adjacent, belowNode)
				}
			}
		}
	}
	sNode.IsLoop = true
	for _, nextNode := range sNode.Adjacent {
		visitNode2(nextNode, sNode, grid)
	}
	numContained := 0
	for i, row := range grid {
		if i%2 == 1 {
			continue
		}
		for j, item := range row {
			if j%2 == 1 {
				continue
			}
			visited := make(map[*Node]struct{})
			if !item.IsLoop && isContained(item, grid, visited) {
				numContained += 1
			}
		}
	}
	return numContained
}

func visitNode2(currNode, parentNode *Node, grid [][]*Node) bool {
	if currNode.Symbol == 'S' {
		return true
	}
	for _, nextNode := range currNode.Adjacent {
		if nextNode == parentNode {
			continue
		}
		foundGoal := visitNode2(nextNode, currNode, grid)
		if foundGoal {
			currNode.IsLoop = true
			var inBetweenRow int
			var inBetweenCol int
			if currNode.Row == nextNode.Row {
				inBetweenRow = currNode.Row
			}
			if currNode.Row > nextNode.Row {
				inBetweenRow = currNode.Row - 1
			}
			if currNode.Row < nextNode.Row {
				inBetweenRow = currNode.Row + 1
			}
			if currNode.Col == nextNode.Col {
				inBetweenCol = currNode.Col
			}
			if currNode.Col > nextNode.Col {
				inBetweenCol = currNode.Col - 1
			}
			if currNode.Col < nextNode.Col {
				inBetweenCol = currNode.Col + 1
			}
			grid[inBetweenRow][inBetweenCol].Symbol = 'Y'
			grid[inBetweenRow][inBetweenCol].IsLoop = true
			return true
		}
	}
	return false
}

func isContained(node *Node, grid [][]*Node, visited map[*Node]struct{}) bool {
	visited[node] = struct{}{}
	if node.Row == 0 {
		return false
	}
	if node.Row == len(grid)-1 {
		return false
	}
	if node.Col == 0 {
		return false
	}
	if node.Col == len(grid[0])-1 {
		return false
	}
	above := grid[node.Row-1][node.Col]
	if _, ok := visited[above]; !ok && !above.IsLoop && !isContained(above, grid, visited) {
		return false
	}
	below := grid[node.Row+1][node.Col]
	if _, ok := visited[below]; !ok && !below.IsLoop && !isContained(below, grid, visited) {
		return false
	}
	left := grid[node.Row][node.Col-1]
	if _, ok := visited[left]; !ok && !left.IsLoop && !isContained(left, grid, visited) {
		return false
	}
	right := grid[node.Row][node.Col+1]
	if _, ok := visited[right]; !ok && !right.IsLoop && !isContained(right, grid, visited) {
		return false
	}
	return true
}
