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

func (c *Coordinate) isWithin(input []Coordinate) bool {
	return slices.Contains(input, *c)
}

func printMap(
	reds []Coordinate,
	greens []Coordinate,
	hiX int,
) {
	for y := range len(reds) + 1 {
		fmt.Printf("\n")
		for x := range hiX + 3 {
			coord := Coordinate{x, y}
			if coord.isWithin(reds) {
				fmt.Printf("%s#%s", RED, NC)
			} else if coord.isWithin(greens) {
				fmt.Printf("%sX%s", GRN, NC)
			} else {
				fmt.Printf(".")
			}
		}
	}
	fmt.Printf("\n")
}

func findLimits(input []Coordinate) ([2]int, [2]int) {
	minX, maxX := input[0].X, input[0].X
	minY, maxY := input[0].Y, input[0].Y
	for _, coord := range input {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}
	return [2]int{minX, maxX}, [2]int{minY, maxY}
}

func extractGreensFrom(reds []Coordinate) []Coordinate {
	greens := []Coordinate{}
	prev := Coordinate{}

	for i, curr := range reds {
		if i > 0 {
			prev = reds[i-1]
			if curr.Y == prev.Y { // if same line
				// fmt.Printf("%d. curr=%d, currY=%d\n", i, curr.X, curr.Y)
				// fmt.Printf("   prevX=%d, prevY=%d\n", prev.X, prev.Y)
				bigX := int(math.Max(float64(curr.X), float64(prev.X)))
				litX := int(math.Min(float64(curr.X), float64(prev.X)))

				for i := litX + 1; i < bigX; i++ {
					greens = append(greens, Coordinate{i, curr.Y})
					// fmt.Printf("     append %d, %d\n", i, curr.Y)
				}
			} else {
				continue
			}
		}
	}
	return greens
}

func infill(
	greens []Coordinate,
	reds []Coordinate,
	xLims [2]int,
	yLims [2]int,
) []Coordinate {
	rows := [][]Coordinate{}
	// build uniform-length rows encompassing color field
	//

	// for y := yLims[0]; y <= yLims[1]; y++ {
	// 	for x := xLims[0]; x <= xLims[1]; x++ {
	//            current := Coordinate{x, y}
	// 		prevRedLine := reds[y-1]
	// 		prevGreenLine := greens[y-1]
	//            if (current.isWithin)
	// 			greens = append(greens, current)
	// 		}
	// 	}
	// }

	return greens
}

func evalPart2(reds []Coordinate, hiX int) int {
	xLims, yLims := findLimits(reds)
	greens := extractGreensFrom(reds)
	greens = infill(reds, greens, xLims, yLims)
	printMap(reds, greens, hiX)

	return len(greens)
}

// start 1765470289 stop 1765472304
// start 1765479977 stop 1765483078
func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	coords, hiX := convertCoords(body)

	result1 := evalPart1(coords)
	result2 := evalPart2(coords, hiX)
	fmt.Println()
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
