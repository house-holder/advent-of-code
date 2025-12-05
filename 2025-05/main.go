package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func evalPart1(ranges []string, ingredients []string) int {
	count := 0

	for _, ingredient := range ingredients {
		id, err := strconv.Atoi(ingredient)
		if err != nil {
			log.Fatalf("strconv.Atoi failed. %v", err)
		}

		for _, rangeStr := range ranges {
			loBound, err := strconv.Atoi(strings.Split(rangeStr, "-")[0])
			if err != nil {
				log.Fatalf("strconv.Atoi failed. %v", err)
			}
			hiBound, err := strconv.Atoi(strings.Split(rangeStr, "-")[1])
			if err != nil {
				log.Fatalf("strconv.Atoi failed. %v", err)
			}

			if id >= loBound && id <= hiBound {
				count++
				break
			}
		}
	}
	return count
}

func evalPart2(ranges []string) int {
	count := 0
	maxValueSeen := 0
	ingrRanges := [][]int{}

	for _, rangeStr := range ranges {
		loBound, err := strconv.Atoi(strings.Split(rangeStr, "-")[0])
		if err != nil {
			log.Fatalf("strconv.Atoi failed. %v", err)
		}
		hiBound, err := strconv.Atoi(strings.Split(rangeStr, "-")[1])
		if err != nil {
			log.Fatalf("strconv.Atoi failed. %v", err)
		}
		ingrRanges = append(ingrRanges, []int{loBound, hiBound})
	}

	sort.Slice(ingrRanges, func(i, j int) bool {
		return ingrRanges[i][0] < ingrRanges[j][0]
	})
	// fmt.Printf("ranges: %v\n", ingrRanges)

	for _, ingrRange := range ingrRanges {
		// fmt.Printf("Evaluating %v", ingrRange)
		if ingrRange[1] < maxValueSeen {
			// fmt.Printf(" skip, upper bound less than max\n")
			continue
		}
		// fmt.Printf("\n")

		upper := ingrRange[1]
		lower := int(math.Max(float64(maxValueSeen)+1, float64(ingrRange[0])))
		maxValueSeen = ingrRange[1]
		// fmt.Printf("\tsetting lower to: %d, maxValueSeen to: %d\n", lower, maxValueSeen)

		diff := upper - lower + 1
		count += diff
		// fmt.Printf("\tfound %d ingredients to add to count\n", diff)
	}

	return count
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)

	ranges := []string{}
	ingredients := []string{}
	for part := range strings.FieldsSeq(body) {
		if strings.Contains(string(part), "-") {
			ranges = append(ranges, string(part))
		} else {
			ingredients = append(ingredients, string(part))
		}
	}

	part1 := evalPart1(ranges, ingredients)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := evalPart2(ranges)
	fmt.Printf("Part 2: %d\n", part2)
}
