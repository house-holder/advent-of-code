package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	RED = "\033[31m"
	GRN = "\033[32m"
	NC  = "\033[0m"
)

type Coordinate struct {
	X int
	Y int
}

func getCoords(input string) Coordinate {
	strPair := strings.Split(input, ",")
	valX, err := strconv.Atoi(strPair[0])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}
	valY, err := strconv.Atoi(strPair[1])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}

	return Coordinate{
		X: valX,
		Y: valY,
	}
}

func maxPossibleArea(input []Coordinate) int {
	maxArea := 0
	for i, cornerA := range input {
		for j, cornerB := range input {
			if i != j {
				thisArea := findArea(cornerA, cornerB)
				if maxArea < thisArea {
					maxArea = thisArea
				}
			}
		}
	}
	return maxArea
}

func findArea(cornerA Coordinate, cornerB Coordinate) int {
	latSize := 1 + cornerB.X - cornerA.X
	vertSize := 1 + cornerB.Y - cornerA.Y
	return int(math.Abs(float64(latSize * vertSize)))
}

func convertCoords(input string) ([]Coordinate, int) {
	output := []Coordinate{}
	hiX := 0
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) > 0 {
			newCoords := getCoords(line)
			output = append(output, newCoords)
			if math.Abs(float64(newCoords.X)) > float64(hiX) {
				hiX = newCoords.X
			}
		}
	}
	return output, hiX
}

func evalPart1(input []Coordinate) int {
	return maxPossibleArea(input)
}

func (c *Coordinate) isRed(input []Coordinate) bool {
	return slices.Contains(input, *c)
}

func (c *Coordinate) isGreen(input []Coordinate) bool {
	return false
}

func printMap(input []Coordinate, hiX int) {
	reds := []Coordinate{}
	greens := []Coordinate{}
	for y := range len(input) + 1 {
		fmt.Printf("\n")
		for x := range hiX + 3 {
			coord := Coordinate{x, y}
			if coord.isRed(input) {
				fmt.Printf("%s#%s", RED, NC)
				reds = append(reds, coord)
			} else if coord.isGreen(input) {
				fmt.Printf("%s#%s", RED, NC)
				greens = append(greens, coord)
			} else {
				fmt.Printf(".")
			}
		}
	}
	fmt.Printf("\n")
	fmt.Println(greens)
}

func evalPart2(input []Coordinate, hiX int) int {
	printMap(input, hiX)

	return len(input)
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	coords, hiX := convertCoords(body)

	fmt.Println()

	result1 := evalPart1(coords)
	result2 := evalPart2(coords, hiX)
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
