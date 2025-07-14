package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"Zero size", 0},
		{"Small size", 10},
		{"Large size", 100_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.size {
				t.Errorf("generateRandomElements() length = %v, want %v", len(got), tt.size)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"All negative", []int{-5, -3, -8}, -3},
		{"Mixed values", []int{-1, 0, 1, -2, 2}, 2},
		{"All equal", []int{7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.expected {
				t.Errorf("maximum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Less than chunks", []int{1, 2, 3}, 3},
		{"Exact chunks", make([]int, CHUNKS), 0},
		{"Large data", make([]int, 1000), rand.Intn(100000)},
	}

	// Initialize large data
	rand.Seed(time.Now().UnixNano())
	for i := range tests[3].data {
		tests[3].data[i] = rand.Intn(100000)
	}
	tests[3].expected = maximum(tests[3].data)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.expected {
				t.Errorf("maxChunks() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConcurrency(t *testing.T) {
	data := generateRandomElements(1000)
	var wg sync.WaitGroup
	results := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = maxChunks(data)
		}(i)
	}
	wg.Wait()

	first := results[0]
	for _, res := range results {
		if res != first {
			t.Errorf("maxChunks() вернул разные результаты: %v", results)
			break
		}
	}
}
