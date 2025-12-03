package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}

	body := string(bytes)
	fmt.Println(body)
}
