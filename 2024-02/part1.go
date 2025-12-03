package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	FAIL = "\033[31m✗\033[0m"
	OK   = "\033[32m✓\033[0m"
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

func safe(line []int) bool {
	fmt.Printf("%d ", line)

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

		if i+1 > len(line) {
			break
		}
		if diff > 3 || diff < -3 || diff == 0 {
			fmt.Printf("%s bad diff\n", FAIL)
			return false
		}
		if positive {
			if first > second {
				fmt.Printf("%s positive trend\n", FAIL)
				return false
			}
		} else {
			if first < second {
				fmt.Printf("%s negative trend\n", FAIL)
				return false
			}
		}
	}
	return true
}

func safeWithRemoval(line []int) bool {
	for i := range len(line) {
		newLine := []int{}
		for j := range len(line) {
			if j != i {
				newLine = append(newLine, line[j])
			}
		}
		fmt.Printf("    remove (i=%d) %d ", i, line[i])
		if safe(newLine) {
			return true
		}
		continue
	}
	return false
}

func process(dataLines [][]int) int {
	safeCount := 0
	for i, line := range dataLines {
		fmt.Printf("> New eval: line %d ", i+1)
		if safe(line) {
			safeCount++
			fmt.Printf("%s count:   %d\n", OK, safeCount)
		} else if safeWithRemoval(line) {
			safeCount++
			fmt.Printf("%s removal: %d\n", OK, safeCount)
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

	fmt.Printf("\nFinal count: %d\n", count)
}
