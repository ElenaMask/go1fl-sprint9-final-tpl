package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SIZE           = 100_000_000
	CHUNKS         = 8
	MIN_CHUNK_SIZE = 1000
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	var rsl []int = make([]int, size)
	if size > 0 {
		for i, _ := range rsl {
			rsl[i] = rand.Int()
		}
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
	l := len(data)
	if l < CHUNKS*MIN_CHUNK_SIZE {
		return maximum(data)
	}
	increment := l / CHUNKS
	start := 0
	maxCh := make(chan int, CHUNKS)
	routineCounter := 0
	for start < l {
		end := start + increment
		if end > l || routineCounter == CHUNKS-1 {
			end = l
		}

		go func(start, end int) {
			max := data[start]
			for i := start; i < end; i++ {
				if data[i] > max {
					max = data[i]
				}
			}
			maxCh <- max
		}(start, end)
		routineCounter++

		start = end
	}

	max := data[0]
	for i := 0; i < routineCounter; i++ {
		localMax := <-maxCh
		if localMax > max {
			max = localMax
		}
	}
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
