package main

import "fmt"

func main() {
	//Ініціалізація змінних
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 //Такий варіант ініціалізації також можливий

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	//Скорочений запис оголошення змінної
	//тільки для нових змінних
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	//Завдання.
	//1. Вивести типи усіх змінних
	//2. Присвоїти змінній intVar змінні userinit16 та userautoinit. Результат вивести.
	fmt.Printf("Value = %d Type = %T\n", userinit8, userinit8)
	fmt.Printf("Value = %d Type = %T\n", userinit16, userinit16)
	fmt.Printf("Value = %d Type = %T\n", userinit64, userinit64)
	fmt.Printf("Value = %d Type = %T\n", userautoinit, userautoinit)

	intVar = int(userinit16)
	fmt.Println("intVar =", intVar)
	intVar = userautoinit
	fmt.Println("intVar =", intVar)
}
