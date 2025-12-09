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

func Debug(format string, a ...any) {
	fmt.Printf(format, a...)
}

type Circuit struct {
	name        string
	connections []*JBox
	size        int
}

var circuitIdx = 0
var circuits = []*Circuit{}

func createCircuit() *Circuit {
	n := fmt.Sprintf("circuit %d", circuitIdx)
	circuitIdx++
	circuit := &Circuit{
		name:        n,
		connections: []*JBox{},
		size:        0,
	}
	circuits = append(circuits, circuit)
	return circuit
}

func (c *Circuit) GetConnections() []string {
	output := []string{}
	for _, n := range c.connections {
		output = append(output, n.name)
	}
	return output
}

func (own *Circuit) Append(target *JBox) {
	target.Assign(own)
	own.connections = append(own.connections, target)
	own.size++
}

func CombineCircuits(a *Circuit, b *Circuit) *Circuit {
	// fmt.Printf("\tCombine %s + %s\n", a.name, b.name)
	newCircuit := createCircuit()
	newConnections := a.connections
	for _, conn := range b.connections {
		if slices.Contains(newConnections, conn) {
			continue
		}
		newConnections = append(newConnections, conn)
	}
	newCircuit.size = len(newCircuit.connections)
	newCircuit.connections = newConnections
	return newCircuit
}

var noCircuit = createCircuit()

type JBox struct {
	name        string
	X           int
	Y           int
	Z           int
	connections []*JBox
	isConnected bool
	circuit     *Circuit
}

func createBox(index int, coordLine string) *JBox {
	parts := strings.Split(coordLine, ",")
	strX, strY, strZ := parts[0], parts[1], parts[2]

	xVal, err := strconv.Atoi(strX)
	if err != nil {
		log.Fatal("strconv.Atoi failed. ", err)
	}
	yVal, err := strconv.Atoi(strY)
	if err != nil {
		log.Fatal("strconv.Atoi failed. ", err)
	}
	zVal, err := strconv.Atoi(strZ)
	if err != nil {
		log.Fatal("strconv.Atoi failed. ", err)
	}

	return &JBox{
		name:        fmt.Sprintf("jbox%d", index),
		X:           xVal,
		Y:           yVal,
		Z:           zVal,
		connections: []*JBox{},
		isConnected: false,
		circuit:     noCircuit,
	}
}

func (box *JBox) Connect(other *JBox) {
	if !slices.Contains(box.connections, other) {
		box.connections = append(box.connections, other)
	}
	if !slices.Contains(other.connections, box) {
		other.connections = append(other.connections, box)
	}
	if slices.Contains(box.circuit.connections, other) {
		fmt.Println("OH NO RETURN")
		return
	}
	if !box.isConnected && !other.isConnected {
		newCkt := createCircuit()
		newCkt.Append(box)
		newCkt.Append(other)
		Debug("    New: %s\n", newCkt.name)
		Debug("\tContents: %v\n", newCkt.GetConnections())
	} else {
		if box.isConnected {
			box.circuit.Append(other)
			Debug("    Box append %s to %s\n", other.name, box.circuit.name)
			Debug("\tContents: %v\n", box.circuit.GetConnections())
		} else if other.isConnected {
			other.circuit.Append(box)
			Debug("    Other append %s to %s\n", box.name, other.circuit.name)
			Debug("\tContents: %v\n", other.circuit.GetConnections())
		} else {
			CombineCircuits(box.circuit, other.circuit)
			Debug("    Combine circuits: %s %s\n", box.circuit.name, other.circuit.name)
			Debug("\tContents: %v\n", box.circuit.GetConnections())
		}
	}
	fmt.Println()
}

func (box *JBox) GetCoords() string {
	return fmt.Sprintf("%d,%d,%d", box.X, box.Y, box.Z)
}

func (box *JBox) Assign(ckt *Circuit) {
	box.circuit = ckt
	box.isConnected = true
}

func (box *JBox) GetDist(other *JBox) float64 {
	dX, dY, dZ := other.X-box.X, other.Y-box.Y, other.Z-box.Z
	dist := math.Sqrt(float64(dX*dX) + float64(dY*dY) + float64(dZ*dZ))
	// Debug("Dist: %s to %s: %.2f\n", box.name, other.name, dist)
	return dist
}

func (b *JBox) FindNearest(boxes []*JBox, selfIdx int) (*JBox, float64) {
	minDist := math.Inf(1)
	minIdx := 0
	for i, box := range boxes {
		if i == selfIdx {
			continue
		}
		dist := b.GetDist(box)
		if dist < minDist {
			minDist = dist
			minIdx = i
		}
	}
	return boxes[minIdx], minDist
}

func GetAllDistances(boxes []*JBox) ([]float64, map[float64][]*JBox) {
	distMap := make(map[float64][]*JBox)
	distances := []float64{}
	var otherBox *JBox
	minDist := 0.0

	for boxIdx, box := range boxes {
		otherBox, minDist = box.FindNearest(boxes, boxIdx)
		pair, ok := distMap[minDist]
		if ok {
			if slices.Contains(pair, box) ||
				slices.Contains(pair, otherBox) {
				continue
			}
		}
		distMap[minDist] = []*JBox{box, otherBox}
		distances = append(distances, minDist)
		// Debug("%s, %s (d=%.2f)\n", box.name, otherBox.name, minDist)
	}

	sort.Float64s(distances)
	fmt.Println(distances)
	return distances, distMap
}

func ConnectCircuits(dists []float64, dMap map[float64][]*JBox, limit int) {
	for i := range limit {
		pair := dMap[dists[i]]
		boxA := pair[0]
		boxB := pair[1]
		Debug("Connect%d (%.2f)  %s [%s]  %s [%s]\n", i, dists[i],
			boxA.name, boxA.GetCoords(), boxB.name, boxB.GetCoords())
		boxA.Connect(boxB)
	}
}

func GetCircuitSizes() []int {
	circuitSizes := []int{}
	for _, circuit := range circuits[1:] {
		// Debug("\nGCS %d: %s s=%d\n", i+1, circuit.name, circuit.size)
		// for _, connection := range circuit.connections {
		// Debug("  %s\n", connection.name)
		// }
		circuitSizes = append(circuitSizes, circuit.size)
	}
	return circuitSizes
}

func evalPart1(input string) int {
	result := 1
	boxes := []*JBox{}
	limit := 10
	for idx, line := range strings.Split(input, "\n") {
		newBox := createBox(idx, line)
		boxes = append(boxes, newBox)
	}

	distances, distanceMap := GetAllDistances(boxes)
	ConnectCircuits(distances, distanceMap, limit)
	for _, box := range boxes {
		if !box.isConnected {
			newCkt := createCircuit()
			newCkt.Append(box)
		}
	}
	circuitSizes := GetCircuitSizes()

	sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))
	fmt.Println("Circuit sizes:", circuitSizes)

	factors := 3
	for n := range factors {
		result *= circuitSizes[n]
	}
	return result
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}

	body := string(bytes)

	result1 := evalPart1(body)
	fmt.Printf("Part 1: %d\n", result1)
}
