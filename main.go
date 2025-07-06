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
	if size < 0 {
		return nil
	}
	var rsl []int = make([]int, size)
	for i := range rsl {
		rsl[i] = rand.Int()
	}
	return rsl
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	len := len(data)
	if len < CHUNKS {
		return maximum(data)
	}
	increment := len / CHUNKS
	start := 0
	localMaxs := make([]int, CHUNKS)
	var wg sync.WaitGroup
	var mu sync.Mutex
	routineCounter := 0
	for start < len {
		end := start + increment
		if routineCounter == CHUNKS-1 {
			end = len
		}
		wg.Add(1)
		go func(start, end, n int) {
			defer wg.Done()
			max := maximum(data[start:end])
			mu.Lock()
			localMaxs[n] = max
			mu.Unlock()
		}(start, end, routineCounter)
		routineCounter++
		start = end
	}
	wg.Wait()
	max := maximum(localMaxs)
	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
