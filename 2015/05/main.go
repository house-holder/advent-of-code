package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func evalPart2(lines []string) int {
	count := 0

	// core logic

	return count
}

func evalPart1(lines []string) int {
	count := 0
	badStrings := []string{"ab", "cd", "pq", "xy"}

	for _, line := range lines {
		badBool := false

		if line != "" {
			for _, badStr := range badStrings {
				if strings.Contains(line, badStr) {
					badBool = true
					break
				}
			}
			if badBool {
				continue
			}

			vowelCount := 0
			hasDouble := false

			for i, char := range line {
				if vowelCount < 3 && strings.Contains("aeiou", string(char)) {
					vowelCount++
				}

				if !hasDouble {
					if (i+1) < len(line) &&
						string(line[i+1]) == string(char) {
						hasDouble = true
					}
				}
			}
			if hasDouble && vowelCount >= 3 {
				count++
			}
		}
	}
	return count
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	lines := strings.Split(body, "\n")

	result1 := evalPart1(lines)
	fmt.Println("Part 1: ", result1)

	// result2 := evalPart2(lines)
	// fmt.Println("Part 2: ", result2)
}
