package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Network struct {
	devices map[string][]string
}

func (n *Network) followPaths(curr string, tgt string) int {
	if curr == tgt {
		return 1
	}
	acc := 0
	for _, connection := range n.devices[curr] {
		acc += n.followPaths(connection, tgt)
	}

	return acc
}

func evalPart1(lines []string) (output int) {
	network := Network{
		make(map[string][]string),
	}
	for _, ln := range lines {
		parts := strings.Split(ln, ":")
		key := parts[0]
		val := strings.Split(strings.Trim(parts[1], " "), " ")
		network.devices[key] = val
	}
	return network.followPaths("you", "out")
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	lines := strings.Split(strings.Trim(body, "\n"), "\n")

	result1 := evalPart1(lines)
	fmt.Println("Part 1:", result1)
}
