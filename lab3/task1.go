package main

import (
	"fmt"
	"math"
)

// 9 варіант

func main() {
	var x0 int32 = 0
	var arr []int32
	RandomValue(x0, &arr) // Передача посилання на arr
	Reverse(&arr)

	// Створюємо масив частот
	var freq [250]int32
	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}

	// Виведення статистики для кожного значення
	for i := 0; i < 250; i++ {
		if freq[i] > 0 {
			fmt.Printf("№ %d %d К-ть повторів: %d Статична ймовірність: %f\n", i+1, i, freq[i], float32(freq[i])/20000)
		}
	}

	// Математичне сподівання
	exp := MathExp(&arr, &freq)
	fmt.Printf("\nМатематичне сподівання випадкових величин: %v", exp)

	// Обчислення дисперсії
	var dysp float64
	for i := 0; i < 250; i++ {
		if freq[i] > 0 {
			dysp += math.Pow(float64(i)-exp, 2) * float64(freq[i]) / 20000
		}
	}
	fmt.Printf("\nДисперсія псевдовипадкових величин: %f", dysp)

	// Середньоквадратичне відхилення
	sigma := Sigma(&arr)
	fmt.Printf("\nСередньоквадратичне відхилення випадкових величин: %f", sigma)
}

// Функція для генерації випадкових чисел
func RandomValue(x0 int32, arr *[]int32) {
	var a = 1140671485
	var c = 12820163
	var m = int32(math.Pow(2, 24))

	// Генерація 20000 випадкових чисел
	for k := 0; k < 20000; k++ {
		x0 = (a*x0 + c) % m
		*arr = append(*arr, x0%250) // Генерація числа від 0 до 249
	}
}

// Функція для обчислення математичного сподівання
func MathExp(arr *[]int32, frequency *[250]int32) float64 {
	var mathExp float64
	for i := 0; i < len(*arr); i++ {
		mathExp += float64((*arr)[i]) * float64((*frequency)[(*arr)[i]]) / 20000
	}
	return mathExp
}

// Функція для обчислення середньоквадратичного відхилення
func Sigma(arr *[]int32) float64 {
	var sigma float64
	var sum float64
	var n = len(*arr)
	for i := 0; i < n; i++ {
		sum += float64((*arr)[i])
	}
	avg := sum / float64(n)

	for i := 0; i < n; i++ {
		sigma += math.Pow(float64((*arr)[i])-avg, 2)
	}
	sigma = math.Sqrt(sigma / float64(n))
	return sigma
}

// Функція для перевертання масиву
func Reverse(numbers *[]int32) {
	for i := 0; i < len(*numbers)/2; i++ {
		j := len(*numbers) - i - 1
		(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
	}
}
