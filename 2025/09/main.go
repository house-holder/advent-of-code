package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type Line struct {
	start, end Coord
	pos        int
	isVert     bool
}

type Corners struct {
	topL, topR, botR, botL Coord
}

type Rect struct {
	corners Corners
	edges   [4]Line
}

func NewLine(a, b Coord) Line {
	startPt, endPt, vertical, position := a, b, false, a.Y
	if a.X == b.X {
		position, vertical = a.X, true
		if a.Y > b.Y {
			startPt, endPt = b, a
		}
	} else {
		if a.X > b.X {
			startPt, endPt = b, a
		}
	}
	return Line{
		start:  startPt,
		end:    endPt,
		pos:    position,
		isVert: vertical,
	}
}

func NewRect(a, b Coord) Rect {
	tL, tR := a, Coord{b.X, a.Y}
	bL, bR := Coord{a.X, b.Y}, b
	top, right := NewLine(tL, tR), NewLine(tR, bR)
	bottom, left := NewLine(bR, bL), NewLine(bL, tL)
	return Rect{
		corners: Corners{
			topL: tL, topR: tR, botR: bR, botL: bL,
		},
		edges: [4]Line{top, right, bottom, left},
	}
}

func getCoords(input string) Coord {
	strCoord := strings.Split(input, ",")
	valX, err := strconv.Atoi(strCoord[0])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}
	valY, err := strconv.Atoi(strCoord[1])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}

	return Coord{
		X: valX,
		Y: valY,
	}
}

func convertCoords(input string) []Coord {
	output := []Coord{}
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) > 0 {
			newCoords := getCoords(line)
			output = append(output, newCoords)
		}
	}
	return output
}

func findArea(topL Coord, botR Coord) int {
	width := 1 + botR.X - topL.X
	height := 1 + botR.Y - topL.Y
	return width * height
}

func (a *Line) intersects(b Line) bool {
	if a.isVert == b.isVert {
		return false
	}
	vLine, hLine := *a, b
	if !a.isVert {
		vLine, hLine = b, *a
	}
	// if vLine.pos > hLine.start.X && vLine.pos < hLine.end.X &&
	// 	hLine.pos > vLine.start.Y && hLine.pos < vLine.end.Y {
	// 	return true
	// }
	// fmt.Printf("  %v and %v\n", *a, b)
	if vLine.pos > hLine.start.X {
		if vLine.pos < hLine.end.X {
			if hLine.pos > vLine.start.Y {
				if hLine.pos < vLine.end.Y {
					return true
				}
			}
		}
	}
	return false
}

func buildBounds(coords []Coord) []Line {
	bounds := []Line{}
	for i := range len(coords) - 1 {
		if i < len(coords) {
			line := NewLine(coords[i], coords[i+1])
			bounds = append(bounds, line)
		}
	}
	line := NewLine(coords[len(coords)-1], coords[0])
	bounds = append(bounds, line)
	return bounds
}

func getSafeArea(a Coord, b Coord, bounds []Line) int {
	r := NewRect(a, b)
	for _, edge := range r.edges {
		for _, line := range bounds {
			if line.intersects(edge) {
				fmt.Println(">>>intersect")
			}
		}
	}
	return findArea(r.corners.topL, r.corners.botR)
}

func evalPart2(coords []Coord) int {
	maxArea := 0
	bounds := buildBounds(coords)

	for i, a := range coords {
		for j, b := range coords {
			if i == j {
				continue
			}
			area := getSafeArea(a, b, bounds)
			fmt.Println("Area:", area)
			if area == 0 {
				continue
			}
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func evalPart1(coords []Coord) int {
	maxArea := 0
	for i, cornerA := range coords {
		for j, cornerB := range coords {
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

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	coords := convertCoords(body)

	result1 := evalPart1(coords)
	result2 := evalPart2(coords)
	fmt.Println()
	fmt.Printf("Part 1: %d\n", result1)
	fmt.Printf("Part 2: %d\n", result2)
}
