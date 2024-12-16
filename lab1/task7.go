package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання.
	// 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести

	var a int8 = 18
	var b float32 = 3.6
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Printf("a + b = %-5f\n", float32(a)+b)
	fmt.Printf("a - b = %-5f\n", float32(a)-b)
	fmt.Printf("a * b = %-5f\n", float32(a)*b)
	fmt.Printf("a / b = %-5f\n", float32(a)/b)
	a++
	fmt.Printf("a++ = %-5d\n", a)
	a--
	fmt.Printf("a-- = %-5d\n", a)
}
