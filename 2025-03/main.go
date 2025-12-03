package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}
	return num
}

func evaluatePart1(line string) int {
	hiL := 0
	hiR := 0
	idx := -1

	digits := []int{}
	for _, digitChar := range strings.Split(line, "") {
		digits = append(digits, strToInt(digitChar))
	}

	for i, digitL := range digits[:len(digits)-1] {
		if digitL > hiL {
			// fmt.Printf("digitL: %d, hiL: %d\n", digitL, hiL)
			hiL = digitL
			idx = i
		}
	}

	for _, digitR := range digits[idx+1:] {
		if digitR > hiR {
			// fmt.Printf("digitR: %d, hiR: %d\n", digitR, hiR)
			hiR = digitR
		}
	}
	return hiL*10 + hiR
}

func evaluatePart2(bank string) int {
	// if resultSize is set, works perfectly for part 1
	resultSize := 12
	outputSum := 0
	digits := []int{}
	for digitChar := range strings.SplitSeq(bank, "") {
		digits = append(digits, strToInt(digitChar))
	}
	// fmt.Printf("New eval: %s\n", bank)

	accum := []int{}
	startIdx := 0

	for len(accum) < resultSize {
		var digit int
		stopIdx := 1 + len(accum) + len(digits) - resultSize
		digit, startIdx = search(digits, startIdx, stopIdx)

		accum = append(accum, digit)
		// fmt.Printf("append %d\n", digit)
	}
	// fmt.Printf("  Accumulator: %v\n", accum)

	// reverse(&accum) // need "least significant decimal" order
	// fmt.Printf("Reversed: %v\n", accum)
	// for i, n := range accum {
	// 	sum := n * int(math.Pow(10, float64(i)))
	// 	fmt.Printf("sum: %d,", sum)
	// 	outputSum += sum
	// 	fmt.Printf("outputSum: %d\n", outputSum)
	// }

	for i, n := range accum {
		power := len(accum) - i - 1
		sum := n * int(math.Pow(10, float64(power)))
		outputSum += sum
	}
	return outputSum
}

func search(digits []int, startIdx int, stopIdx int) (digit int, idx int) {
	// fmt.Printf("    start=%d, stop=%d | ", startIdx, stopIdx)
	digit = 0
	idx = startIdx
	for i := startIdx; i < stopIdx; i++ {
		if digits[i] > digit {
			digit = digits[i]
			idx = i + 1
		}
	}
	// fmt.Printf("(%d) i=%d | ", digit, idx)
	return digit, idx
}

// func reverse(arr *[]int) {
// 	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
// 		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
// 	}
// }

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}

	body := string(bytes)
	lines := strings.Split(body, "\n")

	sum1 := 0
	sum2 := 0
	for _, bank := range lines {
		sum1 += evaluatePart1(bank)
		sum2 += evaluatePart2(bank)
	}
	fmt.Println(">>> Part 1:", sum1)
	fmt.Println(">>> Part 2:", sum2)
}
