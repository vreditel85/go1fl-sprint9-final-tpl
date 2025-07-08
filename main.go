package main

import (
	"fmt"
	"math"
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
	// Обработка случая, когда размер слайса равен 0
	if size <= 0 {
		return nil
	}

	// Инициализация генератора случайных чисел с использованием текущего времени
	rand.Seed(time.Now().UnixNano())

	// Создание слайса заданного размера
	slice := make([]int, size)

	// Заполнение слайса случайными числами
	for i := 0; i < size; i++ {
		slice[i] = rand.Int() // Генерация случайного int
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := math.MinInt // Начинаем с минимально возможного int

	for _, num := range data {
		if num > max {
			max = num
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return math.MinInt
	}

	var wg sync.WaitGroup
	maxes := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(chunkIndex int) {
			defer wg.Done()

			start := chunkIndex * chunkSize
			end := start + chunkSize
			if chunkIndex == CHUNKS-1 {
				end = len(data) // Последний чанк может быть больше
			}

			chunk := data[start:end]
			max := math.MinInt
			for _, num := range chunk {
				if num > max {
					max = num
				}
			}
			maxes[chunkIndex] = max
		}(i)
	}

	wg.Wait()

	// Находим максимум среди всех чанков
	overallMax := math.MinInt
	for _, m := range maxes {
		if m > overallMax {
			overallMax = m
		}
	}

	return overallMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	if len(data) == 0 {
		fmt.Println("Ошибка генерации данных")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
