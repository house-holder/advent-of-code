package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 && os.Args[1] == "-t" {
		filename = "example.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Println(lines)
}
