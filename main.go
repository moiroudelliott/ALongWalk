package main

import (
	"fmt"
	"time"
)

func main() {

	res := algo()

	startTime := time.Now()

	duration := time.Since(startTime)

	fmt.Printf("Longest Walk: %d\n", res)
	fmt.Printf("Execution Time: %v\n", duration)
}
