package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func getMD5Hash(key string) string {
	h := md5.New()
	io.WriteString(h, key)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func evalPart1(input string) int {
	num, running := 0, true

	for running {
		key := input + fmt.Sprintf("%d", num)
		hashStr := getMD5Hash(key)

		if hashStr[:5] == "00000" {
			running = false
		} else {
			num++
		}
	}

	return num
}

func evalPart2(input string) int {
	num, running := 282750, true

	for running {
		key := input + fmt.Sprintf("%d", num)
		hashStr := getMD5Hash(key)

		if hashStr[:6] == "000000" {
			running = false
		} else {
			num++
		}
	}

	return num
}

func main() {
	input := "yzbqklnj"

	result1 := evalPart1(input)
	fmt.Println("Part 1: ", result1)

	result2 := evalPart2(input)
	fmt.Println("Part 2: ", result2)
}
