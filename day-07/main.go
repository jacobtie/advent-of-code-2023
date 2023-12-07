package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	processedHand []int
	bid           int
	score         int
}

var puzzle1CardMatch = map[rune]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var puzzle2CardMatch = map[rune]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 0,
	'Q': 12,
	'K': 13,
	'A': 14,
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
		log.Fatalf("failed to read file: %s\n", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		segments := strings.Split(line, " ")
		bid, err := strconv.Atoi(segments[1])
		if err != nil {
			log.Fatalf("failed to convert %s to a bid number: %s\n", segments[1], err.Error())
		}
		processedHand := make([]int, 0, 5)
		counts := make(map[rune]int)
		for _, char := range segments[0] {
			processedHand = append(processedHand, puzzle1CardMatch[char])
			if _, ok := counts[char]; !ok {
				counts[char] = 0
			}
			counts[char] += 1
		}
		hands = append(hands, Hand{
			processedHand: processedHand,
			bid:           bid,
			score:         getHandScore(counts),
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score != hands[j].score {
			return hands[i].score < hands[j].score
		}
		for k := range hands[i].processedHand {
			if hands[i].processedHand[k] != hands[j].processedHand[k] {
				return hands[i].processedHand[k] < hands[j].processedHand[k]
			}
		}
		return false // should never get here
	})
	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
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
		log.Fatalf("failed to read file: %s\n", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		segments := strings.Split(line, " ")
		bid, err := strconv.Atoi(segments[1])
		if err != nil {
			log.Fatalf("failed to convert %s to a bid number: %s\n", segments[1], err.Error())
		}
		processedHand := make([]int, 0, 5)
		for _, char := range segments[0] {
			processedHand = append(processedHand, puzzle2CardMatch[char])
		}
		subHands := processSubHand(segments[0])
		maxHandScore := 0
		for _, subHand := range subHands {
			counts := make(map[rune]int)
			for _, char := range subHand {
				if _, ok := counts[char]; !ok {
					counts[char] = 0
				}
				counts[char] += 1
			}
			subHandScore := getHandScore(counts)
			if subHandScore > maxHandScore {
				maxHandScore = subHandScore
			}
		}
		hands = append(hands, Hand{
			processedHand: processedHand,
			bid:           bid,
			score:         maxHandScore,
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score != hands[j].score {
			return hands[i].score < hands[j].score
		}
		for k := range hands[i].processedHand {
			if hands[i].processedHand[k] != hands[j].processedHand[k] {
				return hands[i].processedHand[k] < hands[j].processedHand[k]
			}
		}
		return false // should never get here
	})
	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}
	return sum
}

func getHandScore(counts map[rune]int) int {
	tripleCount := 0
	pairCount := 0
	for _, count := range counts {
		if count == 5 {
			return 7
		}
		if count == 4 {
			return 6
		}
		if count == 3 {
			tripleCount += 1
		}
		if count == 2 {
			pairCount += 1
		}
	}
	if tripleCount == 1 {
		if pairCount == 1 {
			return 5
		}
		return 4
	}
	if pairCount == 2 {
		return 3
	}
	if pairCount == 1 {
		return 2
	}
	return 1
}

func processSubHand(startingHand string) []string {
	frontier := []string{startingHand}
	subHands := make([]string, 0)
	for len(frontier) > 0 {
		currentHand := frontier[0]
		frontier = frontier[1:]
		if !strings.Contains(currentHand, "J") {
			subHands = append(subHands, currentHand)
			continue
		}
		jIdx := strings.Index(currentHand, "J")
		for char := range puzzle2CardMatch {
			if char == 'J' {
				continue
			}
			frontier = append(frontier, fmt.Sprintf("%s%c%s", currentHand[:jIdx], char, currentHand[jIdx+1:]))
		}
	}
	return subHands
}
