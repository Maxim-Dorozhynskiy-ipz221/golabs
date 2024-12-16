package main

// Імпорт декількох пакетів
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи скорочений запис і ініціалізацію за замовчуванням. Результат вивести
	var (
		x int16
		y bool
	)
	char := ']'
	compl := 3 + 2.2i
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("char = ", char)
	fmt.Println("complex = ", compl)
}
