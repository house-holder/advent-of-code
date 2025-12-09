package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const ( // NOTE: just for colorizing debug output
	BLK = "\033[30m"
	RED = "\033[31m"
	GRN = "\033[32m"
	YLW = "\033[33m"
	BLU = "\033[34m"
	MGN = "\033[35m"
	CYN = "\033[36m"

	BGK = "\033[40m"
	BGR = "\033[41m"
	BGG = "\033[42m"
	BGY = "\033[43m"
	BGB = "\033[44m"
	BGM = "\033[45m"
	BGC = "\033[46m"

	NC = "\033[0m"
)

type Position struct {
	X int
	Y int
	Z int
}

type Circuit struct {
	name        string
	connections []*Box
}

type Box struct {
	name    string
	circuit *Circuit
	pos     *Position
}

func (p *Position) toString() string {
	return fmt.Sprintf("%s%3d %3d %3d%s", GRN, p.X, p.Y, p.Z, NC)
}

func (b *Box) getDistTo(o *Box) float64 {
	p, w := b.pos, o.pos
	dZ, dX, dY := float64(w.X-p.X), float64(w.Y-p.Y), float64(w.Z-p.Z)
	dist := math.Sqrt((dX * dX) + (dY * dY) + (dZ * dZ))
	return math.Round(dist * 100)
}

func (b *Box) fullName() string {
	pos := fmt.Sprintf("%s(%11s%s)%s", BLU, b.pos.toString(), BLU, NC)
	return fmt.Sprintf("%s%s%s%s%s%s", BLU, b.name, pos, YLW, b.circuit.name, NC)

}

// func (b *Box) connect(o *Box) *Circuit {
//
// }

func buildPosition(input string) *Position {
	parts := strings.Split(input, ",")
	valX, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("strconv.Atoi failed.", err)
	}
	valY, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("strconv.Atoi failed.", err)
	}
	valZ, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal("strconv.Atoi failed.", err)
	}
	return &Position{
		X: valX,
		Y: valY,
		Z: valZ,
	}
}

func buildCircuit(idx int) *Circuit {
	c := &Circuit{
		name:        fmt.Sprintf("ckt%02d", idx),
		connections: []*Box{},
	}
	return c
}

func buildBox(input string, allBoxes []*Box) *Box {
	i := len(allBoxes)
	n := fmt.Sprintf("jb%02d", i)
	box := &Box{
		name:    n,
		circuit: nil,
		pos:     buildPosition(input),
	}
	ckt := buildCircuit(i)
	ckt.connections = append(ckt.connections, box)
	box.circuit = ckt
	return box
}

func mapAllConnections(allBoxes []*Box) map[int][2]*Box {
	output := make(map[int][2]*Box)
	for _, boxA := range allBoxes {
		for _, boxB := range allBoxes {
			if boxA != boxB {
				dist := int(boxA.getDistTo(boxB))
				positionPair := [2]*Box{boxA, boxB}
				output[dist] = positionPair
			}
		}
	}
	return output
}

func (c *Circuit) getStr() string {
	conns := []string{}
	for _, cxn := range c.connections {
		conns = append(conns, cxn.name)
	}
	output := strings.Join(conns, " ")
	return fmt.Sprintf("%s: [%s]", c.name, output)
}

func debugPrintBoxPair(leader string, pair [2]*Box) {
	fmt.Printf("%s%s  ", leader, pair[0].fullName())
	fmt.Printf("%s\n", pair[1].fullName())
}

func debugPrintCircuits(input []*Circuit) {
	for _, ckt := range input {
		if len(ckt.connections) > 0 {
			fmt.Println("   ", ckt.getStr())
		}
	}
}

func getCircuitSizes(input []*Circuit) []int {
	output := []int{}
	for _, ckt := range input {
		if len(ckt.connections) != 0 {
			output = append(output, len(ckt.connections))
		}
	}
	return output
}

func connect(input [2]*Box) {
	if slices.Contains(input[0].circuit.connections, input[1]) {
		fmt.Printf("%sSKIP CASE: %s.connections contains %s%s\n",
			RED, input[0].name, input[1].name, NC)
		return
	}
	if slices.Contains(input[1].circuit.connections, input[0]) {
		fmt.Printf("%sSKIP CASE: %s.connections contains %s%s\n",
			RED, input[1].name, input[0].name, NC)
		return
	}
	cA := input[0].circuit
	cB := input[1].circuit

	if len(cA.connections) < len(cB.connections) {
		cA = input[1].circuit
		cB = input[0].circuit
	}

	for _, box := range cB.connections {
		fmt.Printf("  Target: %s", box.name)

		if !slices.Contains(cA.connections, box) {
			fmt.Printf(", added to %s\n", cA.name)
			cA.connections = append(cA.connections, box)
			continue
		}
		fmt.Printf("%sOH NO OH MY GOD HELP ME OH THE HORROR%s\n", RED, NC)
	}

	cB.connections = []*Box{}
	input[1].circuit = cA
}

func evalPart1(input string, filename string) int {
	result := 1
	connectLimit := 10
	if filename == "input.txt" {
		connectLimit = 1000
	}

	allBoxes := []*Box{}
	allCircuits := []*Circuit{}
	for line := range strings.SplitSeq(input, "\n") {
		newBox := buildBox(line, allBoxes)
		allBoxes = append(allBoxes, newBox)
		allCircuits = append(allCircuits, newBox.circuit)
	}

	allPossibleConnections := mapAllConnections(allBoxes)

	keys := []int{}
	for k := range allPossibleConnections {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := range connectLimit {
		pair := allPossibleConnections[keys[i]]
		debugPrintBoxPair("Eval: ", pair)
		connect(pair)
	}
	sizes := getCircuitSizes(allCircuits)
	debugPrintCircuits(allCircuits)

	slices.Sort(sizes)
	fmt.Printf("\nSorted sizes: %s%v%s\n", GRN, sizes, NC)
	for i := range 3 {
		result *= sizes[(len(sizes)-1)-i]
	}
	return result
}

func main() {
	filename := os.Args[1]
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)

	result1 := evalPart1(body, filename)
	fmt.Printf("Part 1: %d\n", result1)
}
