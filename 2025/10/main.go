package main

import (
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func parseU16(c string, msg string) uint16 {
	num, err := strconv.Atoi(c)
	if err != nil {
		log.Fatalf("%s failed. %v", msg, err)
	}
	return uint16(num)
}

func idt(depth int) string {
	return strings.Repeat("  ", depth)
}

type Machine struct {
	voltages []uint16
	buttons  []uint16
	stateTgt uint16
	state    uint16
}

func NewMachine(line string) Machine {
	var tgtState uint16 = 0
	accB, accV := []uint16{}, []uint16{}

	for cmp := range strings.SplitSeq(line, " ") {
		switch leader := cmp[0]; leader {
		case '[':
			tCmp := strings.Trim(cmp, "[]")
			for i, c := range tCmp {
				if c == '#' {
					tgtState |= 1 << i
				}
			}
		case '(':
			var newB uint16 = 0
			tCmp := strings.Trim(cmp, "()")
			for c := range strings.SplitSeq(tCmp, ",") {
				n := parseU16(c, "button parse")
				newB |= 1 << n
			}
			accB = append(accB, newB)
		case '{':
			tCmp := strings.Trim(cmp, "{}")
			for c := range strings.SplitSeq(tCmp, ",") {
				newV := parseU16(c, "voltage parse")
				accV = append(accV, newV)
			}
		}
	}
	return Machine{
		voltages: accV,
		buttons:  accB,
		stateTgt: tgtState,
		state:    0,
	}
}

func evalPart1(machines []*Machine) (operations int) {
	operations = 0
	for _, m := range machines {
		minOps := math.MaxInt
		k := len(m.buttons)
		limit := 1 << k

		for subset := 1; subset < limit; subset++ {
			ops := bits.OnesCount(uint(subset))
			if ops >= minOps {
				continue
			}
			m.state = 0
			for i := range k {
				if subset&(1<<i) != 0 {
					m.state ^= m.buttons[i]
				}
			}
			if m.state == m.stateTgt {
				minOps = ops
			}
		}
		operations += minOps
	}
	return operations
}

func evalPart2(machines []*Machine) (operations int) {
	operations = 0
	for _, m := range machines {
		numBtns := len(m.buttons)
		totalCombos := 1 << numBtns
		numVTargets := len(m.voltages)
		results := make(map[int][]uint16)
		patterns := make(map[string][]int)

		for combo := range totalCombos {
			voltageVals := make([]uint16, numVTargets)

			for bIdx := range numBtns {
				if combo&(1<<bIdx) != 0 {
					activeBtn := m.buttons[bIdx]

					for cIdx := range numVTargets {
						if activeBtn&(1<<cIdx) != 0 {
							voltageVals[cIdx]++
						}
					}
				}
			}
			parity := make([]uint16, numVTargets)
			for i := range numVTargets {
				parity[i] = voltageVals[i] % 2
			}

			parityKey := fmt.Sprint(parity)
			patterns[parityKey] = append(patterns[parityKey], combo)
			results[combo] = voltageVals
		}

		machineMinOps := m.execute(m.voltages, patterns, results, 0)
		operations += machineMinOps
	}
	return operations
}

func (m *Machine) execute(
	targets []uint16,
	patterns map[string][]int,
	results map[int][]uint16,
	depth int,
) (machineMinOps int) {
	base := true
	for _, v := range targets {
		if v != 0 {
			base = false
			break
		}
	}
	if base {
		// fmt.Printf("%sBASE: tgts=%v → return 0\n", idt(depth), targets)
		return 0
	}

	parity := make([]uint16, len(targets))
	for i := range targets {
		parity[i] = targets[i] % 2
	}
	parityKey := fmt.Sprint(parity)
	matches := patterns[parityKey]
	// fmt.Printf("%s>>> exec(targets=%v) parity=%s, %d matches\n",
	// 	idt(depth), targets, parityKey, len(matches))

	minOps := math.MaxInt
	for _, combo := range matches {
		diff := results[combo]
		// fmt.Printf("%s  combo=%b, diff=%v\n", idt(depth), combo, diff)
		valid := true
		for i := range targets {
			if diff[i] > targets[i] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		newTarget := make([]uint16, len(targets))
		for i := range targets {
			newTarget[i] = (targets[i] - diff[i]) / 2
		}
		// fmt.Printf("%s  newTarget = (targets - diff) / 2 = %v\n",
		//  idt(depth), newTarget)
		presses := bits.OnesCount(uint(combo))
		subResult := m.execute(newTarget, patterns, results, depth+1)
		// fmt.Printf("%s  !R: presses=%d, subR=%d, tot=%d+2*%d=%d\n",
		// 	idt(depth), presses, subResult, presses, subResult, presses+2*subResult)

		if subResult < math.MaxInt {
			total := presses + 2*subResult
			if total < minOps {
				minOps = total
			}
		}
	}
	// fmt.Printf("%s<<< Return minOps=%d for targets=%v\n",
	//  idt(depth), minOps, targets)
	return minOps
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	lines := strings.Split(string(bytes), "\n")
	machines := []*Machine{}
	for _, line := range lines {
		if line != "" {
			newM := NewMachine(line)
			machines = append(machines, &newM)
		}
	}

	result1 := evalPart1(machines)
	fmt.Println("Part 1:", result1)
	result2 := evalPart2(machines)
	fmt.Println("Part 2:", result2)
}
