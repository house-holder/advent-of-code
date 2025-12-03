package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func collect(data string) [][]int {
	enabled := true
	output := [][]int{}

	mulRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	onRegex := regexp.MustCompile(`do\(\)`)
	offRegex := regexp.MustCompile(`don\'t\(\)`)

	tokens := mulRegex.FindAllString(data, -1)

	for _, token := range tokens {
		if offRegex.MatchString(token) {
			enabled = false
		} else if onRegex.MatchString(token) {
			enabled = true
		}
		if enabled {
			trimmed := strings.Trim(token, "mul()")
			parts := strings.Split(trimmed, ",")
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])

			output = append(output, []int{num1, num2})
		}
	}
	return output
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

	stringData := ""
	bios := bufio.NewScanner(file)
	for bios.Scan() {
		stringData += bios.Text()
	}

	data := collect(stringData)

	sum := 0
	for _, pair := range data {
		sum += pair[0] * pair[1]
	}
	fmt.Println(sum)
}
