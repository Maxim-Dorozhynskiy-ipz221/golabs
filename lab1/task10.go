package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Завдання.
	// 1. Вивести українську літеру 'Ї'
	// 2. Пояснити призначення типу "rune"
	var ch = "Її"
	fmt.Println(ch)
	/* rune є еквівалентом для типу даних int32.
	Використовується, щоб відрізнити символьні значення від цілих значень.*/
}
