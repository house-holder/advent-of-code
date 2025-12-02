package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convert(lines []string) [][]int {
	output := [][]int{}

	for _, line := range lines {
		intLine := []int{}
		for char := range strings.SplitSeq(line, " ") {
			intVal, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal("Failed to convert int: ", err)
			}
			intLine = append(intLine, intVal)
		}
		output = append(output, intLine)
	}
	return output
}

func walk(line []int) bool {
	fmt.Printf("Line: %d\n", line)
	first := line[0]
	second := line[1]
	diff := second - first
	positive := true
	if diff < 0 {
		positive = false
	}

	for i := range len(line) - 1 {
		first := line[i]
		second := line[i+1]
		diff = second - first
		fmt.Printf("  > Evaluating %d and %d...\n", first, second)
		if i+1 > len(line) {
			break
		}
		if diff > 3 || diff < -3 || diff == 0 {
			fmt.Printf("        Bad diff: %d\n", diff)
			return false
		}
		if positive {
			if first > second {
				fmt.Println("        Failed positive trend")
				return false
			}
		} else {
			if first < second {
				fmt.Println("        Failed negative trend")
				return false
			}
		}
	}
	return true
}

func process(dataLines [][]int) int {
	safeCount := 0
	for _, line := range dataLines {
		safe := walk(line)
		if safe {
			safeCount++
			fmt.Printf("        OK - Count: %d\n", safeCount)
		}
	}
	return safeCount
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 && os.Args[1] == "-t" {
		filename = "example.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open: %v", err)
	}
	defer file.Close()

	dataStrings := []string{}
	bios := bufio.NewScanner(file)
	for bios.Scan() {
		line := bios.Text()
		dataStrings = append(dataStrings, line)
	}

	dataLines := convert(dataStrings)
	count := process(dataLines)

	fmt.Printf("Safe count: %d\n", count)
}
