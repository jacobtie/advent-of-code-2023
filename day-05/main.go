package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type MappingNode struct {
	DestStart   int
	SourceStart int
	Length      int
}

type Mapping struct {
	nodes []MappingNode
}

func (m *Mapping) Add(destStart, sourceStart, length int) {
	m.nodes = append(m.nodes, MappingNode{
		DestStart:   destStart,
		SourceStart: sourceStart,
		Length:      length,
	})
}

func (m *Mapping) Process() {
	sort.Slice(m.nodes, func(i, j int) bool {
		return m.nodes[i].SourceStart < m.nodes[j].SourceStart
	})
}

func (m *Mapping) Lookup(sourceNum int) int {
	var lastNode MappingNode
	lastNode.SourceStart = -1
	for _, node := range m.nodes {
		if node.SourceStart > sourceNum {
			break
		}
		lastNode = node
	}
	if lastNode.SourceStart == -1 {
		return sourceNum
	}
	delta := sourceNum - lastNode.SourceStart
	if delta > lastNode.Length {
		return sourceNum
	}
	return lastNode.DestStart + delta
}

func NewMapping() *Mapping {
	return &Mapping{nodes: make([]MappingNode, 0)}
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
		log.Fatalf("failed to read file: %s", err.Error())
	}
	lines := strings.Split(string(input), "\n")
	var seeds []int
	seedToSoil := NewMapping()
	soilToFertilizer := NewMapping()
	fertilizerToWater := NewMapping()
	waterToLight := NewMapping()
	lightToTemperature := NewMapping()
	temperatureToHumidity := NewMapping()
	humidityToLocation := NewMapping()
	for i := range lines {
		if strings.HasPrefix(lines[i], "seeds: ") {
			segments := strings.Split(lines[i], ": ")
			seeds = getNums(segments[1])
			continue
		}
		if lines[i] == "seed-to-soil map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				seedToSoil.Add(nums[0], nums[1], nums[2])
			}
			seedToSoil.Process()
		}
		if lines[i] == "soil-to-fertilizer map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				soilToFertilizer.Add(nums[0], nums[1], nums[2])
			}
			soilToFertilizer.Process()
		}
		if lines[i] == "fertilizer-to-water map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				fertilizerToWater.Add(nums[0], nums[1], nums[2])
			}
			fertilizerToWater.Process()
		}
		if lines[i] == "water-to-light map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				waterToLight.Add(nums[0], nums[1], nums[2])
			}
			waterToLight.Process()
		}
		if lines[i] == "light-to-temperature map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				lightToTemperature.Add(nums[0], nums[1], nums[2])
			}
			lightToTemperature.Process()
		}
		if lines[i] == "temperature-to-humidity map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				temperatureToHumidity.Add(nums[0], nums[1], nums[2])
			}
			temperatureToHumidity.Process()
		}
		if lines[i] == "humidity-to-location map:" {
			for {
				i += 1
				if i == len(lines) || lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				humidityToLocation.Add(nums[0], nums[1], nums[2])
			}
			humidityToLocation.Process()
		}
	}
	minLocation := math.MaxInt
	for _, seed := range seeds {
		soil := seedToSoil.Lookup(seed)
		fertilizer := soilToFertilizer.Lookup(soil)
		water := fertilizerToWater.Lookup(fertilizer)
		light := waterToLight.Lookup(water)
		temperature := lightToTemperature.Lookup(light)
		humidity := temperatureToHumidity.Lookup(temperature)
		location := humidityToLocation.Lookup(humidity)
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

func getNums(line string) []int {
	nums := make([]int, 0)
	numsStr := strings.Split(line, " ")
	for _, numStr := range numsStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("failed to convert %s: %s", numStr, err.Error())
		}
		nums = append(nums, num)
	}
	return nums
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
	seeds := make([][]int, 0)
	seedToSoil := NewMapping()
	soilToFertilizer := NewMapping()
	fertilizerToWater := NewMapping()
	waterToLight := NewMapping()
	lightToTemperature := NewMapping()
	temperatureToHumidity := NewMapping()
	humidityToLocation := NewMapping()
	for i := range lines {
		if strings.HasPrefix(lines[i], "seeds: ") {
			segments := strings.Split(lines[i], ": ")
			for j, numStr := range strings.Split(segments[1], " ") {
				idx := j / 2
				if j%2 == 0 {
					seeds = append(seeds, make([]int, 2))
				}
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatalf("failed to convert %s: %s", numStr, err.Error())
				}
				seeds[idx][j%2] = num
			}
			continue
		}
		if lines[i] == "seed-to-soil map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				seedToSoil.Add(nums[0], nums[1], nums[2])
			}
			seedToSoil.Process()
		}
		if lines[i] == "soil-to-fertilizer map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				soilToFertilizer.Add(nums[0], nums[1], nums[2])
			}
			soilToFertilizer.Process()
		}
		if lines[i] == "fertilizer-to-water map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				fertilizerToWater.Add(nums[0], nums[1], nums[2])
			}
			fertilizerToWater.Process()
		}
		if lines[i] == "water-to-light map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				waterToLight.Add(nums[0], nums[1], nums[2])
			}
			waterToLight.Process()
		}
		if lines[i] == "light-to-temperature map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				lightToTemperature.Add(nums[0], nums[1], nums[2])
			}
			lightToTemperature.Process()
		}
		if lines[i] == "temperature-to-humidity map:" {
			for {
				i += 1
				if lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				temperatureToHumidity.Add(nums[0], nums[1], nums[2])
			}
			temperatureToHumidity.Process()
		}
		if lines[i] == "humidity-to-location map:" {
			for {
				i += 1
				if i == len(lines) || lines[i] == "" {
					break
				}
				nums := getNums(lines[i])
				humidityToLocation.Add(nums[0], nums[1], nums[2])
			}
			humidityToLocation.Process()
		}
	}
	minLocation := math.MaxInt
	for _, seedGroup := range seeds {
		for i := seedGroup[0]; i < seedGroup[0]+seedGroup[1]; i++ {
			soil := seedToSoil.Lookup(i)
			fertilizer := soilToFertilizer.Lookup(soil)
			water := fertilizerToWater.Lookup(fertilizer)
			light := waterToLight.Lookup(water)
			temperature := lightToTemperature.Lookup(light)
			humidity := temperatureToHumidity.Lookup(temperature)
			location := humidityToLocation.Lookup(humidity)
			if location < minLocation {
				minLocation = location
			}
		}
	}
	return minLocation
}
