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

var debug = true

func Dbg(format string, a ...any) {
	if debug {
		fmt.Printf(format, a...)
	}
}

func parseU16(c string, msg string) uint16 {
	num, err := strconv.Atoi(c)
	if err != nil {
		log.Fatalf("%s failed. %v", msg, err)
	}
	return uint16(num)
}

type Machine struct {
	voltageTargets []uint16
	voltages       []uint16
	buttons        []uint16
	stateTarget    uint16
	state          uint16
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
		voltageTargets: accV,
		voltages:       make([]uint16, len(accV)),
		buttons:        accB,
		stateTarget:    tgtState,
		state:          0,
	}
}

func (m *Machine) runState() (minOps int) {
	k := len(m.buttons)
	limit := 1 << k
	minOps = math.MaxInt

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
		if m.state == m.stateTarget {
			minOps = ops
		}
	}
	return minOps
}

// what the eff do i do for part 2

func evalPart1(machines []*Machine) int {
	operations := 0

	for _, machine := range machines {
		ops := machine.runState()
		operations += ops
	}
	return operations
}

func evalPart2(machines []*Machine) (operations int) {
	operations = 0

	for _, m := range machines {
		operations += len(m.voltages)
	}
	return operations
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
	fmt.Printf("Part 2: DUMMY VALUE %d\n", result2)
}
