package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Benchmark for the naive algorithm with display
func BenchmarkNaive(b *testing.B) {
	fmt.Println("Starting Benchmark for Naive Algorithm...")
	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration %d of Naive Algorithm\n", i+1)
		naive("input.txt")
	}
	fmt.Println("Completed Benchmark for Naive Algorithm.")
}

// Benchmark for the optimized algorithm with display
func BenchmarkOptimized(b *testing.B) {
	fmt.Println("Starting Benchmark for Optimized Algorithm...")
	for i := 0; i < b.N; i++ {
		fmt.Printf("Iteration %d of Optimized Algorithm\n", i+1)
		algo("input.txt")
	}
	fmt.Println("Completed Benchmark for Optimized Algorithm.")
}

// Unit tests with verbose output
func TestNaive(t *testing.T) {
	fmt.Println("\n🔍 Running TestNaive: Testing with very small maze (input_test.txt)")
	result := algo("input_small.txt")
	assert.Equal(t, 3, result, "Expected 3 for very small maze")
	fmt.Println("✅ TestNaive Passed!")
}

func TestOptimized(t *testing.T) {
	fmt.Println("\n🔍 Running TestOptimized: Testing with large maze (input_small.txt)")
	result := algo("input_large.txt")
	assert.Equal(t, 154, result, "Expected 154 for large maze")
	fmt.Println("✅ TestOptimized Passed!")
}

func TestLargeScale(t *testing.T) {
	fmt.Println("\n🔍 Running TestLargeScale: Testing with a very large maze (input.txt)")
	result := algo("input_very_large.txt")
	assert.Equal(t, 155, result, "Expected 155 for a very large maze")
	fmt.Println("✅ TestLargeScale Passed!")
}
