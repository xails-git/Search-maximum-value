package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	slice := make([]int, 0, size)
	if size == 0 {
		fmt.Printf("size для слайса = 0")
		return slice
	}
	n := 100000

	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(n))
	}

	return slice

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	counter := data[0]

	for i := 0; i < len(data); i++ {
		if data[i] > counter {
			counter = data[i]
		}
	}
	return counter

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS
	maxValues := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {

		start := i * chunkSize
		end := start + chunkSize

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if start >= len(data) {
				return
			}
			if end > len(data) {
				end = len(data)
			}

			chunk := data[start:end]
			if len(chunk) == 0 {
				return
			}

			max := maximum(chunk)

			maxValues[i] = max

		}(i)
	}
	wg.Wait()

	if len(maxValues) == 0 {
		return 0
	}

	counter := maximum(maxValues)
	return counter
}

func main() {

	fmt.Printf("Генерируем %d целых чисел", SIZE)
	generateRandomElements(SIZE)

	fmt.Println("Ищём максимальное значение в один поток")

	start := time.Now()
	max := maximum(generateRandomElements(SIZE))
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	startCh := time.Now()
	maxChunk := maxChunks(generateRandomElements(SIZE))
	elapsedCh := time.Since(startCh)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxChunk, elapsedCh)
}
