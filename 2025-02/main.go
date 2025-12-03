package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func hasPattern(id int) bool {
	idStr := strconv.Itoa(id)

	for i := 1; i <= len(idStr)/2; i++ {
		if len(idStr)%i == 0 {
			pattern := idStr[0:i]
			if strings.Repeat(pattern, len(idStr)/i) == idStr {
				fmt.Printf("   > Found: %s\n", idStr)
				return true
			}
		}
	}
	return false
}

func evaluateRange(firstID int, lastID int, accumulator *int) {
	for i := firstID; i <= lastID; i++ {
		if hasPattern(i) {
			*accumulator += i
		}
	}
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 && os.Args[1] == "-t" {
		filename = "example.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open: %v", err)
	}
	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	accumulator := 0

	for ID := range strings.SplitSeq(string(body), ",") {
		thisRange := strings.Split(ID, "-")
		fmt.Printf("Evaluating range: %v\n", thisRange)

		firstID, err := strconv.Atoi(thisRange[0])
		if err != nil {
			log.Fatalf("Failed to convert: %v", err)
		}

		lastID, err := strconv.Atoi(thisRange[1])
		if err != nil {
			log.Fatalf("Failed to convert: %v", err)
		}

		evaluateRange(firstID, lastID, &accumulator)
		fmt.Printf("  Acc: %v\n", accumulator)
	}
}
