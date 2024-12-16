package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Синоніми цілих типів\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, або int64, залежно від ОС")
	fmt.Println("uint    - uint32, або uint64, залежно від ОС")

	// Завдання.
	// 1. Визначити розрядність ОС
	fmt.Println("\nПорівнюємо максимальні int значення")
	fmt.Println("Значення int32", math.MaxInt32)
	fmt.Println("Значення int64", math.MaxInt64)
	fmt.Println("Значення int", math.MaxInt)
	fmt.Println("Розрядність ОС - 64")
	// Розрядність ОС - 64
}
